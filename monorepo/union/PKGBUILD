# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=union
pkgver=6.8.1
_dirver=$(echo $pkgver | cut -d. -f1-3)
pkgrel=1
arch=(x86_64)
pkgdesc='A Qt style supporting both QtQuick and QtWidgets'
url='https://kde.org/plasma-desktop/'
license=(LGPL-2.1-only OR LGPL-3.0-only OR LicenseRef-KDE-Accepted-LGPL)
depends=(cxx-rust-cssparser
         glibc
         kcolorscheme
         kconfig
         kcoreaddons
         kguiaddons
         kiconthemes
         kirigami
         libgcc
         libstdc++
         qt6-base
         qt6-declarative
         zlib)
makedepends=(extra-cmake-modules
             git
             qt6-shadertools)
groups=(plasma)
source=(https://download.kde.org/stable/plasma/$_dirver/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('abea413c79fe9a52337a9b8a50ee8fb29be7ed38bd764eb7555d724e99b39f00'
            'SKIP')
validpgpkeys=('E0A3EB202F8E57528E13E72FD7574483BB57B18D'  # Jonathan Esk-Riddell <jr@jriddell.org>
              '0AAC775BB6437A8D9AF7A3ACFE0784117FBCE11D'  # Bhushan Shah <bshah@kde.org>
              'D07BD8662C56CB291B316EB2F5675605C74E02CF'  # David Edmundson <davidedmundson@kde.org>
              '1FA881591C26B276D7A5518EEAAF29B42A678C20') # Marco Martin <notmart@gmail.com>

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DBUILD_OUTPUT_QTWIDGETS=ON \
    -DBUILD_TESTING=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
