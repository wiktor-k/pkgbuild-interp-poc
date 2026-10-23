# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=plasma-firewall
pkgver=6.8.1
_dirver=$(echo $pkgver | cut -d. -f1-3)
pkgrel=1
pkgdesc='Control Panel for your system firewall'
arch=(x86_64)
url='https://kde.org/plasma-desktop/'
license=(GPL-2.0-or-later LGPL-2.0-or-later)
depends=(glibc
         kauth
         kcmutils
         kconfig
         kcoreaddons
         ki18n
         kirigami
         libstdc++
         qt6-base
         qt6-declarative)
makedepends=(extra-cmake-modules
             python)
optdepends=('firewalld: firewalld backend'
            'iproute2: netstat backend'
            'python: ufw backend'
            'ufw: ufw backend')
groups=(plasma)
source=(https://download.kde.org/stable/plasma/$_dirver/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('ebde7709a5df6cd1db20c57a503203b70d39dcb0bdbf5e2a9fb46df3d4ae3c5f'
            'SKIP')
validpgpkeys=('E0A3EB202F8E57528E13E72FD7574483BB57B18D'  # Jonathan Esk-Riddell <jr@jriddell.org>
              '0AAC775BB6437A8D9AF7A3ACFE0784117FBCE11D'  # Bhushan Shah <bshah@kde.org>
              'D07BD8662C56CB291B316EB2F5675605C74E02CF'  # David Edmundson <davidedmundson@kde.org>
              '1FA881591C26B276D7A5518EEAAF29B42A678C20') # Marco Martin <notmart@gmail.com>

build() {
  cmake -B build  -S $pkgname-$pkgver \
    -DKDE_INSTALL_LIBEXECDIR=lib \
    -DBUILD_TESTING=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
