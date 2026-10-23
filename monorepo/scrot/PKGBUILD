# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Maintainer: T.J. Townsend <blakkheim@archlinux.org>
# Contributor: Daniel Wallace <danielwallace at gtmanfred dot com>
# Contributor: Giovanni Scafora <giovanni@archlinux.org>

pkgname=scrot
pkgver=2.0.0
pkgrel=1
pkgdesc='Simple command-line screenshot utility for X'
url='https://github.com/resurrecting-open-source-projects/scrot'
arch=('x86_64')
license=('MIT')
depends=('imlib2' 'libxext' 'libxfixes' 'libxcomposite' 'libbsd' 'xorg-xrandr')
makedepends=('autoconf-archive' 'git')
source=("git+https://github.com/resurrecting-open-source-projects/scrot.git#tag=${pkgver}")
sha512sums=('407e9edabd1c7ecc5fdb6f7a08d910f18f6f76b881c4b9906c41715465f9877229820be80b27774bb71f887d30a6a55c5d884168cc73154f8708a0e326833fc3')
b2sums=('71213d3ccd85349014594c54376d480164182c7cd02176404116a31a19a1276ffeb819928a7672ad85cb9341200e328ca45508c806f94188f911d71adcb07953')

prepare(){
  cd ${pkgname}
  autoreconf -fiv
}

build() {
  cd ${pkgname}
  ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man
  make
}

package() {
  cd ${pkgname}
  make DESTDIR="${pkgdir}" install
  install -Dm644 README.md ChangeLog -t "${pkgdir}"/usr/share/doc/${pkgname}
  install -Dm644 COPYING -t "${pkgdir}"/usr/share/licenses/${pkgname}
}

# vim: ts=2 sw=2 et:
