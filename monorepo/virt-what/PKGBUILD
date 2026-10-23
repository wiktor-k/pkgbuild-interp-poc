# Maintainer: Tim Meusel <tim@bastelfreak.de>
# Contributor: Christian Hesse <mail@eworm.de>
# Contributor: Thomas S Hatch <thatch45 (at) Gmail (dot) com>

pkgname=virt-what
pkgver=1.27
pkgrel=1
license=('GPL-2.0-only')
arch=('x86_64')
depends=('dmidecode' 'bash')
pkgdesc='A small program that prints out a list of facts about a virtual machine'
url='http://people.redhat.com/~rjones/virt-what/'
validpgpkeys=('F7774FB1AD074A7E8C8767EA91738F73E1B768A0')
source=("https://people.redhat.com/~rjones/virt-what/files/${pkgname}-${pkgver}.tar.gz"{,.sig})
sha512sums=('8d10111e38a87504daccf36e9386863aa92308c96be65164412d92c56c12d42f38e6e251fabe7945ce74e1419fc82c0ac3cdbd0f25b19cd72feda2e199307d6f'
            'SKIP')

build() {
  cd "${srcdir}/${pkgname}-${pkgver}"

  ./configure
  sed -i '/^sbindir/c sbindir = /usr/bin' Makefile
  make
}

package() {
  cd "${srcdir}/${pkgname}-${pkgver}"

  make libexecdir=/usr/bin prefix=/usr DESTDIR=${pkgdir} install
}
