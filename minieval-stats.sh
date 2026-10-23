#!/usr/bin/env bash
#
# For every <name>/PKGBUILD in a monorepo directory, let minieval print a
# .SRCINFO and compare it byte for byte (cmp) with <name>/.SRCINFO on disk.
# minieval itself never sees the .SRCINFO, so it can't report its own success.
#
# Usage: minieval-stats.sh MONOREPO_DIR [path/to/minieval]   (default: ./minieval)
#
# Results:
#   ./minieval-results.tsv     name <TAB> status <TAB> detail, one line per package
#   ./minieval-diffs/<name>.diff   diff -u (.SRCINFO on disk vs minieval output) for each mismatch
#
# Status: match        output identical to the .SRCINFO
#         mismatch     output differs (see the diff)
#         unsupported  minieval refused the PKGBUILD (detail: its reason)
#         parse-error  minieval could not parse it
#         crash        minieval failed unexpectedly (a bug)
#         no-srcinfo   nothing to compare with

set -u

dir=${1:?usage: ${0##*/} MONOREPO_DIR [path/to/minieval]}
bin=${2:-./minieval}
results=./minieval-results.tsv
diffs=./minieval-diffs
[[ -x $bin ]] || { echo "minieval binary not found at $bin" >&2; exit 1; }

out=$(mktemp) err=$(mktemp)
trap 'rm -f "$out" "$err"' EXIT
rm -rf "$diffs"
mkdir -p "$diffs"

# 1. Run minieval on every PKGBUILD and compare.
shopt -s nullglob
files=("$dir"/*/PKGBUILD)
echo "Evaluating ${#files[@]} PKGBUILDs..." >&2
for f in "${files[@]}"; do
	pkgdir=${f%/PKGBUILD}
	name=${pkgdir##*/}
	srcinfo=$pkgdir/.SRCINFO
	"$bin" "$f" >"$out" 2>"$err"
	rc=$?
	detail=''
	read -r detail <"$err"
	if ((rc == 0)); then
		if [[ ! -f $srcinfo ]]; then
			status=no-srcinfo detail=''
		elif cmp -s "$out" "$srcinfo"; then
			status=match detail=''
		else
			status=mismatch detail="$diffs/$name.diff"
			diff -u --label "$name/.SRCINFO" --label "minieval $name/PKGBUILD" "$srcinfo" "$out" >"$detail"
		fi
	elif ((rc == 1)); then
		status=unsupported detail=${detail#unsupported: }
	elif ((rc == 3)); then
		status=parse-error detail=${detail#parse-error: }
	else
		status=crash
	fi
	printf '%s\t%s\t%s\n' "$name" "$status" "$detail"
done >"$results"

# 2. Counts and percentages per status.
total=$(($(wc -l <"$results")))
echo
echo "Results for $total PKGBUILDs:"
cut -f2 "$results" | sort | uniq -c | sort -rn |
	awk -v t="$total" '{ printf "  %-12s %6d  %5.1f%%\n", $2, $1, 100 * $1 / t }'

matched=$(cut -f2 "$results" | grep -c '^match$')
mismatched=$(cut -f2 "$results" | grep -c '^mismatch$')
if ((total)); then
	# Two different totals: how many files minieval gave an answer for at all
	# (it refuses "unsupported" ones rather than guess), and how many of those
	# answers are identical to the .SRCINFO.
	awk -v m="$matched" -v x="$mismatched" -v t="$total" 'BEGIN {
		a = m + x
		printf "\nAnswered: %d of %d files (%.1f%%); the rest were refused or had no .SRCINFO.\n", a, t, 100 * a / t
		if (a) printf "Identical: %d of those %d answers (%.2f%%); %d differ.\n", m, a, 100 * m / a, x }'
fi

# 3. Why files were refused, most common first (some reasons shortened).
if awk -F'\t' '$2 == "unsupported" { found = 1 } END { exit !found }' "$results"; then
	echo
	echo "Why files were refused (unsupported):"
	awk -F'\t' '$2 == "unsupported" { print $3 }' "$results" |
		sed -E -e 's/^pathname expansion of .*/pathname expansion (unquoted glob)/' \
			-e 's/^top-level ForClause$/top-level for loop/' -e 's/^top-level WhileClause$/top-level while loop/' \
			-e 's/^top-level IfClause$/top-level if/' -e 's/^top-level CaseClause$/top-level case/' \
			-e 's/^top-level BinaryCmd$/top-level \&\& or || or pipe/' -e 's/^top-level (Block|Subshell)$/top-level { } or ( )/' |
		sort | uniq -c | sort -rn | head -n 15 | sed 's/^/  /'
fi

# 4. Files that need a look.
for s in mismatch parse-error crash; do
	n=$(cut -f2 "$results" | grep -c "^$s\$")
	if ((n)); then
		echo
		echo "Files with status $s:"
		awk -F'\t' -v s="$s" '$2 == s { print "  " $1 ($3 != "" ? "  (" $3 ")" : "") }' "$results" | head -n 20
		((n > 20)) && echo "  ... and $((n - 20)) more (all in $results)"
	fi
done
