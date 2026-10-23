# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Contributor: Laurent Carlier <lordheavym@gmail.com>
# Contributor: gumper <gumper1034@gmail.com>

pkgname=libalkimia
pkgver=8.2.1
pkgrel=1
pkgdesc="A library with common classes and functionality used by finance applications for the KDE SC"
arch=('x86_64')
url="https://community.kde.org/Alkimia/libalkimia"
license=('LGPL')
depends=('qt6-webengine' 'knewstuff' 'ktextwidgets')
makedepends=('cmake' 'doxygen' 'extra-cmake-modules' 'kdoctools')
source=(https://download.kde.org/stable/alkimia/$pkgver/alkimia-$pkgver.tar.xz{,.sig})
validpgpkeys=('91F17B048CF4668CD27FE9FBD3694D8B346FCA80'  # Ralf Habacker <ralf.habacker@freenet.de>
              'D69A745A55331F44F404D8258D4DE062AA2EB01C') # Thomas Baumgart <thb@net-bembel.de>
sha256sums=('f921410e180e0a5811e1ee2926954920c6576a72b3b65f53791faa6c85fcb689'
            'SKIP')

build() {
  cmake -B build -S alkimia-$pkgver \
    -DQT_MAJOR_VERSION=6 \
    -DBUILD_APPLETS=OFF
  cmake --build build
}

package() {
  DESTDIR="${pkgdir}" cmake --install build
}
