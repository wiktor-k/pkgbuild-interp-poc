# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Marcin Wieczorek <marcin@marcin.co>
# Contributor: ovi chis <ovi@ovios.org>

pkgname=libqb
pkgver=2.0.10
pkgrel=1
pkgdesc='Library for providing high performance, reusable features for client-server architecture'
arch=('x86_64')
url='https://github.com/ClusterLabs/libqb'
license=('LGPL-2.1-or-later')
depends=(
  'glibc'
  'libxml2'
)
makedepends=(
  'doxygen'
  'splint'
)
checkdepends=('check')
provides=('libqb.so')
source=(https://github.com/ClusterLabs/${pkgname}/releases/download/v${pkgver}/${pkgname}-${pkgver}.tar.gz{,.asc})
sha512sums=('c7e833c3bffbd40bd095c0866b870f7357d2960ff1a668bfd6251e46f6914db05c285574316518886ff7fe17c1a3dc990a023a538e521540782d84f011884543'
            'SKIP')
validpgpkeys=(
  'C5E29348A2B634E9F71B2014791890532CB5CDDE' # libqb Release Signing Key <users@clusterlabs.org>
  '746F4C0CA8C2B82ECC0A694808C81B2C6CE95CA7' # Cluster Release Team <cluster-devel@redhat.com>
)

prepare() {
  cd ${pkgname}-${pkgver}
  ./autogen.sh
}

build() {
  cd ${pkgname}-${pkgver}
  ./configure \
    --prefix=/usr \
    --disable-fatal-warnings \
    --disable-static \
    --libdir=/usr/lib \
    --sbindir=/usr/bin \
    --localstatedir=/var
  make
}

check() {
  cd ${pkgname}-${pkgver}
  make -C tests VERBOSE=1 check
}

package() {
  cd ${pkgname}-${pkgver}
  make DESTDIR="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
