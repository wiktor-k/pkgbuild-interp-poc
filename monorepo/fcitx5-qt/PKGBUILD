# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: csslayer <wengxt AT gmail com>

pkgname=fcitx5-qt
pkgver=5.1.16
pkgrel=2
pkgdesc="Fcitx5 Qt Library (Qt5 & Qt6 integrations)"
arch=('x86_64')
url="https://github.com/fcitx/fcitx5-qt"
license=('LGPL-2.1-or-later' 'BSD-3-Clause' 'GPL-2.0-or-later')
groups=('fcitx5-im')
depends=('glibc' 'libgcc' 'libstdc++' 'libxcb' 'libxkbcommon' 'fcitx5' 'qt6-base' 'wayland')
makedepends=('git' 'extra-cmake-modules' 'qt5-base' 'ninja')
source=("git+https://github.com/fcitx/fcitx5-qt.git#tag=$pkgver?signed")
sha512sums=('7f50a7530a666182342d0dd59d6a614de29576a9718e3048922fe759e3f0fa1f65ac40388360fee8fd67ec3c402472518739d02f1367692bd6e7d9c5f346a258')
validpgpkeys=('2CC8A0609AD2A479C65B6D5C8E8B898CBF2412F9') # Weng Xuetian <wengxt@gmail.com>

build() {
  cd $pkgname
  cmake . -GNinja \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DCMAKE_INSTALL_LIBEXECDIR=/usr/lib
  ninja
}

check() {
  cd $pkgname
  ninja test
}

package() {
  cd $pkgname
  DESTDIR="$pkgdir" ninja install
  install -Dm644 LICENSES/BSD-3-Clause.txt -t "$pkgdir/usr/share/licenses/$pkgname/"
}
