# Maintainer: Anatol Pomozov
# Contributor: Devin J. Pohly <djpohly+arch@gmail.com>
# Contributor: Andrej Gelenberg <andrej.gelenberg@udo.edu>

pkgname=evtest
pkgver=1.36
pkgrel=1
pkgdesc='Input device event monitor and query tool'
arch=(x86_64)
url='https://cgit.freedesktop.org/evtest/'
license=(GPL)
depends=(glibc)
makedepends=(asciidoc xmlto)
source=(https://gitlab.freedesktop.org/libevdev/evtest/-/archive/evtest-${pkgver}/evtest-evtest-${pkgver}.tar.gz)
sha256sums=('3b9a66c92e48b0cd13b689530b5729c031bc1bcbfe9d19c258f9245e2f8d2a0f')

build() {
  cd $pkgname-$pkgname-$pkgver
  ./autogen.sh
  ./configure --prefix=/usr 
  make
}

package() {
  cd $pkgname-$pkgname-$pkgver
  make DESTDIR="$pkgdir" install
}
