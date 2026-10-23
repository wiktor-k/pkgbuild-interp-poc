# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Maintainer: Bruno Pagani <archange@archlinux.org>

pkgbase=kvantum
pkgname=(kvantum
         kvantum-qt5)
pkgver=1.1.8
pkgrel=1
pkgdesc='SVG-based theme engine for Qt6 (including config tool and extra themes)'
arch=(x86_64)
url='https://github.com/tsujan/Kvantum'
license=(GPL-3.0-or-later)
depends=(glibc
         libgcc
         libstdc++
         libx11)
makedepends=(cmake
             git
             qt5-svg
             qt5-tools
             qt5-x11extras
             kwindowsystem
             qt6-svg
             qt6-tools)
source=(git+https://github.com/tsujan/Kvantum#tag=V$pkgver?signed)
sha256sums=('c407e9fb1e9e985b3af78d7535445da31f03473387b1c45b4c810883b8c27ec6')
validpgpkeys=(19DFDF3A579BD509DBB572D8BE793007AD22DF7E) # Pedram Pourang <tsujan2000@gmail.com>

build() {
  cmake -B build5 -S Kvantum/Kvantum \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DENABLE_QT5=ON
  make -C build5

  cmake -B build6 -S Kvantum/Kvantum \
    -DCMAKE_INSTALL_PREFIX=/usr
  make -C build6
}

package_kvantum() {
  depends+=(kwindowsystem
            qt6-base
            qt6-svg)
  optdepends=('kvantum-qt5: Qt5 style')

  DESTDIR="$pkgdir" cmake --install build6
}

package_kvantum-qt5() {
  pkgdesc='SVG-based theme engine for Qt5'
  depends+=(kvantum
            qt5-base
            qt5-svg
            qt5-x11extras)

  DESTDIR="$pkgdir" cmake --install build5/style
}
