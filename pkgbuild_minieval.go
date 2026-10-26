// minieval: a deliberately small, non-executing evaluator for the metadata
// part of PKGBUILDs. It evaluates top-level assignments (scalars, arrays,
// brace expansion, word splitting, ${...} operators), records functions,
// emulates makepkg's package_*() override extraction, and prints the result
// in .SRCINFO form. Anything outside that subset makes it bail out with a
// reason instead of guessing. It never runs commands.
//
// Usage:
//
//	minieval PKGBUILD
//	minieval DIR        (reads DIR/PKGBUILD, e.g. monorepo/linux)
//
// Prints the .SRCINFO that `makepkg --printsrcinfo` would produce and exits 0.
// If the PKGBUILD uses something outside the supported subset, prints nothing
// on stdout, a reason on stderr, and exits non-zero:
//
//	1  unsupported   (stderr: "unsupported: <reason>")
//	2  usage error
//	3  parse error   (stderr: "parse-error: <message>")
//	4  internal error (a bug; stderr: "crash: <message>")
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"mvdan.cc/sh/v3/syntax"
)

// unsupported is raised (as a panic) when the PKGBUILD uses something outside
// the supported subset. root names the construct; why may add which variable
// needed it.
type unsupported struct{ why, root string }

func bail(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	panic(unsupported{why: msg, root: msg})
}

// ---------- variables ----------

type variable struct {
	arr    map[int]string // indexed array (scalars live at index 0)
	array  bool
	poison string // if set, the value is unknown because of this unsupported construct
}

