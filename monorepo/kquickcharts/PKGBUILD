# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=kquickcharts
pkgver=6.31.0
pkgrel=1
pkgdesc='A QtQuick plugin providing high-performance charts'
arch=(x86_64)
url='https://develop.kde.org/products/frameworks/'
license=(LGPL-2.0-only LGPL-3.0-only)
depends=(glibc
         kirigami
         libstdc++
         qt6-base
         qt6-declarative)
makedepends=(extra-cmake-modules
             qt6-shadertools
             qt6-tools)
groups=(kf6)
source=(https://download.kde.org/stable/frameworks/${pkgver%.*}/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('2c2c18c1eb1ea1c8d9df0dc084bd2770b0099dc10cc0cb6827c7dc6bd5d3c21d'
            'SKIP')
validpgpkeys=(53E6B47B45CEA3E0D5B7457758D0EE648A48B3BB # David Faure <faure@kde.org>
              90A968ACA84537CC27B99EAF2C8DF587A6D4AAC1 # Nicolas Fella <nicolas.fella@kde.org>
              )

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DBUILD_TESTING=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
