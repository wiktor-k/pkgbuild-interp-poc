# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: ObserverOfTime <chronobserver@disroot.org>

pkgname=tree-sitter-bash
pkgver=0.25.1
pkgrel=2
pkgdesc='Bash grammar for tree-sitter'
arch=(x86_64)
url=https://github.com/tree-sitter/tree-sitter-bash
license=(MIT)
groups=(tree-sitter-grammars)
depends=(
  glibc
)
makedepends=(
  git
  tree-sitter-cli
)
optdepends=('tree-sitter: core library')
provides=("lib$pkgname.so")
source=("git+$url.git#commit=v$pkgver")
b2sums=('a7fc497386f2f72d1e04c60bbf825e443e816ed1be1bb8591a8ea233b5e0aebe059713035ebfd21f23f5b19c9d4f4d1b6e2e6564f5f5a0958470d46cc834bc63')
#validpgpkeys=(FCC13F47A6900D64239FF13BE67890ADC4227273) # Amaan Qureshi <amaanq12@gmail.com>

prepare() {
  cd $pkgname
  tree-sitter generate src/grammar.json
}

build() {
  cd $pkgname
  make PREFIX=/usr PARSER_URL=$url
}

package() {
  install -d "$pkgdir"/usr/lib/tree_sitter
  ln -s /usr/lib/lib$pkgname.so \
    "$pkgdir"/usr/lib/tree_sitter/${pkgname#tree-sitter-}.so

  cd $pkgname
  make DESTDIR="$pkgdir" PREFIX=/usr install
  install -Dm644 -t "$pkgdir"/usr/share/doc/$pkgname README.md
  install -Dm644 -t "$pkgdir"/usr/share/licenses/$pkgname LICENSE
}
