# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=fcitx5-bamboo
pkgver=1.0.11
pkgrel=1
pkgdesc="Bamboo (Vietnamese Input Method) engine support for Fcitx"
arch=('x86_64')
url="https://github.com/fcitx/fcitx5-bamboo"
license=('LGPL-2.1-or-later')
depends=('fcitx5')
makedepends=('git' 'extra-cmake-modules' 'go')
source=("git+https://github.com/fcitx/fcitx5-bamboo.git#tag=$pkgver?signed"
        "git+https://github.com/BambooEngine/bamboo-core.git")
sha512sums=('45a0d003c3099c7d0dc5f92f8dee129c62f8baeddb7de12d9c9079bb29a539a025749db8db942e151eb6f7c5cfbde3ec18f67881c81e614a6123ec4d73c64fea'
            'SKIP')
validpgpkeys=('2CC8A0609AD2A479C65B6D5C8E8B898CBF2412F9') # Weng Xuetian <wengxt@gmail.com>

prepare() {
  cd fcitx5-bamboo
  git submodule init
  git config submodule.bamboo/bamboo-core.url "$srcdir"/bamboo-core
  git -c protocol.file.allow=always submodule update
}

build(){
  cd $pkgname
  cmake -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_INSTALL_LIBDIR=/usr/lib .
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
