# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Leonidas Spyropoulos <artafinde at gmail com>
# Contributor: Miguel Revilla <yo@miguelrevilla.com>
# Contributor: Joshua Stiefer <facedelajunk@gmail.com>
# Contributor: Alexander Fehr <pizzapunk gmail com>
# Contributor: Daniel J Griffiths <ghost1227@archlinux.us>

pkgname=wxsvg
pkgver=1.5.25
pkgrel=7
pkgdesc='C++ library to create, manipulate and render SVG files'
arch=(x86_64)
url='http://wxsvg.sourceforge.net/'
license=('LGPL-2.0-or-later WITH WxWindows-exception-3.1')
depends=(
  cairo
  expat
  ffmpeg
  glib2
  glibc
  libexif
  libgcc
  libstdc++
  pango
  wxwidgets-common
  wxwidgets-gtk3
)
source=("https://downloads.sourceforge.net/$pkgname/$pkgname-$pkgver.tar.bz2")
b2sums=(5efdda5e5e546f0d191f615b1699b88ec3dc9dfd5c4d3a618fa7bbc70416ddb8e9df870a96814ed788662bc2e7d9ba2e00744b53880df14d8c15baf46af0377c)

prepare() {
  cd $pkgname-$pkgver
  autoreconf -fiv
}

build() {
  cd $pkgname-$pkgver
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var
  make
}

package() {
  cd $pkgname-$pkgver
  make DESTDIR="$pkgdir" install
  install -Dm644 COPYING "$pkgdir/usr/share/licenses/$pkgname/COPYING"
}
