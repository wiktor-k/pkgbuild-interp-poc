# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=kommit
pkgver=1.8.1
pkgrel=2
pkgdesc='Git gui client'
arch=(x86_64)
url='https://apps.kde.org/kommit'
license=(GPL-3.0-only)
depends=(glibc
         kconfig
         kconfigwidgets
         kcoreaddons
         kcrash
         kdbusaddons
         ki18n
         kiconthemes
         kio
         kparts
         kwidgetsaddons
         kxmlgui
         libgit2
         libstdc++
         openssl
         qt6-base
         qt6-charts
         sh
         syntax-highlighting)
makedepends=(extra-cmake-modules
             dolphin
             git
             kdoctools
             ktexteditor)
source=(git+https://invent.kde.org/sdk/kommit#tag=v$pkgver)
sha256sums=('3280d24a3d2519f914cdfdaf2b61ff25b735fd2805254b15be42822547ff58cd')
validpgpkeys=('E0A3EB202F8E57528E13E72FD7574483BB57B18D') # Jonathan Esk-Riddell <jr@jriddell.org>

build() {
  cmake -B build -S $pkgname \
    -DCMAKE_BUILD_TYPE=None \
    -DQT_MAJOR_VERSION=6
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
