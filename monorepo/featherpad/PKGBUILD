# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Contributor: Peter Mattern <pmattern at arcor dot de>

pkgname=featherpad
pkgver=1.6.4
pkgrel=1
pkgdesc='Lightweight Qt plain text editor'
arch=('x86_64')
url='https://github.com/tsujan/FeatherPad'
license=('GPL3')
depends=('hicolor-icon-theme' 'hunspell' 'qt6-svg')
makedepends=('cmake' 'git' 'ninja' 'qt6-tools')
source=("git+https://github.com/tsujan/FeatherPad.git#tag=V${pkgver}?signed")
sha256sums=('0ad200b8d0e9b1298a8c39b8415f3467535ffa28b2319e91311bca8fc8526ca3')
validpgpkeys=('19DFDF3A579BD509DBB572D8BE793007AD22DF7E')

build() {
  cmake -B build -S FeatherPad \
        -G Ninja \
        -DCMAKE_INSTALL_PREFIX='/usr' \
        -DCMAKE_BUILD_TYPE=None \
        -DENABLE_QT5=OFF \
        -Wno-dev
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
