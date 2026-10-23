# Maintainer: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Angel Velasquez <angvp@archlinux.org>
# Contributor: Patrick Melo <patrick@patrickmelo.eti.br>

pkgname=ctpl
pkgver=0.3.5
pkgrel=3
pkgdesc='Template engine library written in C'
arch=(x86_64)
url='https://ctpl.tuxfamily.org/'
license=(GPL-2.0-or-later)
makedepends=(gtk-doc)
source=("https://git.tuxfamily.org/ctpl/ctpl.git/snapshot/ctpl-$pkgver.tar.gz")
b2sums=('101f83a4154104d23bbb0e8540923dab506df653f9c8e9eafb2035e0e4bbc094b66a5a5d4b67bd0572f4f60a20593aedc145df82ef55227fb6bbb5c8f451babf')

build() {
  cd $pkgname-$pkgver
  ./autogen.sh
  ./configure --prefix=/usr
  make
}

package() {
  make -C $pkgname-$pkgver DESTDIR="$pkgdir" install
}
