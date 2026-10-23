# Maintainer: Konstantin Gizdov <arch at kge dot pw>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: Brian Nguyen <bhnguyen@gmail.com>
# Original Packager: Tom K <tomk@runbox.com>

pkgname=hevea
pkgver=2.38
pkgrel=1
pkgdesc="A quite complete and fast LATEX to HTML translator"
arch=(x86_64)
depends=('texlive-bin')
makedepends=('ocaml' 'ocamlbuild')
license=('GPL')
url="http://pauillac.inria.fr/~maranget/hevea/"
source=(http://pauillac.inria.fr/~maranget/$pkgname/distri/$pkgname-$pkgver.tar.gz)
sha256sums=('722038065007226f0fa3de4629127294d2e29bfbbc41042c83a570fa0c455a47')

build() {
  cd $pkgname-$pkgver
  make PREFIX=/usr
}

package() {
  cd $pkgname-$pkgver
  DESTDIR="$pkgdir" make install
  mkdir -p "$pkgdir"/usr/share/texmf/tex/latex
  ln -s /usr/lib/hevea/hevea.sty "$pkgdir"/usr/share/texmf/tex/latex/hevea.sty
}
