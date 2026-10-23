# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=kirigami-app-components
pkgver=1.1.1
pkgrel=1
pkgdesc='Kirigami addons and modules necessary to do a full featured KDE application'
url='https://invent.kde.org/libraries/kirigami-app-components'
arch=(x86_64)
license=(GPL-2.0-or-later
         LGPL-2.1-or-later)
depends=(glibc
         kconfig
         kcoreaddons
         kguiaddons
         ki18n
         kirigami
         kitemmodels
         libstdc++
         qt6-base
         qt6-declarative)
makedepends=(extra-cmake-modules)
source=(https://download.kde.org/stable/$pkgname/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('f68f4d082933213d82b805217f2fa0690f51df7d00214ddb873812a60a6b830b'
            'SKIP')
validpgpkeys=(1FA881591C26B276D7A5518EEAAF29B42A678C20) # Marco Martin <notmart@gmail.com>

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DBUILD_TESTING=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
