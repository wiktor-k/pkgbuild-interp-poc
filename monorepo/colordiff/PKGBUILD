# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=colordiff
pkgver=1.0.22
pkgrel=1
pkgdesc="A Perl script wrapper for 'diff' that produces the same output but with pretty 'syntax' highlighting"
arch=('any')
url="https://www.colordiff.org/"
license=('GPL-2.0-or-later')
depends=('perl' 'diffutils')
backup=('etc/colordiffrc' 'etc/colordiffrc-lightbg')
source=(https://www.colordiff.org/archive/${pkgname}-${pkgver}.tar.gz{,.sig})
sha256sums=('f96f73c54521c53f14dc164d5a3920c9ca21a0e5f8e9613f43812a98af3e22af'
            'SKIP')
validpgpkeys=('CF3A93EF01E616C5AE7A1D2745E1E473378BB197') #Dave Ewart <davee@sungate.co.uk>

package() {
  cd ${pkgname}-${pkgver}

  make DESTDIR="${pkgdir}" INSTALL_DIR=/usr/bin MAN_DIR=/usr/share/man/man1 install
  install -m644 colordiffrc-lightbg "${pkgdir}/etc/colordiffrc-lightbg"
}
