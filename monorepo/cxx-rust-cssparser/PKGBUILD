# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=cxx-rust-cssparser
pkgver=1.1.0
pkgrel=1
arch=(x86_64)
pkgdesc='Library for parsing CSS using the Rust cssparser crate'
url='https://kde.org/plasma-desktop/'
license=(LGPL-2.1-only
         LGPL-3.0-only)
depends=(glibc
         libgcc
         libstdc++)
makedepends=(corrosion
             extra-cmake-modules
             qt6-base)
source=(https://download.kde.org/stable/$pkgname/$pkgname-$pkgver.tar.xz)
sha256sums=('c88f883f67d132240f7ef5c8b8fe4f743a79c448b98ac88f2be3b68264840358')
options=(!lto)

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DCMAKE_INSTALL_PREFIX=/usr
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