func (v *variable) keys() []int {
	ks := make([]int, 0, len(v.arr))
	for k := range v.arr {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	return ks
}
func (v *variable) list() []string {
	var out []string
	for _, k := range v.keys() {
		out = append(out, v.arr[k])
	}
	return out
}
func (v *variable) next() int {
	m := -1
	for k := range v.arr {
		if k > m {
			m = k
		}
	}
	return m + 1
}

type env struct {
	vars  map[string]*variable
	funcs map[string]*syntax.FuncDecl
}

// get returns a variable for reading. Reading a variable whose value is
// unknown (see poison) gives up, naming the construct that made it unknown.
func (e *env) get(name string) (*variable, bool) {
	v, ok := e.vars[name]
	if ok && v.poison != "" {
		panic(unsupported{why: fmt.Sprintf("%s (needed for %s)", v.poison, name), root: v.poison})
	}
	return v, ok
}

// poison marks a variable's value as unknown. That is only a problem if
// something needed for the .SRCINFO reads it later.
func (e *env) poison(name, why string) {
	e.vars[name] = &variable{arr: map[int]string{}, poison: why}
}

func (e *env) scalar(name string) (string, bool) {
	v, ok := e.get(name)
	if !ok {
		return "", false
	}
	s, ok := v.arr[0]
	return s, ok
}
func (e *env) setScalar(name, val string, appendTo bool) {
	if appendTo {
		if old, ok := e.get(name); ok {
			val = old.arr[0] + val
		}
	}
	v, ok := e.vars[name]
	if !ok || v.poison != "" {
		v = &variable{arr: map[int]string{}}
		e.vars[name] = v
	}
	v.arr[0] = val
}

// ---------- arithmetic (enough for ${x:off:len} and [i]) ----------

func (e *env) arith(x syntax.ArithmExpr) int {
	if x == nil {
		return 0
	}
	switch n := x.(type) {
	case *syntax.Word:
		s := e.expandScalar(n)
		s = strings.TrimSpace(s)
		if i, err := strconv.Atoi(s); err == nil {
			return i
		}
		if s == "" {
			return 0
		}
		if v, ok := e.scalar(s); ok { // bare variable name
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			return i
		}
		return 0
	case *syntax.ParenArithm:
		return e.arith(n.X)
	case *syntax.UnaryArithm:
		v := e.arith(n.X)
		switch n.Op {
		case syntax.Minus:
			return -v
		case syntax.Plus:
			return v
		}
	case *syntax.BinaryArithm:
		a, b := e.arith(n.X), e.arith(n.Y)
		switch n.Op {
		case syntax.Add:
			return a + b
		case syntax.Sub:
			return a - b
		case syntax.Mul:
			return a * b
		case syntax.Quo:
			if b != 0 {
				return a / b
			}
		case syntax.Rem:
			if b != 0 {
				return a % b
			}
		}
	}
	bail("arithmetic %T", x)
	return 0
}

// ---------- glob patterns (for # % / operators) ----------

func globLitToRegex(sb *strings.Builder, s string) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '\\':
			if i+1 < len(s) {
				i++
				sb.WriteString(regexp.QuoteMeta(string(s[i])))
			}
		case '*':
			sb.WriteString("(?s:.*)")
		case '?':
			sb.WriteString("(?s:.)")
		case '[':
			j := i + 1
			if j < len(s) && (s[j] == '!' || s[j] == '^') {
				j++
			}
			if j < len(s) && s[j] == ']' {
				j++
			}
			for j < len(s) && s[j] != ']' {
				if s[j] == '[' && j+1 < len(s) && s[j+1] == ':' {
					if k := strings.Index(s[j:], ":]"); k > 0 {
						j += k + 2
						continue
					}
				}
				j++
			}
			if j >= len(s) {
				sb.WriteString(`\[`)
				continue
			}
			body := s[i+1 : j]
			if strings.HasPrefix(body, "!") {
				body = "^" + body[1:]
			}
			sb.WriteString("[" + strings.ReplaceAll(body, `\`, `\\`) + "]")
			i = j
		default:
			sb.WriteString(regexp.QuoteMeta(string(c)))
		}
	}
}

func (e *env) pattern(w *syntax.Word) *regexp.Regexp {
	var sb strings.Builder
	sb.WriteString("^(?:")
	if w != nil {
		for _, p := range w.Parts {
			switch x := p.(type) {
			case *syntax.Lit:
				globLitToRegex(&sb, x.Value)
			case *syntax.SglQuoted:
				sb.WriteString(regexp.QuoteMeta(x.Value))
			case *syntax.DblQuoted:
				sb.WriteString(regexp.QuoteMeta(e.expandScalar(&syntax.Word{Parts: []syntax.WordPart{x}})))
			case *syntax.ParamExp: // unquoted expansion in a pattern stays a pattern
				vals, _ := e.param(x)
				globLitToRegex(&sb, strings.Join(vals, " "))
			default:
				bail("pattern part %T", p)
			}
		}
	}
	sb.WriteString(")$")
	re, err := regexp.Compile(sb.String())
	if err != nil {
		bail("pattern regex: %v", err)
	}
	return re
}

// replaceOp implements ${s/pat/rep}; with(m) returns the replacement for matched text m.
func replaceOp(s string, re *regexp.Regexp, with func(string) string, all, anchorStart, anchorEnd bool) string {
	n := len(s)
	if anchorStart {
		for j := n; j >= 0; j-- {
			if re.MatchString(s[:j]) {
				return with(s[:j]) + s[j:]
			}
		}
		return s
	}
	if anchorEnd {
		for i := 0; i <= n; i++ {
			if re.MatchString(s[i:]) {
				return s[:i] + with(s[i:])
			}
		}
		return s
	}
	var out strings.Builder
	i := 0
	replaced := false
	for i <= n {
		if replaced && !all {
			break
		}
		found := -1
		for j := n; j > i; j-- { // longest non-empty match starting at i
			if re.MatchString(s[i:j]) {
				found = j
				break
			}
		}
		if found < 0 {
			if i < n {
				out.WriteByte(s[i])
			}
			i++
			continue
		}
		out.WriteString(with(s[i:found]))
		i = found
		replaced = true
	}
	if i < n {
		out.WriteString(s[i:])
	}
	return out.String()
}

// ---------- parameter expansion ----------

// replPiece is one piece of a ${x/pat/rep} replacement: literal text, or the matched text.
type replPiece struct {
	text  string
	match bool
}

// replacement turns the rep part of ${x/pat/rep} into pieces, following bash >= 5.2
// (patsub_replacement, on by default): an unquoted & stands for the matched text;
// \& and quoted & ('&', "&", "$v") are a literal &; \\& is a backslash followed by
// the match. An & coming from an unquoted variable is special too, and \& in such
// a value is a literal &; other backslashes in the value are kept.
func (e *env) replacement(w *syntax.Word) []replPiece {
	var out []replPiece
	if w == nil {
		return out
	}
	var b strings.Builder
	flush := func() {
		out = append(out, replPiece{text: b.String()})
		b.Reset()
	}
	scan := func(s string, fromVar bool) {
		for i := 0; i < len(s); i++ {
			c := s[i]
			switch {
			case c == '\\' && i+1 < len(s) && s[i+1] == '&':
				b.WriteByte('&')
				i++
			case c == '\\' && fromVar:
				if i+1 < len(s) && s[i+1] == '\\' {
					bail("backslashes in a variable used as ${x/pat/rep} replacement")
				}
				b.WriteByte('\\')
			case c == '\\': // unquoted backslash escapes the next character
				if i+1 < len(s) {
					i++
					if s[i] != '\n' {
						b.WriteByte(s[i])
					}
				}
			case c == '&':
				flush()
				out = append(out, replPiece{match: true})
			default:
				b.WriteByte(c)
			}
		}
	}
	for _, p := range w.Parts {
		switch x := p.(type) {
		case *syntax.Lit:
			scan(x.Value, false)
		case *syntax.ParamExp:
			vals, _ := e.param(x)
			scan(strings.Join(vals, " "), true)
		default: // quoted text etc.: expanded as usual, & is literal there
			b.WriteString(e.expandScalar(&syntax.Word{Parts: []syntax.WordPart{p}}))
		}
	}
	flush()
	return out
}

var specialParams = regexp.MustCompile(`^([0-9@*#?$!-])$`)

// param returns the values of a parameter expansion; list=true when it
// came from ${a[@]} / ${a[*]} (one field per element when quoted).
func (e *env) param(pe *syntax.ParamExp) (vals []string, list bool) {
	if pe.Param == nil {
		bail("param without name")
	}
	name := pe.Param.Value
	if specialParams.MatchString(name) {
		bail("special parameter $%s", name)
	}
	if pe.Names != 0 {
		bail("${!prefix*}")
	}
	if pe.Excl && pe.Index == nil {
		ref, _ := e.scalar(name)
		name = ref
	}
	v, set := e.get(name)
	isAll := false
	if pe.Index != nil {
		if w, ok := pe.Index.(*syntax.Word); ok && (w.Lit() == "@" || w.Lit() == "*") {
			isAll = true
		}
	}
	if isAll {
		if set {
			vals = v.list()
		}
		list = true
	} else {
		idx := 0
		if pe.Index != nil {
			idx = e.arith(pe.Index)
			if set && idx < 0 {
				idx += v.next()
			}
		}
		if set {
			if s, ok := v.arr[idx]; ok {
				vals = []string{s}
			} else {
				set = false
			}
		}
	}
	if pe.Length {
		if isAll {
			return []string{strconv.Itoa(len(vals))}, false
		}
		s := ""
		if len(vals) > 0 {
			s = vals[0]
		}
		return []string{strconv.Itoa(len([]rune(s)))}, false
	}
	if pe.Width {
		bail("${%%...} width")
	}
	null := len(vals) == 0 || (len(vals) == 1 && vals[0] == "")
	if pe.Exp != nil {
		op := pe.Exp.Op
		switch op {
		case syntax.DefaultUnset, syntax.DefaultUnsetOrNull, syntax.AssignUnset, syntax.AssignUnsetOrNull:
			if (op == syntax.DefaultUnset || op == syntax.AssignUnset) && set || !null && (op == syntax.DefaultUnsetOrNull || op == syntax.AssignUnsetOrNull) {
				return vals, list
			}
			d := e.expandScalar(pe.Exp.Word)
			if op == syntax.AssignUnset || op == syntax.AssignUnsetOrNull {
				e.setScalar(name, d, false)
			}
			return []string{d}, false
		case syntax.AlternateUnset, syntax.AlternateUnsetOrNull:
			if op == syntax.AlternateUnset && set || op == syntax.AlternateUnsetOrNull && !null {
				return []string{e.expandScalar(pe.Exp.Word)}, false
			}
			return nil, false
		case syntax.ErrorUnset, syntax.ErrorUnsetOrNull:
			bail("${x:?}")
		}
	}
	if !set && !isAll {
		vals = []string{""}
	}
	apply := func(f func(string) string) {
		for i := range vals {
			vals[i] = f(vals[i])
		}
	}
	if pe.Slice != nil {
		if isAll {
			off := e.arith(pe.Slice.Offset)
			if off < 0 {
				off += len(vals)
			}
			if off < 0 || off > len(vals) {
				vals = nil
			} else {
				vals = vals[off:]
			}
			if pe.Slice.Length != nil {
				l := e.arith(pe.Slice.Length)
				if l < len(vals) && l >= 0 {
					vals = vals[:l]
				}
			}
		} else {
			apply(func(s string) string {
				r := []rune(s)
				off := e.arith(pe.Slice.Offset)
				if off < 0 {
					off += len(r)
				}
				if off < 0 || off > len(r) {
					return ""
				}
				r = r[off:]
				if pe.Slice.Length != nil {
					l := e.arith(pe.Slice.Length)
					if l < 0 {
						l += len(r)
					}
					if l < 0 {
						bail("negative slice length")
					}
					if l < len(r) {
						r = r[:l]
					}
				}
				return string(r)
			})
		}
	}
	if pe.Repl != nil {
		orig := pe.Repl.Orig
		anchorS, anchorE := false, false
		if orig != nil && len(orig.Parts) > 0 {
			if l, ok := orig.Parts[0].(*syntax.Lit); ok && (strings.HasPrefix(l.Value, "#") || strings.HasPrefix(l.Value, "%")) {
				anchorS, anchorE = l.Value[0] == '#', l.Value[0] == '%'
				rest := *l
				rest.Value = l.Value[1:]
				orig = &syntax.Word{Parts: append([]syntax.WordPart{&rest}, orig.Parts[1:]...)}
			}
		}
		re := e.pattern(orig)
		tmpl := e.replacement(pe.Repl.With)
		with := func(m string) string {
			var b strings.Builder
			for _, pc := range tmpl {
				if pc.match {
					b.WriteString(m)
				} else {
					b.WriteString(pc.text)
				}
			}
			return b.String()
		}
		apply(func(s string) string { return replaceOp(s, re, with, pe.Repl.All, anchorS, anchorE) })
	}
	if pe.Exp != nil {
		switch op := pe.Exp.Op; op {
		case syntax.RemSmallPrefix, syntax.RemLargePrefix, syntax.RemSmallSuffix, syntax.RemLargeSuffix:
			re := e.pattern(pe.Exp.Word)
			apply(func(s string) string {
				n := len(s)
				switch op {
				case syntax.RemSmallPrefix:
					for i := 0; i <= n; i++ {
						if re.MatchString(s[:i]) {
							return s[i:]
						}
					}
				case syntax.RemLargePrefix:
					for i := n; i >= 0; i-- {
						if re.MatchString(s[:i]) {
							return s[i:]
						}
					}
				case syntax.RemSmallSuffix:
					for i := n; i >= 0; i-- {
						if re.MatchString(s[i:]) {
							return s[:i]
						}
					}
				case syntax.RemLargeSuffix:
					for i := 0; i <= n; i++ {
						if re.MatchString(s[i:]) {
							return s[:i]
						}
					}
				}
				return s
			})
		case syntax.UpperFirst, syntax.UpperAll, syntax.LowerFirst, syntax.LowerAll:
			if pe.Exp.Word != nil && pe.Exp.Word.Lit() != "" {
				bail("case-mod with pattern")
			}
			apply(func(s string) string {
				r := []rune(s)
				if len(r) == 0 {
					return s
				}
				switch op {
				case syntax.UpperAll:
					return strings.ToUpper(s)
				case syntax.LowerAll:
					return strings.ToLower(s)
				case syntax.UpperFirst:
					r[0] = unicode.ToUpper(r[0])
				case syntax.LowerFirst:
					r[0] = unicode.ToLower(r[0])
				}
				return string(r)
			})
		case syntax.OtherParamOps:
			switch pe.Exp.Word.Lit() {
			case "L":
				apply(strings.ToLower)
			case "U":
				apply(strings.ToUpper)
			case "u":
				apply(func(s string) string {
					if r := []rune(s); len(r) > 0 {
						r[0] = unicode.ToUpper(r[0])
						return string(r)
					}
					return s
				})
			default:
				bail("${x@%s}", pe.Exp.Word.Lit())
			}
		}
	}
	return vals, list
}

// ---------- word expansion ----------

func unescapeUnquoted(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			i++
			if s[i] == '\n' {
				continue
			}
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

func unescapeDouble(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) && strings.ContainsRune("$`\"\\\n", rune(s[i+1])) {
			i++
			if s[i] == '\n' {
				continue
			}
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

func ansiC(s string) string {
	r := strings.NewReplacer(`\n`, "\n", `\t`, "\t", `\\`, `\`, `\'`, `'`, `\"`, `"`, `\e`, "\x1b", `\a`, "\a")
	return r.Replace(s)
}

type fields struct {
	out     []string
	cur     strings.Builder
	started bool
	globby  bool // an unquoted variable expansion produced glob characters
}

func (f *fields) lit(s string) { f.cur.WriteString(s); f.started = true }
func (f *fields) flush() {
	if f.started {
		f.out = append(f.out, f.cur.String())
	}
	f.cur.Reset()
	f.started = false
}
func isIFS(r rune) bool { return r == ' ' || r == '\t' || r == '\n' }
func (f *fields) split(s string) {
	if s == "" {
		return
	}
	pieces := strings.FieldsFunc(s, isIFS)
	if isIFS(rune(s[0])) {
		f.flush()
	}
	for i, p := range pieces {
		if i > 0 {
			f.flush()
		}
		f.lit(p)
	}
	if isIFS(rune(s[len(s)-1])) {
		f.flush()
	}
}

func (e *env) expandParts(parts []syntax.WordPart, f *fields, quoted, split bool) {
	for _, p := range parts {
		switch x := p.(type) {
		case *syntax.Lit:
			if quoted {
				f.lit(unescapeDouble(x.Value))
			} else {
				f.lit(unescapeUnquoted(x.Value))
			}
		case *syntax.SglQuoted:
			if x.Dollar {
				f.lit(ansiC(x.Value))
			} else {
				f.lit(x.Value)
			}
		case *syntax.DblQuoted:
			if x.Dollar {
				bail(`$"..."`)
			}
			// "${a[@]}" of an empty array yields no word at all, unlike "" or
			// "${a[*]}", which yield one empty word.
			if split && len(x.Parts) > 0 {
				onlyEmpty := true
				for _, qp := range x.Parts {
					if !e.isEmptyAtList(qp) {
						onlyEmpty = false
						break
					}
				}
				if onlyEmpty {
					continue
				}
			}
			f.started = true
			e.expandParts(x.Parts, f, true, split)
		case *syntax.ParamExp:
			vals, list := e.param(x)
			switch {
			case quoted && list && split:
				for i, v := range vals {
					if i > 0 {
						f.flush()
						f.started = true
					}
					f.lit(v)
				}
			case quoted || !split:
				f.lit(strings.Join(vals, " "))
			default:
				for i, v := range vals {
					if i > 0 {
						f.flush()
					}
					if strings.ContainsAny(v, "*?[") {
						f.globby = true
					}
					f.split(v)
				}
			}
		case *syntax.CmdSubst:
			out := e.cmdSubst(x)
			if quoted || !split {
				f.lit(out)
			} else {
				if strings.ContainsAny(out, "*?[") {
					f.globby = true
				}
				f.split(out)
			}
		case *syntax.ArithmExp:
			f.lit(strconv.Itoa(e.arith(x.X)))
		case *syntax.ProcSubst:
			bail("process substitution")
		case *syntax.ExtGlob:
			bail("extglob")
		default:
			bail("word part %T", p)
		}
	}
}

// isEmptyAtList reports whether p is a ${a[@]}-style expansion with no elements.
func (e *env) isEmptyAtList(p syntax.WordPart) bool {
	pe, ok := p.(*syntax.ParamExp)
	if !ok || pe.Index == nil || pe.Length {
		return false
	}
	if w, ok := pe.Index.(*syntax.Word); !ok || w.Lit() != "@" {
		return false
	}
	vals, list := e.param(pe)
	return list && len(vals) == 0
}

func (e *env) expandScalar(w *syntax.Word) string {
	if w == nil {
		return ""
	}
	f := &fields{}
	e.expandParts(w.Parts, f, false, false)
	return f.cur.String()
}

// expandFields: brace expansion, then expansion with word splitting (array / argument context).
func (e *env) expandFields(w *syntax.Word) []string {
	var out []string
	for _, bw := range braceWords(w) {
		f := &fields{}
		e.expandParts(bw.Parts, f, false, true)
		f.flush()
		for _, s := range f.out {
			if strings.ContainsAny(s, "*?[") && (hasUnquotedGlob(bw) || f.globby) && !urlShaped(s) {
				bail("pathname expansion of %q", s)
			}
		}
		out = append(out, f.out...)
	}
	return out
}

// urlShaped reports whether an unquoted word with glob characters is safe to keep
// as it is. Bash would match it against files, e.g. source=(https://x.org/f?id=1)
// is a pattern for a directory named "https:" in the PKGBUILD's directory. If the
// first path component looks like a URL scheme ("https:", "git+https:",
// "name::https:") and has no glob characters, nothing can match, and bash keeps
// the word unchanged (no nullglob/failglob in makepkg). Other patterns, like
// *.patch, depend on which files exist, so they are still refused.
func urlShaped(s string) bool {
	i := strings.IndexByte(s, '/')
	return i > 0 && strings.HasSuffix(s[:i], ":") && !strings.ContainsAny(s[:i], "*?[")
}

func hasUnquotedGlob(w *syntax.Word) bool {
	for _, p := range w.Parts {
		if l, ok := p.(*syntax.Lit); ok && strings.ContainsAny(strings.ReplaceAll(l.Value, `\`, ""), "*?[") {
			return true
		}
	}
	return false
}

func braceWords(w *syntax.Word) []*syntax.Word {
	cp := &syntax.Word{Parts: append([]syntax.WordPart(nil), w.Parts...)}
	if !syntax.SplitBraces(cp) {
		return []*syntax.Word{w}
	}
	return expandBraceParts(cp.Parts)
}

func expandBraceParts(parts []syntax.WordPart) []*syntax.Word {
	words := []*syntax.Word{{}}
	for _, p := range parts {
		be, ok := p.(*syntax.BraceExp)
		if !ok {
			for _, w := range words {
				w.Parts = append(w.Parts, p)
			}
			continue
		}
		var alts []*syntax.Word
		if be.Sequence {
			alts = braceSeq(be)
		} else {
			for _, el := range be.Elems {
				alts = append(alts, expandBraceParts(el.Parts)...)
			}
		}
		var next []*syntax.Word
		for _, w := range words {
			for _, a := range alts {
				nw := &syntax.Word{Parts: append(append([]syntax.WordPart(nil), w.Parts...), a.Parts...)}
				next = append(next, nw)
			}
		}
		words = next
	}
	return words
}

func braceSeq(be *syntax.BraceExp) []*syntax.Word {
	a, b := be.Elems[0].Lit(), be.Elems[1].Lit()
	step := 1
	if len(be.Elems) > 2 {
		step, _ = strconv.Atoi(be.Elems[2].Lit())
		if step < 0 {
			step = -step
		}
		if step == 0 {
			step = 1
		}
	}
	lit := func(s string) *syntax.Word { return &syntax.Word{Parts: []syntax.WordPart{&syntax.Lit{Value: s}}} }
	var out []*syntax.Word
	x, err1 := strconv.Atoi(a)
	y, err2 := strconv.Atoi(b)
	if err1 == nil && err2 == nil {
		width := 0
		if (len(a) > 1 && a[0] == '0') || (len(b) > 1 && b[0] == '0') {
			width = max(len(a), len(b))
		}
		for i := x; ; {
			out = append(out, lit(fmt.Sprintf("%0*d", width, i)))
			if i == y {
				break
			}
			if x < y {
				i += step
				if i > y {
					break
				}
			} else {
				i -= step
				if i < y {
					break
				}
			}
		}
		return out
	}
	if len(a) == 1 && len(b) == 1 {
		for c := a[0]; ; {
			out = append(out, lit(string(c)))
			if c == b[0] {
				break
			}
			if a[0] < b[0] {
				c += byte(step)
				if c > b[0] {
					break
				}
			} else {
				c -= byte(step)
				if c < b[0] {
					break
				}
			}
		}
		return out
	}
	bail("brace sequence %s..%s", a, b)
	return nil
}

// ---------- command substitution ----------

// cmdSubst evaluates $(...) for the pipelines PKGBUILDs use to reformat
// versions: echo ... | cut ..., echo ... | sed 's/.../.../'. Only the options
// those use are supported; anything else is refused.
func (e *env) cmdSubst(cs *syntax.CmdSubst) string {
	var cmds []*syntax.CallExpr
	var names []string
	ok := len(cs.Stmts) == 1
	var collect func(st *syntax.Stmt)
	collect = func(st *syntax.Stmt) {
		if st.Negated || st.Background || st.Coprocess || len(st.Redirs) > 0 {
			ok = false
		}
		switch x := st.Cmd.(type) {
		case *syntax.CallExpr:
			if len(x.Assigns) > 0 || len(x.Args) == 0 {
				ok = false
				return
			}
			cmds = append(cmds, x)
			name := x.Args[0].Lit()
			if name == "" {
				name = "?"
			}
			names = append(names, name)
		case *syntax.BinaryCmd:
			if x.Op != syntax.Pipe {
				ok = false
			}
			collect(x.X)
			collect(x.Y)
		default:
			ok = false
		}
	}
	if ok {
		collect(cs.Stmts[0])
	}
	if !ok || len(cmds) == 0 || names[0] != "echo" {
		if len(names) > 0 {
			bail("command substitution $(%s)", strings.Join(names, " | "))
		}
		bail("command substitution $(...)")
	}
	text := ""
	for i, c := range cmds {
		var args []string
		for _, w := range c.Args[1:] {
			args = append(args, e.expandFields(w)...)
		}
		switch {
		case i == 0: // echo
			if len(args) > 0 && regexp.MustCompile(`^-[neE]+$`).MatchString(args[0]) {
				bail("command substitution $(echo %s ...)", args[0])
			}
			text = strings.Join(args, " ") + "\n"
		case names[i] == "cut":
			text = cut(args, text)
		case names[i] == "sed":
			text = sed(args, text)
		default:
			bail("command substitution $(%s)", strings.Join(names, " | "))
		}
	}
	return strings.TrimRight(text, "\n") // $(...) drops trailing newlines
}

// lines splits command output into lines (a final newline ends the last line).
func lines(text string) []string {
	if text == "" {
		return nil
	}
	return strings.Split(strings.TrimSuffix(text, "\n"), "\n")
}

// cut supports -d CHAR and -f LIST (N, N-M, N-, -M, comma-separated).
func cut(args []string, text string) string {
	delim, list := "\t", ""
	for i := 0; i < len(args); i++ {
		a := args[i]
		opt, val := a, ""
		if len(a) > 2 && strings.HasPrefix(a, "-") {
			opt, val = a[:2], a[2:]
		} else if i+1 < len(args) {
			val = args[i+1]
			i++
		}
		switch opt {
		case "-d":
			delim = val
		case "-f":
			list = val
		default:
			bail("command substitution: cut %s", a)
		}
	}
	if len(delim) != 1 || list == "" {
		bail("command substitution: cut -d %q -f %q", delim, list)
	}
	var ranges [][2]int // 1-based, inclusive; 0 = open end
	for _, r := range strings.Split(list, ",") {
		lo, hi, isRange := strings.Cut(r, "-")
		a, errA := strconv.Atoi(lo)
		b, errB := strconv.Atoi(hi)
		switch {
		case !isRange && errA == nil && a > 0:
			ranges = append(ranges, [2]int{a, a})
		case isRange && lo == "" && errB == nil && b > 0:
			ranges = append(ranges, [2]int{1, b})
		case isRange && hi == "" && errA == nil && a > 0:
			ranges = append(ranges, [2]int{a, 0})
		case isRange && errA == nil && errB == nil && a > 0 && b >= a:
			ranges = append(ranges, [2]int{a, b})
		default:
			bail("command substitution: cut -f %s", list)
		}
	}
	var out strings.Builder
	for _, line := range lines(text) {
		if !strings.Contains(line, delim) { // no delimiter: line printed as is
			out.WriteString(line + "\n")
			continue
		}
		fields := strings.Split(line, delim)
		var picked []string
		for i, fld := range fields { // in field order, each field once
			for _, r := range ranges {
				if i+1 >= r[0] && (r[1] == 0 || i+1 <= r[1]) {
					picked = append(picked, fld)
					break
				}
			}
		}
		out.WriteString(strings.Join(picked, delim) + "\n")
	}
	return out.String()
}

// sedSubst is one s/regex/replacement/[g] command.
type sedSubst struct {
	re     *regexp.Regexp
	repl   string
	global bool
}

// sed supports scripts made of s commands: sed SCRIPT, or sed -e SCRIPT [-e SCRIPT]...
func sed(args []string, text string) string {
	var scripts []string
	for i := 0; i < len(args); i++ {
		switch {
		case args[i] == "-e" && i+1 < len(args):
			scripts = append(scripts, args[i+1])
			i++
		case strings.HasPrefix(args[i], "-"):
			bail("command substitution: sed %s", args[i])
		case len(scripts) == 0 && i == len(args)-1:
			scripts = append(scripts, args[i])
		default:
			bail("command substitution: sed with input files")
		}
	}
	if len(scripts) == 0 {
		bail("command substitution: sed without a script")
	}
	var cmds []sedSubst
	for _, sc := range scripts {
		cmds = append(cmds, parseSedScript(sc)...)
	}
	var out strings.Builder
	for _, line := range lines(text) {
		for _, c := range cmds {
			line = c.apply(line)
		}
		out.WriteString(line + "\n")
	}
	return out.String()
}

func parseSedScript(sc string) []sedSubst {
	var cmds []sedSubst
	i := 0
	for i < len(sc) {
		switch c := sc[i]; {
		case c == ';' || c == '\n' || c == ' ' || c == '\t':
			i++
			continue
		case c != 's' || i+1 >= len(sc):
			bail("command substitution: sed command %q", sc[i:])
		}
		delim := sc[i+1]
		if delim == '\\' || delim == '\n' {
			bail("command substitution: sed delimiter %q", delim)
		}
		i += 2
		part := func() string { // up to the next unescaped delimiter
			start := i
			for i < len(sc) && sc[i] != delim {
				if sc[i] == '\\' {
					i++
				}
				i++
			}
			if i >= len(sc) {
				bail("command substitution: unterminated sed s command")
			}
			p := sc[start:i]
			i++
			return p
		}
		re, repl := part(), part()
		global := false
		for i < len(sc) && sc[i] != ';' && sc[i] != '\n' {
			if sc[i] != 'g' {
				bail("command substitution: sed flag %q", sc[i])
			}
			global = true
			i++
		}
		rx, err := regexp.Compile(breToRE2(re, delim))
		if err != nil {
			bail("command substitution: sed regex %q", re)
		}
		rx.Longest() // POSIX leftmost-longest matching, like sed
		cmds = append(cmds, sedSubst{rx, unescapeDelim(repl, delim), global})
	}
	return cmds
}

// unescapeDelim turns \<delim> into <delim> in a replacement.
func unescapeDelim(s string, delim byte) string {
	return strings.ReplaceAll(s, "\\"+string(delim), string(delim))
}

// breToRE2 translates a POSIX basic regular expression (with the GNU \+ \?
// \| extensions) into Go regexp syntax. Back-references are refused.
func breToRE2(bre string, delim byte) string {
	var b strings.Builder
	atStart := true // where ^ is an anchor and * is literal
	for i := 0; i < len(bre); i++ {
		c := bre[i]
		start := atStart
		atStart = false
		switch {
		case c == '\\' && i+1 < len(bre):
			i++
			d := bre[i]
			switch {
			case d == '(':
				b.WriteString("(")
				atStart = true
			case d == ')' || d == '{' || d == '}' || d == '+' || d == '?':
				b.WriteByte(d)
			case d == '|':
				b.WriteByte('|')
				atStart = true
			case d >= '1' && d <= '9':
				bail("command substitution: sed back-reference \\%c", d)
			case d == 'n':
				b.WriteString(`\n`)
			case d == delim || strings.IndexByte(`.*[]^$\/`, d) >= 0:
				b.WriteString(regexp.QuoteMeta(string(d)))
			default:
				bail("command substitution: sed escape \\%c", d)
			}
		case c == '[':
			j := i + 1
			if j < len(bre) && bre[j] == '^' {
				j++
			}
			if j < len(bre) && bre[j] == ']' {
				j++
			}
			for j < len(bre) && bre[j] != ']' {
				if bre[j] == '[' && j+1 < len(bre) && (bre[j+1] == ':' || bre[j+1] == '.' || bre[j+1] == '=') {
					if k := strings.Index(bre[j+2:], string(bre[j+1])+"]"); k >= 0 {
						j += k + 4
						continue
					}
				}
				j++
			}
			if j >= len(bre) {
				bail("command substitution: sed bracket expression")
			}
			b.WriteString(strings.ReplaceAll(bre[i:j+1], `\`, `\\`)) // backslash is literal in brackets
			i = j
		case c == '*' && start:
			b.WriteString(`\*`)
		case c == '^' && start:
			b.WriteByte('^')
			atStart = true
		case c == '$' && (i == len(bre)-1 || strings.HasPrefix(bre[i+1:], `\)`) || strings.HasPrefix(bre[i+1:], `\|`)):
			b.WriteByte('$')
		case c == '.' || c == '*':
			b.WriteByte(c)
		default: // everything else, including ( ) { } + ? | ^ $, is literal in a BRE
			b.WriteString(regexp.QuoteMeta(string(c)))
		}
	}
	return b.String()
}

// apply runs the substitution on one line: & is the match, \1-\9 groups,
// \& and \\ literal, \n a newline.
func (c sedSubst) apply(line string) string {
	expand := func(m []int) string {
		var b strings.Builder
		for i := 0; i < len(c.repl); i++ {
			ch := c.repl[i]
			switch {
			case ch == '&':
				b.WriteString(line[m[0]:m[1]])
			case ch == '\\' && i+1 < len(c.repl):
				i++
				d := c.repl[i]
				switch {
				case d >= '0' && d <= '9':
					g := int(d - '0')
					if 2*g+1 >= len(m) {
						bail("command substitution: sed group \\%c doesn't exist", d)
					}
					if m[2*g] >= 0 {
						b.WriteString(line[m[2*g]:m[2*g+1]])
					}
				case d == 'n':
					b.WriteByte('\n')
				default:
					b.WriteByte(d)
				}
			default:
				b.WriteByte(ch)
			}
		}
		return b.String()
	}
	n := 1
	if c.global {
		n = -1
	}
	matches := c.re.FindAllStringSubmatchIndex(line, n)
	if matches == nil {
		return line
	}
	var b strings.Builder
	last := 0
	for _, m := range matches {
		b.WriteString(line[last:m[0]])
		b.WriteString(expand(m))
		last = m[1]
	}
	b.WriteString(line[last:])
	return b.String()
}

// ---------- statements ----------

func (e *env) assign(a *syntax.Assign) {
	if a.Name == nil {
		bail("assign without name")
	}
	name := a.Name.Value
	if a.Array != nil {
		// Build into a fresh variable and install it at the end: the right-hand
		// side may refer to the old value (provides=("${provides[@]}")).
		v := &variable{arr: map[int]string{}, array: true}
		if old, ok := e.get(name); ok && a.Append {
			for k, s := range old.arr {
				v.arr[k] = s
			}
		}
		defer func() { e.vars[name] = v }()
		next := v.next()
		for _, el := range a.Array.Elems {
			if el.Index != nil {
				i := e.arith(el.Index)
				v.arr[i] = e.expandScalar(el.Value)
				next = i + 1
				continue
			}
			if el.Value == nil {
				continue
			}
			for _, s := range e.expandFields(el.Value) {
				v.arr[next] = s
				next++
			}
		}
		return
	}
	val := e.expandScalar(a.Value)
	if a.Index != nil {
		v, ok := e.get(name)
		if !ok {
			v = &variable{arr: map[int]string{}, array: true}
			e.vars[name] = v
		}
		i := e.arith(a.Index)
		if a.Append {
			val = v.arr[i] + val
		}
		v.arr[i] = val
		return
	}
	if a.Naked {
		return
	}
	e.setScalar(name, val, a.Append)
}

func (e *env) runTop(f *syntax.File) {
	// Functions defined anywhere in the file: calling one at top level could change anything.
	defined := map[string]bool{}
	syntax.Walk(f, func(n syntax.Node) bool {
		if fd, ok := n.(*syntax.FuncDecl); ok {
			defined[fd.Name.Value] = true
		}
		return true
	})
	for _, st := range f.Stmts {
		e.runStmt(st, defined)
	}
}

// tryAssign evaluates a top-level assignment. If the value can't be computed,
// the variable is marked unknown instead of giving up right away.
func (e *env) tryAssign(a *syntax.Assign) {
	if a.Name == nil {
		bail("assign without name")
	}
	if a.Name.Value == "IFS" {
		bail("IFS assignment") // word splitting always uses the default IFS
	}
	defer func() {
		if r := recover(); r != nil {
			u, ok := r.(unsupported)
			if !ok {
				panic(r)
			}
			e.poison(a.Name.Value, u.root)
		}
	}()
	e.assign(a)
}

func (e *env) runStmt(st *syntax.Stmt, defined map[string]bool) {
	if !st.Negated && !st.Background && !st.Coprocess && len(st.Redirs) == 0 {
		switch c := st.Cmd.(type) {
		case nil:
			return
		case *syntax.FuncDecl:
			e.funcs[c.Name.Value] = c
			return
		case *syntax.CallExpr:
			if len(c.Args) == 0 {
				for _, a := range c.Assigns {
					e.tryAssign(a)
				}
				return
			}
			if len(c.Assigns) == 0 {
				switch c.Args[0].Lit() {
				case ":", "true":
					return
				case "unset":
					for _, a := range c.Args[1:] {
						n := a.Lit()
						if n == "" {
							bail("unset with a dynamic name")
						}
						if strings.HasPrefix(n, "-") {
							if n == "-f" {
								bail("unset -f")
							}
							continue
						}
						delete(e.vars, n)
					}
					return
				}
			}
		case *syntax.DeclClause:
			flags := ""
			for _, a := range c.Args {
				if a.Name == nil {
					l := a.Value.Lit()
					if !strings.HasPrefix(l, "-") && !strings.HasPrefix(l, "+") {
						bail("%s with a dynamic name", c.Variant.Value)
					}
					flags += l
				}
			}
			switch {
			case strings.Contains(flags, "n"):
				bail("%s %s (nameref)", c.Variant.Value, flags)
			case strings.ContainsAny(flags, "Ailu"):
				// associative, integer, lower/upper-case attributes: not evaluated
				for _, a := range c.Args {
					if a.Name != nil {
						e.poison(a.Name.Value, c.Variant.Value+" "+flags)
					}
				}
			case strings.ContainsAny(flags, "fFp"):
				// only prints or exports functions
			default:
				for _, a := range c.Args {
					if a.Name != nil {
						e.tryAssign(a)
					}
				}
			}
			return
		}
	}
	e.skip(st, defined)
}

// Builtins that could change the shell's state in ways skip can't follow.
var unsafeBuiltins = map[string]bool{
	"eval": true, "source": true, ".": true, "exec": true, "exit": true, "return": true,
	"set": true, "alias": true, "unalias": true, "trap": true, "enable": true, "builtin": true,
	"command": true, "shift": true, "getopts": true, "mapfile": true, "readarray": true,
}

// skip handles a top-level statement that isn't evaluated (if, for, case, &&,
// a command, ...). It works out from the code which variables the statement
// could set and marks them unknown. Statements that could change anything
// (eval, source, calling one of the file's functions, exit, ...) still give up.
func (e *env) skip(st *syntax.Stmt, defined map[string]bool) {
	what := describe(st)
	writes := map[string]bool{}
	temporary := map[*syntax.Assign]bool{} // FOO=x cmd: only set for cmd
	arithTarget := func(x syntax.ArithmExpr) {
		w, ok := x.(*syntax.Word)
		name := ""
		if ok {
			name, _, _ = strings.Cut(w.Lit(), "[")
		}
		if name == "" {
			bail("%s: arithmetic assignment to a dynamic name", what)
		}
		writes[name] = true
	}
	syntax.Walk(st, func(n syntax.Node) bool {
		switch x := n.(type) {
		case *syntax.Assign:
			if x.Name == nil { // a declare/local/export flag, or a dynamic name
				if l := x.Value.Lit(); !strings.HasPrefix(l, "-") && !strings.HasPrefix(l, "+") {
					bail("%s: declare with a dynamic name", what)
				}
				return true
			}
			if x.Name.Value == "IFS" && !temporary[x] {
				bail("%s changes IFS", what)
			}
			writes[x.Name.Value] = true
		case *syntax.BinaryArithm:
			switch x.Op {
			case syntax.Assgn, syntax.AddAssgn, syntax.SubAssgn, syntax.MulAssgn, syntax.QuoAssgn,
				syntax.RemAssgn, syntax.AndAssgn, syntax.OrAssgn, syntax.XorAssgn, syntax.ShlAssgn, syntax.ShrAssgn:
				arithTarget(x.X)
			}
		case *syntax.UnaryArithm:
			if x.Op == syntax.Inc || x.Op == syntax.Dec {
				arithTarget(x.X)
			}
		case *syntax.DeclClause:
			for _, a := range x.Args {
				if a.Name != nil {
					continue
				}
				if l := a.Value.Lit(); strings.HasPrefix(l, "-") && strings.Contains(l, "n") {
					bail("%s: %s %s (nameref)", what, x.Variant.Value, l)
				}
			}
		case *syntax.WordIter:
			writes[x.Name.Value] = true
		case *syntax.FuncDecl:
			bail("%s defines a function", what)
		case *syntax.CoprocClause:
			bail("%s: coproc", what)
		case *syntax.TestClause:
			writes["BASH_REMATCH"] = true
		case *syntax.Redirect:
			if x.N != nil && strings.HasPrefix(x.N.Value, "{") {
				bail("%s: {var}> redirection", what)
			}
		case *syntax.CallExpr:
			if len(x.Args) == 0 {
				return true
			}
			for _, a := range x.Assigns {
				temporary[a] = true
			}
			name := x.Args[0].Lit()
			switch {
			case name == "":
				bail("%s: command with a dynamic name", what)
			case defined[name]:
				if x == st.Cmd {
					bail("top-level call to function %s()", name)
				}
				bail("%s calls function %s()", what, name)
			case unsafeBuiltins[name]:
				bail("%s: %s", what, name)
			case name == "shopt":
				if len(x.Args) != 3 || x.Args[2].Lit() != "extglob" {
					bail("%s: shopt", what)
				}
			case name == "printf":
				for _, a := range x.Args[1:] {
					if a.Lit() == "-v" {
						bail("%s: printf -v", what)
					}
				}
			case name == "read":
				// read [-r] [-a NAME] [-d DELIM] ... [NAME...]: the names get set
				args, named := x.Args[1:], false
				for i := 0; i < len(args); i++ {
					l := args[i].Lit()
					if l == "" {
						bail("%s: read with a dynamic argument", what)
					}
					if !strings.HasPrefix(l, "-") || l == "-" {
						writes[l], named = true, true
						continue
					}
					for j := 1; j < len(l); j++ {
						if !strings.ContainsRune("adinNptu", rune(l[j])) {
							continue
						}
						optarg := l[j+1:]
						if optarg == "" {
							if i++; i < len(args) {
								optarg = args[i].Lit()
							}
						}
						if l[j] == 'a' {
							if optarg == "" {
								bail("%s: read -a with a dynamic name", what)
							}
							writes[optarg], named = true, true
						}
						break
					}
				}
				if !named {
					writes["REPLY"] = true
				}
			case name == "unset":
				for _, a := range x.Args[1:] {
					l := a.Lit()
					if l == "" || l == "-f" {
						bail("%s: unset", what)
					}
					if !strings.HasPrefix(l, "-") {
						writes[l] = true
					}
				}
			}
		}
		return true
	})
	for name := range writes {
		e.poison(name, what)
	}
}

// describe names a top-level statement for messages.
func describe(st *syntax.Stmt) string {
	switch c := st.Cmd.(type) {
	case *syntax.IfClause:
		return "top-level if"
	case *syntax.ForClause:
		return "top-level for loop"
	case *syntax.WhileClause:
		if c.Until {
			return "top-level until loop"
		}
		return "top-level while loop"
	case *syntax.CaseClause:
		return "top-level case"
	case *syntax.BinaryCmd:
		return "top-level && or || or pipe"
	case *syntax.Block, *syntax.Subshell:
		return "top-level { } or ( )"
	case *syntax.TestClause:
		return "top-level [[ ]]"
	case *syntax.DeclClause:
		return "top-level " + c.Variant.Value
	case *syntax.CallExpr:
		if len(c.Args) > 0 && c.Args[0].Lit() != "" {
			return "top-level command " + c.Args[0].Lit()
		}
		return "top-level command"
	}
	return "top-level " + strings.TrimPrefix(fmt.Sprintf("%T", st.Cmd), "*syntax.")
}

// ---------- .SRCINFO output (mirrors libmakepkg/srcinfo.sh) ----------

var hashAlgos = []string{"ck", "md5", "sha1", "sha224", "sha256", "sha384", "sha512", "b2"}

func sums() []string {
	var s []string
	for _, h := range hashAlgos {
		s = append(s, h+"sums")
	}
	return s
}

var (
	globalSingle = []string{"pkgdesc", "pkgver", "pkgrel", "epoch", "url", "install", "changelog"}
	globalMulti  = append([]string{"arch", "groups", "license", "checkdepends", "makedepends", "depends", "optdepends", "provides", "conflicts", "replaces", "noextract", "options", "backup", "source", "validpgpkeys"}, sums()...)
	pkgSingle    = []string{"pkgdesc", "url", "install", "changelog"}
	pkgMulti     = []string{"arch", "groups", "license", "checkdepends", "depends", "optdepends", "provides", "conflicts", "replaces", "options", "backup"}
	archMulti    = append([]string{"source", "provides", "conflicts", "depends", "replaces", "optdepends", "makedepends", "checkdepends"}, sums()...)
	wsRun        = regexp.MustCompile(`[ \t\n\r\f\v]+`)
)

type section struct {
	Head string
	KV   map[string][]string
	Keys []string // keys in output order
}

func (s *section) add(k, v string) {
	if _, ok := s.KV[k]; !ok {
		s.Keys = append(s.Keys, k)
	}
	s.KV[k] = append(s.KV[k], v)
}

func norm(s string) string {
	s = wsRun.ReplaceAllString(s, " ")
	s = strings.TrimPrefix(s, " ")
	return strings.TrimSuffix(s, " ")
}

func (e *env) globalAttr(name string, multi bool) ([]string, bool) {
	v, ok := e.get(name)
	if !ok {
		return nil, false
	}
	if multi {
		l := v.list()
		return l, len(l) > 0
	}
	s := v.arr[0]
	return []string{s}, s != ""
}

// overrides emulates extract_function_variable: every statement-level
// assignment to the attribute inside package_<name>() (or package()) is
// evaluated, in order, in the global scope.
func (e *env) override(pkg, attr string, multi bool) ([]string, bool) {
	fd := e.funcs["package_"+pkg]
	if fd == nil {
		fd = e.funcs["package"]
	}
	if fd == nil {
		return nil, false
	}
	var found bool
	var result []string
	syntax.Walk(fd.Body, func(n syntax.Node) bool {
		switch x := n.(type) {
		case *syntax.CmdSubst, *syntax.ProcSubst, *syntax.BinaryCmd:
			// `a && depends=(x)` prints on one line in `declare -f`, so
			// makepkg's line-anchored grep never sees it.
			return false
		case *syntax.CallExpr:
			if len(x.Args) != 0 || len(x.Assigns) == 0 {
				return true
			}
			a := x.Assigns[0]
			if a.Name == nil || a.Name.Value != attr || (a.Array != nil) != multi || a.Index != nil {
				return true
			}
			out := &variable{arr: map[int]string{}}
			if a.Append {
				if !found { // makepkg seeds the output with the global value first
					result, _ = e.globalAttr(attr, multi)
				}
				for i, v := range result {
					out.arr[i] = v
				}
			}
			saved, had := e.vars["__out"]
			e.vars["__out"] = out
			// get_pkgbuild_attribute has `local pkgname=$1`, which shadows the
			// global array while the override is eval'd.
			savedPkg := e.vars["pkgname"]
			e.vars["pkgname"] = &variable{arr: map[int]string{0: pkg}}
			b := *a
			b.Name = &syntax.Lit{Value: "__out"}
			e.assign(&b)
			result = e.vars["__out"].list()
			e.vars["pkgname"] = savedPkg
			if had {
				e.vars["__out"] = saved
			} else {
				delete(e.vars, "__out")
			}
			found = true
		}
		return true
	})
	return result, found
}

func (e *env) srcinfo() []section {
	var secs []section
	pkgnames, _ := e.globalAttr("pkgname", true)
	if len(pkgnames) == 0 {
		bail("no pkgname")
	}
	base, ok := e.scalar("pkgbase")
	if !ok || base == "" {
		base = pkgnames[0]
		e.setScalar("pkgbase", base, false) // makepkg: pkgbase=${pkgbase:-${pkgname[0]}}
	}
	write := func(sec *section, k string, vals []string) {
		if len(vals) == 0 {
			// makepkg prints "key = " for an override that empties an array
			sec.add(k, "")
		}
		for _, v := range vals {
			sec.add(k, norm(v))
		}
	}
	g := section{Head: "pkgbase = " + base, KV: map[string][]string{}}
	for _, k := range globalSingle {
		if v, ok := e.globalAttr(k, false); ok {
			write(&g, k, v)
		}
	}
	for _, k := range globalMulti {
		if v, ok := e.globalAttr(k, true); ok {
			write(&g, k, v)
		}
	}
	arches, _ := e.globalAttr("arch", true)
	for _, a := range arches {
		if a == "any" {
			continue
		}
		for _, k := range archMulti {
			if v, ok := e.globalAttr(k+"_"+a, true); ok {
				write(&g, k+"_"+a, v)
			}
		}
	}
	secs = append(secs, g)
	for _, p := range pkgnames {
		s := section{Head: "pkgname = " + p, KV: map[string][]string{}}
		for _, k := range pkgSingle {
			if v, ok := e.override(p, k, false); ok {
				write(&s, k, v)
			}
		}
		for _, k := range pkgMulti {
			if v, ok := e.override(p, k, true); ok {
				write(&s, k, v)
			}
		}
		parch := arches
		if v, ok := e.override(p, "arch", true); ok {
			parch = v
		}
		for _, a := range parch {
			if a == "any" {
				continue
			}
			for _, k := range archMulti {
				if v, ok := e.override(p, k+"_"+a, true); ok {
					write(&s, k+"_"+a, v)
				}
			}
		}
		secs = append(secs, s)
	}
	return secs
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: minieval PKGBUILD|DIR   (prints the .SRCINFO makepkg would generate)")
		os.Exit(2)
	}
	path := os.Args[1]
	if fi, err := os.Stat(path); err == nil && fi.IsDir() {
		path = filepath.Join(path, "PKGBUILD")
	}
	defer func() {
		if r := recover(); r != nil {
			if u, ok := r.(unsupported); ok {
				fmt.Fprintf(os.Stderr, "unsupported: %s\n", u.why)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "crash: %v\n", r)
			os.Exit(4)
		}
	}()
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
	f, err := syntax.NewParser(syntax.Variant(syntax.LangBash)).Parse(strings.NewReader(string(src)), path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse-error: %v\n", err)
		os.Exit(3)
	}
	e := &env{vars: map[string]*variable{}, funcs: map[string]*syntax.FuncDecl{}}
	e.setScalar("CARCH", "x86_64", false)
	e.runTop(f)
	// Everything is evaluated before anything is printed, so a refusal never
	// leaves partial output.
	fmt.Print(render(e.srcinfo()))
}

// render produces the .SRCINFO text in the same layout as `makepkg --printsrcinfo`.
func render(secs []section) string {
	var b strings.Builder
	for i, s := range secs {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(s.Head + "\n")
		for _, k := range s.Keys {
			for _, v := range s.KV[k] {
				fmt.Fprintf(&b, "\t%s = %s\n", k, v)
			}
		}
	}
	return b.String()
}
