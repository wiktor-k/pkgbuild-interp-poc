# Maintainer: Jan de Groot <jgc@archlinux.org>
# Contributor: Eduardo Romero <eduardo@archlinux.org>
# Contributor: Jason Chu <jason@archlinux.org>

pkgname=gphoto2
pkgver=2.5.32
pkgrel=1
pkgdesc="A digital camera download and access program."
arch=(x86_64)
url="http://www.gphoto.org"
license=('GPL' 'LGPL')
depends=('libgphoto2' 'popt')
source=(https://downloads.sourceforge.net/gphoto/${pkgname}-${pkgver}.tar.bz2{,.asc})
sha256sums=('4e379a0f12f72b49ee5ee2283ffd806b5d12d099939d75197a3f4bbc7f27a1a1'
            'SKIP')
validpgpkeys=('7C4AFD61D8AAE7570796A5172209D6902F969C95')

build() {
  cd $pkgname-$pkgver
  ./configure --prefix=/usr
  make
}

package() {
  cd $pkgname-$pkgver
  make DESTDIR="${pkgdir}" install
}
