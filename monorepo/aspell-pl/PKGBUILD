# Maintainer: Johannes Löthberg <johannes@kyriasis.com>
# Contributor: Bartłomiej Piotrowski <barthalion@gmail.com>
# Contributor: Jaroslaw Swierczynski <swiergot@aur.archlinux.org>
# Contributor: Arkadiusz Laczynski <alaczynski@gmail.com>
# Contributor: Jaroslav Lichtblau <svetlemodry@archlinux.org>

pkgname=aspell-pl
pkgver=20251101
pkgrel=1
pkgdesc='Polish dictionary for aspell'
arch=('any')
url="https://www.sjp.pl/slownik/en/"
license=('GPL' 'LGPL' 'MPL' 'CC SA' 'Apache')
depends=('aspell')
source=(https://sjp.pl/sl/ort/sjp-aspell6-pl-6.0_$pkgver-0.tar.bz2)
sha256sums=('a4dd526bc5004cb73c204f0734b1c8107578ed7d7771a42d00844d995f2964a5')

build() {
  cd aspell6-pl-6.0_$pkgver-0

  ./configure
  make
}

package() {
  cd aspell6-pl-6.0_$pkgver-0

  make DESTDIR="$pkgdir" install
}
