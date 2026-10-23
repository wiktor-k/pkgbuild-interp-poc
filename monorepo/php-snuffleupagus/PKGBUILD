# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>
# Contributor: Julien (jvoisin) Voisin <julien.voisin+snuffleupagus@dustri.org>

pkgname=php-snuffleupagus
_pkgname=snuffleupagus
pkgver=0.14.0
pkgrel=1
pkgdesc="Security module for PHP"
url='https://snuffleupagus.readthedocs.io/'
arch=('x86_64')
license=('LGPL-3.0-only')
depends=(
  'glibc'
  'pcre2' 'libpcre2-8.so'
  'php'
)
makedepends=('re2c')
backup=('etc/php/conf.d/snuffleupagus.ini'
        'etc/php/conf.d/snuffleupagus.rules')
source=(https://github.com/jvoisin/${_pkgname}/archive/v${pkgver}/${_pkgname}-${pkgver}.tar.gz)
sha256sums=('080cf7e24d15a8650e271837030fca546e627c7a4c7317710c683282fdcb71c6')
b2sums=('9062306e10c1c05c552d73037300f7050bdb2f4b0ecd64f14b9a01aace7c93c0a0134f8d3d40a3c79e9bd50623246aa53ad5c7c35cb32c28435617d5213e5353')

prepare() {
  cd ${_pkgname}-${pkgver}/src
  phpize
  ./configure --prefix=/usr --enable-"${_pkgname}"
}

build() {
  cd ${_pkgname}-${pkgver}/src
  make
}

check() {
  cd ${_pkgname}-${pkgver}/src
  NO_INTERACTION=1 REPORT_EXIT_STATUS=1 make test
}

package() {
  cd ${_pkgname}-${pkgver}
  make -C src INSTALL_ROOT="${pkgdir}" install
  install -D -m644 "config/default.rules" "${pkgdir}/etc/php/conf.d/${_pkgname}.rules"
  install -D -m644 "config/${_pkgname}.ini" "${pkgdir}/etc/php/conf.d/${_pkgname}.ini"
}

# vim: ts=2 sw=2 et:
