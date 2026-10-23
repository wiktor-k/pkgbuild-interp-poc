# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Antonio Rojas <arojas@archlinux.org>
# Contributor: Devin Lin <devin@kde.org>

pkgname=plasma-nano
pkgver=6.8.1
_dirver=$(echo $pkgver | cut -d. -f1-3)
pkgrel=1
pkgdesc='A minimal Plasma shell package intended for embedded devices'
arch=(x86_64)
url='https://invent.kde.org/plasma/plasma-nano'
license=(GPL-2.0-or-later)
depends=(
  glibc
  kdeclarative
  kirigami
  kitemmodels
  knewstuff
  kservice
  ksvg
  kwayland
  kwindowsystem
  libplasma
  libstdc++
  plasma-workspace
  qt6-5compat
  qt6-base
  qt6-declarative
)
makedepends=(
  cmake
  extra-cmake-modules
)
source=("https://download.kde.org/stable/plasma/$_dirver/$pkgname-$pkgver.tar.xz"{,.sig})
sha256sums=('dc860d7e07040501eac7e19365be16cec9f056308b1440405213bbc14c32bae2'
            'SKIP')
validpgpkeys=(
  E0A3EB202F8E57528E13E72FD7574483BB57B18D # Jonathan Esk-Riddell <jr@jriddell.org>
  0AAC775BB6437A8D9AF7A3ACFE0784117FBCE11D # Bhushan Shah <bshah@kde.org>
  D07BD8662C56CB291B316EB2F5675605C74E02CF # David Edmundson <davidedmundson@kde.org>
  1FA881591C26B276D7A5518EEAAF29B42A678C20 # Marco Martin <notmart@gmail.com>
)

build() {
  cmake -B build -S $pkgname-$pkgver
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
