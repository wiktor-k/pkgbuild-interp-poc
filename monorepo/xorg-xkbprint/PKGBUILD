# Maintainer: AndyRTR <andyrtr@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Brian Bidulock <bidulock@openss7.org>

pkgname=xorg-xkbprint
pkgver=1.0.8
pkgrel=1
pkgdesc="Generates a PostScript description of an XKB keyboard description"
arch=(x86_64)
url="https://gitlab.freedesktop.org/xorg/app/xkbprint"
license=('LicenseRef-xkbprint')
depends=('libxkbfile' 'libx11' 'glibc')
makedepends=('xorg-util-macros')
groups=('xorg-apps')
source=(https://xorg.freedesktop.org/archive/individual/app/xkbprint-${pkgver}.tar.xz{,.sig})
sha512sums=('b2c1716f0dc30e528886190cfccfb72fb17f337bb321f988058759a1cbbf93f2effc68c1b2b9a0c808cc2831b62a8e894ca341ba0a35c22eaff4cc195a20d8da'
            'SKIP')
validpgpkeys=('3AB285232C46AE43D8E192F4DAB0F78EA6E7E2D2') # Alan Coopersmith <alan.coopersmith@oracle.com>

build() {
  cd xkbprint-$pkgver
  ./configure --prefix=/usr
  make
}

package() {
  cd xkbprint-$pkgver
  make DESTDIR="$pkgdir" install
  install -Dm644 COPYING -t "$pkgdir"/usr/share/licenses/$pkgname/
}
