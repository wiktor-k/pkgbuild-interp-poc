# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Jan de Groot <jgc@archlinux.org>

pkgname=qqwing
pkgver=1.3.4
pkgrel=8
pkgdesc='Sudoku generating and solving software'
arch=(x86_64)
url='https://qqwing.com/'
license=(GPL-2.0-or-later)
depends=(
  glibc
  libgcc
  libstdc++
)
makedepends=(git)
source=("git+https://github.com/stephenostermiller/qqwing.git#tag=v$pkgver")
b2sums=(15ba43b7ea118d3bae484f956125abdc5ca2b2713545183986ccbe72e6bbe5b40fe86215376795fe2926e645d92cecea03611db431a04895ffff095a6c841a24)

prepare() {
  cd $pkgname
  cp build/configure.ac build/Makefile.am build/qqwing.pc.in src/cpp
  cp doc/qqwing.man src/cpp/qqwing.1

  cd src/cpp
  autoreconf -fiv
}

build() {
  cd $pkgname/src/cpp
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var
  sed -i -e 's/ -shared / -Wl,-O1,--as-needed\0/g' libtool
  make
}

package() {
  cd $pkgname/src/cpp
  make DESTDIR="$pkgdir" install
}
