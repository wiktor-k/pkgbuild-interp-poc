# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=libmodbus
pkgver=3.2.0
pkgrel=1
pkgdesc='A Modbus library for Linux, Mac OS, FreeBSD and Windows'
arch=(x86_64)
url=https://github.com/stephane/libmodbus
license=(LGPL-2.1-or-later)
depends=(glibc)
makedepends=(git)
source=("git+$url#tag=v$pkgver")
b2sums=('63073bfc05498b40106bc3a2eab7b10c9355943eb3296048fc0bf4585270e95fd72edf1323c78272068794ea15128d7dd2ec93bfb10dba066e84db7352f1a0c3')

prepare() {
  cd $pkgname
  ./autogen.sh
}

build() {
  cd $pkgname
  ./configure --prefix=/usr
  make
}

check() {
  cd $pkgname
  make check
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
