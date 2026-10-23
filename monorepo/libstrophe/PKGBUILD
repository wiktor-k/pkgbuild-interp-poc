# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Bartłomiej Piotrowski <bpiotrowski@archlinux.org>

pkgbase=libstrophe
pkgname=('libstrophe' 'libstrophe-doc')
pkgver=0.14.0
pkgrel=1
epoch=1
pkgdesc='Simple, lightweight C library for writing XMPP clients'
url='http://strophe.im/libstrophe/'
arch=('x86_64')
license=('GPL-3.0-only' 'MIT')
depends=('glibc' 'openssl' 'expat' 'zlib')
makedepends=('doxygen')
checkdepends=('check')
source=(https://github.com/strophe/libstrophe/archive/${pkgver}/${pkgname}-${pkgver}.tar.gz
        https://github.com/strophe/libstrophe/commit/dfb3e868.patch)
sha512sums=('d3764044d03ad3b74013fbc08ac717da0c7e140b9fd1c9269a7956fad1fb0f642be48fc1b4c363e504fe48b3f6540c6577c5cac9b7c3a88df6aa23843dd931d7'
            'ef622617f68aa87cc379e5312baa2a02ef8057f311ab1d9215ffa2765d32f5d3f5374eca15bfd98bd947a76290f32ebf831d89e54098b690c9014c0919ca1159')
b2sums=('0fb443969d3f3c8d855be8953b944f335c3d958967acf39cf30a55daf67d56b85a1a994c7deae44788e79336e8213c6e4d4b6d99ea9c7d40f19c89cbd3e2f880'
        '738af35808a452712ed8124de6938022fa442ae287e095cd432d38139641344685e073f1c1e21bd99f154eb583965984556d07653caf01fbdc0177908ac1b333')

prepare() {
  cd ${pkgbase}-${pkgver}
# Fix build with GCC 15
  patch -p1 -i ../dfb3e868.patch
  autoreconf -fiv
}

build() {
  cd ${pkgbase}-${pkgver}
  ./configure \
    --prefix=/usr
  make
  doxygen
}

check() {
  make -C ${pkgbase}-${pkgver} check
}

package_libstrophe() {
  depends+=('libexpat.so')
  provides=('libstrophe.so')
  cd ${pkgbase}-${pkgver}
  make DESTDIR="${pkgdir}" install
  install -Dm 644 MIT-LICENSE.txt "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE-MIT"
}

package_libstrophe-doc() {
  pkgdesc+=' (documentation and samples)'
  depends=()
  options=('!strip')
  optdepends=('libstrophe')
  cd ${pkgbase}-${pkgver}
  install -Dm 644 docs/html/* -t "${pkgdir}/usr/share/doc/${pkgname}/html"
  install -Dm 644 examples/{README.md,*.c} -t "${pkgdir}/usr/share/doc/${pkgname}/examples"
  install -Dm 644 MIT-LICENSE.txt -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
