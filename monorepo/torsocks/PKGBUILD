# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Daniel Micay <danielmicay@gmail.com>
# Contributor: Nicolas Pouillard <nicolas.pouillard@gmail.com>
# Contributor: Rorschach <r0rschach@lavabit.com>

pkgname=torsocks
pkgver=2.5.0
pkgrel=1
pkgdesc='Wrapper to safely torify applications'
url='https://gitlab.torproject.org/tpo/core/torsocks'
arch=('x86_64')
license=(GPL-2.0-or-later)
depends=(
  bash
  glibc
  tor
)
backup=("etc/tor/${pkgname}.conf")
#source=(https://people.torproject.org/~dgoulet/${pkgname}/${pkgname}-${pkgver}.tar.xz{,.asc})
source=("${url}/-/archive/v${pkgver}/torsocks-v${pkgver}.tar.gz")
sha512sums=('dcf4d4d817eb79e52930afff12915eb040b8fc7acd75ec60f28e0d06528cf011b20968ae6ea9bcb3dc51aebe26f88e89fc53d0630487e9216371ed7f71d25773')
validpgpkeys=('B74417EDDF22AC9F9E90F49142E86A2A11F48D36')

build() {
  cd ${pkgname}-v${pkgver}
  autoreconf -vfi
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc
  make
}

package() {
  cd ${pkgname}-v${pkgver}
  make DESTDIR="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
