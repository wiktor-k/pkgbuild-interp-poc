# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=ktexttemplate
pkgver=6.31.0
pkgrel=1
pkgdesc='Library to allow application developers to separate the structure of documents from the data they contain'
arch=(x86_64)
url='https://develop.kde.org/products/frameworks/'
license=(LGPL-2.1-or-later)
depends=(glibc
         libgcc
         libstdc++
         qt6-base
         qt6-declarative)
makedepends=(doxygen
             extra-cmake-modules
             qt6-tools)
groups=(kf6)
source=(https://download.kde.org/stable/frameworks/${pkgver%.*}/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('461c0d1672430f646d98d89ec226912f7ace031fda020a209b9ce270fd146c0c'
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
