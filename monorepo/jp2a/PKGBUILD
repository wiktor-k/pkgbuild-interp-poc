# Maintainer: Morten Linderud <foxboron@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=jp2a
pkgver=1.3.3
pkgrel=1
pkgdesc='A small utility for converting JPG images to ASCII'
arch=('x86_64')
url='https://github.com/Talinx/jp2a'
license=('GPL-2.0-only')
depends=(
  'curl'
  'glibc'
  'libexif'
  'libjpeg-turbo'
  'libpng'
  'libwebp'
  'ncurses'
)
makedepends=(
  'doxygen'
  'git'
)
source=("git+${url}.git#tag=v${pkgver}")
b2sums=('6bb4b258f7595c5a9599f4651a74aa292a23a27228d7a60603a90d224d3418cfe69a16630ea62baf872064293f5d9e4967b605e399df9fc4df758a6b924d68c4')

prepare() {
  cd ${pkgname}
  autoreconf -vi
}

build() {
  cd ${pkgname}
  bashcompdir=/usr/share/bash-completion/completions ./configure --prefix=/usr
  make
}

check() {
  cd ${pkgname}
  make check
}

package() {
  cd ${pkgname}
  make DESTDIR="${pkgdir}" install
}
