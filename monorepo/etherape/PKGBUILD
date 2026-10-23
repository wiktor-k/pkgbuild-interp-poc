# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Thorsten Töpper <atsutane-tu@freethoughts.de>
# Contributor: kevin <kevin@archlinux.org>
# Contributor: Damir Perisa <damir.perisa@bluewin.ch>

pkgname=etherape
pkgver=0.9.22
pkgrel=1
pkgdesc='Graphical network monitor for various OSI layers and protocols'
url='https://etherape.sourceforge.io/'
arch=('x86_64')
license=('GPL-2.0-or-later')
depends=('glibc' 'glib2' 'libpcap' 'popt' 'cairo' 'gdk-pixbuf2' 'gtk3' libpcap.so)
makedepends=('itstool' 'yelp-tools')
source=(https://downloads.sourceforge.net/${pkgname}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('3a991d52812db6d0a027f03a4cef8b2e9fef8bb83ae0e63ae1fd106a5710ce74bbf918b027179325ac45bcc07421f3ba8ec7342a4d34393684b599f7a7dd71bc')
b2sums=('0e4427cf95e0d619c60a5541d1c79dbed5cf36a4402c9c3939057824b8615a0a1b8c0588e9ec538e837b99213c896f7ed09c793a1f41cab715242a630d95efa5')

build() {
  cd ${pkgname}-${pkgver}
  ./configure --prefix=/usr --sysconfdir=/etc
  make
}

package() {
  cd ${pkgname}-${pkgver}
  make DESTDIR="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
