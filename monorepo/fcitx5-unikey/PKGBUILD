# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: csslayer <wengxt AT gmail com>

pkgname=fcitx5-unikey
pkgver=5.1.12
pkgrel=1
pkgdesc="Unikey engine support for Fcitx5"
arch=('x86_64')
url="https://github.com/fcitx/fcitx5-unikey"
license=('GPL-2.0-or-later')
depends=('fcitx5-qt')
makedepends=('git' 'extra-cmake-modules' 'ninja')
source=("git+https://github.com/fcitx/fcitx5-unikey.git#tag=$pkgver?signed")
sha512sums=('5b4e08e29120a60359fee25e1a1a75b91da5fe0dc0a6f033ecc861df417ab9405edc95d28c7922b39b98188329f79a0d65e0d712f15bc56bdc6ecdea2f805f68')
validpgpkeys=('2CC8A0609AD2A479C65B6D5C8E8B898CBF2412F9') # Weng Xuetian <wengxt@gmail.com>

build() {
  cd $pkgname
  cmake -GNinja -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_INSTALL_LIBDIR=/usr/lib .
  ninja
}

package() {
  cd $pkgname
  DESTDIR="$pkgdir" ninja install
}
