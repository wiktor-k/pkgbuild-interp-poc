# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Alexander F Rødseth <xyproto@archlinux.org>
# Contributor: Dominik Picheta <morfeusz8@gmail.com>
# Contributor: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Jesus Alvarez <jeezusjr@gmail.com>

pkgname=nim
pkgver=2.2.12
pkgrel=1
pkgdesc='Imperative, multi-paradigm, compiled programming language'
url='https://nim-lang.org/'
arch=(x86_64)
license=(MIT)
depends=(
  bash
  gcc
  libgcc
  glibc
)
makedepends=(
  git
  nodejs
  help2man
)
provides=(nimble)
conflicts=(nimble)
replaces=(nimble)
options=(!emptydirs)
backup=(
  etc/nim/config.nims
  etc/nim/nim.cfg
  etc/nim/nimdoc.cfg
  etc/nim/nimdoc.tex.cfg
  etc/nim/rename.rules.cfg
)
source=(
  "$pkgname::git+https://github.com/nim-lang/Nim#tag=v$pkgver"
  build-docs.patch
  'nim-gh-pr-24405.patch::https://patch-diff.githubusercontent.com/raw/nim-lang/Nim/pull/24405.patch'
)
sha512sums=('45384d28615bb00fd7eb15ba07ebe61e5d611404d324a06c9dfb2f1ebf090fb634ba5237a339f69927bbee612278fb850c3d5952d2d50ffed00c300931f7dd24'
            '73e359917fa721d069bf4c3de63a217624aedc334529161f1976fd9d6b38b8d24c5bfca4e2f0ba34e3e6e0fc3394268c4d785e76ad7d24355bb8394db0f0fbf0'
            '82b63ffded3672ca42ba2185d3c2478989685974ee96d8a791b0f660cd273a79ec8916a2da9f29f1bd31a444377f31aa046f6b6e6974c4bea58e54b2befebf9a')
b2sums=('53c7b74909ef4791ad9b150279278b637fd20761a5f9845f0b7aeb72bc1525771cca1b96462217417c036996be2a70840b0b510731280858cc2466bdbb211800'
        'adfffe6b6df889e0cf3c1e37f43ba5c147ce50636150f74b2d73d8dd7df7f60b570c835af3d84d8b249852fa00eb180a4c54330d830947c4e829a2c87b22c4d0'
        '1e28435933d66178bf718851d6e93e714434a53ddab28175d20bd89f6327a8bcb451a4178318bc7b13ae30a4617d1a5137031ad7c082f382c868d6ce8a41918b')

prepare() {
  cd "$pkgname"

  # ensure build_all.sh generates docs
  patch -p1 -i "$srcdir/build-docs.patch"

  # migrate from pcre to pcre2
  # https://github.com/nim-lang/Nim/pull/24405
  git apply "$srcdir/nim-gh-pr-24405.patch"
}

build() {
  cd "$pkgname"

  # https://nim-lang.org/docs/intern.html#bootstrapping-the-compiler-reproducible-builds
  export SOURCE_DATE_EPOCH=$(git log -n 1 --format=%at)

  ./build_all.sh

  # generate man pages
  local h2m_args=(
    --section=1
    --no-info
    --version-string="$pkgver"
  )
  help2man --name='Nim Language Compiler' "${h2m_args[@]}" -o nim.1 ./bin/nim
  help2man --name='Nimsuggest' "${h2m_args[@]}" -o nimsuggest.1 ./bin/nimsuggest
  help2man --name='Nimgrep' "${h2m_args[@]}" -o nimgrep.1 ./bin/nimgrep
  help2man --name='Nimpretty' "${h2m_args[@]}" -o nimpretty.1 ./bin/nimpretty
  help2man --name='Nim Package Installer' "${h2m_args[@]}" -o nimble.1 ./bin/nimble
  help2man --name='Atlas' "${h2m_args[@]}" -o atlas.1 ./bin/atlas

  # generate install.sh
  ./koch distrohelper
}

package() {
  cd "$pkgname"

  DESTDIR="$pkgdir" ./install.sh /usr/bin

  # install atlas+nimpretty, as they are not in `install.sh`
  # https://gitlab.archlinux.org/archlinux/packaging/packages/nim/-/work_items/6
  install -vDm755 -t "$pkgdir/usr/bin" bin/{atlas,nimpretty}

  # documentation
  install -vd "$pkgdir/usr/share/doc/$pkgname"
  cp -vr doc/html "$pkgdir/usr/share/doc/$pkgname" 
  find "$pkgdir/usr/share/doc/$pkgname" -name '*.idx' -delete

  # tools
  for fn in nimble nimsuggest nimgrep nim-gdb; do cp ./bin/$fn $pkgdir/usr/bin/; done
  install -vDm644 -t "$pkgdir/usr/lib/nim/doc" doc/nimdoc.{css,cls}
  install -vDm644 -t "$pkgdir/usr/lib/nim/tools" tools/debug/nim-gdb.py
  install -vDm644 -t "$pkgdir/usr/lib/nim/tools/dochack" tools/dochack/{dochack.js,fuzzysearch.nim}

  # man pages
  install -vDm644 -t "$pkgdir/usr/share/man/man1" ./*.1

  # shell completions
  for comp in {tools,dist/nimble}/*.bash-completion; do
    install -Dm 644 "${comp}" "${pkgdir}/usr/share/bash-completion/completions/$(basename "${comp/.bash-completion}")"
  done
  for comp in {tools,dist/nimble}/*.zsh-completion; do
    install -Dm 644 "${comp}" "${pkgdir}/usr/share/zsh/site-functions/_$(basename "${comp/.zsh-completion}")"
  done

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" copying.txt
}

# vim: ts=2 sw=2 et:
