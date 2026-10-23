# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Evangelos Foutras <foutrelis@archlinux.org>
# Contributor: Aaron Griffin <aaron@archlinux.org>
# Contributor: Tom Newsom <Jeepster@gmx.co.uk>

pkgname=ltrace
pkgver=0.8.1
pkgrel=1
pkgdesc="Tracks runtime library calls in dynamically linked programs"
arch=('x86_64')
url="https://www.ltrace.org/"
license=('GPL-2.0-or-later')
depends=(
  'glibc'
  'libelf'
  'libgcc'
)
makedepends=(
  'dejagnu'
  'git'
)
source=("git+https://gitlab.com/cespedes/ltrace.git#tag=$pkgver")
b2sums=('ca0f1b7ff9c7b0aa19f0f4156f17f9dee6a67c1dc9007e91cae5fcaf2b3797db36cd803d9e42a24161fc6f112dd322ca6265de3675a5a21080e337971d83b0ae')

prepare() {
  cd $pkgname
  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --disable-werror \
    --without-libunwind
  make
}

check() {
  cd $pkgname
  # Many test expectations use the old output format from before upstream
  # commit 6645e9a which changed whitespace/delimiter formatting in struct
  # and array output (e.g. "{ 1, 2 }" -> "{1, 2}", "[ x ]" -> "[x]").
  # The test suite was not updated to match.
  make -k check || true
}

package(){
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
