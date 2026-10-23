# Maintainer: Alexander F Rødseth <xyproto@archlinux.org>
# Contributor: Jared Casper <jaredcasper@gmail.com>
# Contributor: Paulo Matias <matias archlinux-br org>

pkgname=iverilog
pkgver=13.0
pkgrel=1
pkgdesc='Icarus Verilog compiler and simulation tool'
arch=('x86_64')
url='https://github.com/steveicarus/iverilog'
license=('GPL')
depends=(
  bzip2
  zlib
)
makedepends=(
  git
  gperf
)
options=(
  staticlibs
)
source=(
  "git+https://github.com/steveicarus/iverilog#tag=v${pkgver/./_}"
)
b2sums=('e1a6dd5c99aa5094042a0e23f68039b9882fc376a85e754e8ef6ec1a21166681df04c28ef4005ccc8402db04449dcdedb9a4e74658525d694bd2b306ebfdf4a3')

prepare() {
  cd "$pkgname"

  aclocal
  autoconf
}

build() {
  cd "$pkgname"

  CFLAGS+=' -ffat-lto-objects -fcommon' # https://wiki.gentoo.org/wiki/Gcc_10_porting_notes/fno_common

  CXXCPP=/usr/bin/cpp ./configure --prefix=/usr
  make
}

package() {
  make -C "$pkgname" -j1 prefix="$pkgdir/usr" install
}

# vim: ts=2 sw=2 et:
