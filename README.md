# PKGBUILD interpreter PoC

> [!WARNING]
> This is just an experimental prototype.
> Do not use it for anything serious!

## Compiling the interpreter

```bash
go build -o minieval pkgbuild_minieval.go
```

The interpreter takes a single `PKGBUILD` argument and prints a computed `.SRCINFO` file contents that then can be compared with the original.

```bash
$ ./minieval monorepo/z/PKGBUILD
pkgbase = z
	pkgdesc = Tracks your most used directories, based on 'frecency'
	pkgver = 1.12
	pkgrel = 2
	url = https://github.com/rupa/z/
	install = z.install
	arch = any
	license = WTFPL
	depends = bash
	source = z-1.12.tar.gz::https://codeload.github.com/rupa/z/tar.gz/v1.12
	sha256sums = 7d8695f2f5af6805f0db231e6ed571899b8b375936a8bfca81a522b7082b574e

pkgname = z
```

## Sample results

To check if the interpreter produces values similar to official files a small wrapper script is attached:

```
> ./minieval-stats.sh monorepo/
Evaluating 12394 PKGBUILDs...

Results for 12394 PKGBUILDs:
  match         12233   98.7%
  unsupported      75    0.6%
  no-srcinfo       57    0.5%
  parse-error      18    0.1%
  mismatch         11    0.1%

Answered: 12244 of 12394 files (98.8%); the rest were refused or had no .SRCINFO.
Identical: 12233 of those 12244 answers (99.91%); 11 differ.

Why files were refused (unsupported):
       24 top-level for loop: eval
       17 top-level if
       10 declare -gA
        8 top-level for loop
        4 top-level && or || or pipe
        3 command substitution $(echo | awk)
        2 top-level for loop: declare with a dynamic name
        2 command substitution $(printf)
        2 command substitution $(awk)
        1 command substitution $(sed)
        1 command substitution $(perl)
        1 command substitution $(date)

Files with status mismatch:
  imagemagick  (./minieval-diffs/imagemagick.diff)
  lib32-libxv  (./minieval-diffs/lib32-libxv.diff)
  libxrandr  (./minieval-diffs/libxrandr.diff)
  libxres  (./minieval-diffs/libxres.diff)
  libxt  (./minieval-diffs/libxt.diff)
  libxv  (./minieval-diffs/libxv.diff)
  papirus-icon-theme  (./minieval-diffs/papirus-icon-theme.diff)
  smali  (./minieval-diffs/smali.diff)
  xf86-video-fbdev  (./minieval-diffs/xf86-video-fbdev.diff)
  xorg-appres  (./minieval-diffs/xorg-appres.diff)
  xorg-xbacklight  (./minieval-diffs/xorg-xbacklight.diff)

Files with status parse-error:
  lib32-libffi  (monorepo//lib32-libffi/PKGBUILD:13:28: array element values must be words)
  lib32-libjpeg-turbo  (monorepo//lib32-libjpeg-turbo/PKGBUILD:20:9: array element values must be words)
  lib32-libsndfile  (monorepo//lib32-libsndfile/PKGBUILD:15:9: array element values must be words)
  lib32-ncurses  (monorepo//lib32-ncurses/PKGBUILD:16:9: array element values must be words)
  lib32-nettle  (monorepo//lib32-nettle/PKGBUILD:14:49: array element values must be words)
  lib32-pipewire  (monorepo//lib32-pipewire/PKGBUILD:77:25: [ must be followed by an expression)
  nginx-mainline  (monorepo//nginx-mainline/PKGBUILD:288:20: array element values must be words)
  nginx  (monorepo//nginx/PKGBUILD:273:20: array element values must be words)
  pipewire-media-session  (monorepo//pipewire-media-session/PKGBUILD:36:25: [ must be followed by an expression)
  pipewire  (monorepo//pipewire/PKGBUILD:111:25: [ must be followed by an expression)
  python-ansible-compat  (monorepo//python-ansible-compat/PKGBUILD:52:71: "[x]" must be followed by =)
  python-b2sdk  (monorepo//python-b2sdk/PKGBUILD:64:105: "[x]" must be followed by =)
  python-datasets  (monorepo//python-datasets/PKGBUILD:101:74: ternary operator missing ? before :)
  python-distributed  (monorepo//python-distributed/PKGBUILD:105:86: "[x]" must be followed by =)
  python-partd  (monorepo//python-partd/PKGBUILD:50:71: "[x]" must be followed by =)
  python-s3fs  (monorepo//python-s3fs/PKGBUILD:48:92: ternary operator missing : after ?)
  python-scikit-image  (monorepo//python-scikit-image/PKGBUILD:74:90: "[x]" must be followed by =)
  python-virtualenv  (monorepo//python-virtualenv/PKGBUILD:87:70: "[x]" must be followed by =)
```

When the output of the interpreter differs from the official file it produces a diff in the `minieval-diffs` directory.
