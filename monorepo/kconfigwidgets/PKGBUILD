# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Andrea Scarpino <andrea@archlinux.org>

pkgname=kconfigwidgets
pkgver=6.31.0
pkgrel=1
pkgdesc='Widgets for KConfig'
arch=(x86_64)
url='https://develop.kde.org/products/frameworks/'
license=(LGPL-2.0-only LGPL-3.0-only)
depends=(libstdc++
         glibc
         kcodecs
         kcolorscheme
         kconfig
         kcoreaddons
         ki18n
         kwidgetsaddons
         qt6-base)
makedepends=(doxygen
             extra-cmake-modules
             qt6-tools)
groups=(kf6)
source=(https://download.kde.org/stable/frameworks/${pkgver%.*}/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('a9aa37565702ac1353c36cf16cc968c8aa336d33bcab9d0ff5419dd0b1dc92df'
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
