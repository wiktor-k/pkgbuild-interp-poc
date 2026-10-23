# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Andrea Scarpino <andrea@archlinux.org>

pkgname=plasma-nm
pkgver=6.8.1
_dirver=$(echo $pkgver | cut -d. -f1-3)
pkgrel=1
pkgdesc='Plasma applet written in QML for managing network connections'
arch=(x86_64)
url='https://kde.org/plasma-desktop/'
license=(GPL-2.0-or-later)
depends=(glib2
         glibc
         kcmutils
         kcolorscheme
         kcompletion
         kconfig
         kcoreaddons
         kdbusaddons
         kdeclarative
         ki18n
         kio
         kirigami
         kirigami-addons
         kjobwidgets
         knotifications
         kquickcharts
         kservice
         kwidgetsaddons
         kwindowsystem
         libnm
         libplasma
         libstdc++
         modemmanager-qt
         networkmanager-qt
         openssl
         plasma-workspace
         prison
         qcoro
         qt6-base
         qt6-declarative
         qtkeychain-qt6
         solid)
makedepends=(extra-cmake-modules
             openconnect
	     qt6-webengine)
optdepends=('openconnect: Cisco AnyConnect VPN plugin'
            'qrca: QR code scanner'
            'qt6-webengine: Cisco AnyConnect VPN plugin')
groups=(plasma)
source=(https://download.kde.org/stable/plasma/$_dirver/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('1cc6a5ffb35d8f11be2ff4e8f4c558de074100889fecb004c07eeb78ca169a4c'
            'SKIP')
validpgpkeys=('E0A3EB202F8E57528E13E72FD7574483BB57B18D'  # Jonathan Esk-Riddell <jr@jriddell.org>
              '0AAC775BB6437A8D9AF7A3ACFE0784117FBCE11D'  # Bhushan Shah <bshah@kde.org>
              'D07BD8662C56CB291B316EB2F5675605C74E02CF'  # David Edmundson <davidedmundson@kde.org>
              '1FA881591C26B276D7A5518EEAAF29B42A678C20') # Marco Martin <notmart@gmail.com>

build() {
  cmake -B build  -S $pkgname-$pkgver \
    -DBUILD_TESTING=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
