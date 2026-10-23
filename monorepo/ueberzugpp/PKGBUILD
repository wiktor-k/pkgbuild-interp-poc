# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Contributor: JustKidding <jk@vin.ovh>

pkgname=ueberzugpp
pkgver=2.9.10
pkgrel=1
pkgdesc="Command line utility which allows to display images in the terminal, written in C++"
arch=('x86_64')
url="https://github.com/jstkdng/ueberzugpp"
license=("GPL-3.0-or-later")
provides=('ueberzug')
conflicts=("ueberzug")
makedepends=("cmake" "cli11" "nlohmann-json" "wayland-protocols" "extra-cmake-modules" "range-v3")
depends=(
  "opencv"
  "libvips"
  "glib2"
  "libxcb"
  "xcb-util-image"
  "libsixel"
  "openssl"
  "spdlog"
  "libglvnd"
  "fmt"
  "chafa"
  "wayland"
  "onetbb"
  "glibc"
  "libgcc"
  "libstdc++"
  "xcb-util-errors"
)
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/jstkdng/${pkgname}/archive/v${pkgver}.tar.gz")
sha256sums=('225b74ae3c71fd7d6dbd8b3a1740dacda95c40146a4fd00c792e4f18f200854d')

build() {
  # -DENABLE_TURBOBASE64=ON \
  cmake -B build -S "$pkgname-$pkgver" \
    -DCMAKE_BUILD_TYPE='None' \
    -DCMAKE_INSTALL_PREFIX='/usr' \
    -DENABLE_WAYLAND=ON \
    -DENABLE_XCB_ERRORS=ON \
    -DENABLE_OPENGL=ON \
    -Wno-dev
  cmake --build build -j "$(nproc)"
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}

# vim:set ts=2 sw=2 et:
