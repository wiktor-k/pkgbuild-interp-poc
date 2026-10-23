# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Donald Carr <sirspudd@gmail.com>
# Contributor: Jake Petroules <jake.petroules@qt.io>

pkgname=qbs
pkgver=3.3.1
pkgrel=1
pkgdesc='Cross platform build tool'
arch=(x86_64)
url='https://wiki.qt.io/Qbs'
license=(LGPL-2.1-only)
depends=(glibc
         libgcc
         libstdc++
         qt6-base
         qt6-5compat)
makedepends=(cmake
             git)
source=(git+https://code.qt.io/qbs/qbs#tag=v$pkgver
        qt-6.10.patch)
sha256sums=('a8da288699897592fd376b06fda8d2802d6dae2325b33cfe89603221b7ae276a'
            '463fc358c4693048b7db61ad5f6ba2801c54e1ead5daa4f8fc4ca8d70558088b')

prepare() {
  patch -d $pkgname -p1 < qt-6.10.patch
}

build() {
  cmake -B build -S $pkgname \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DWITH_TESTS=OFF \
    -DQBS_LIBEXEC_INSTALL_DIR=lib/qbs
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
