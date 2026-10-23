# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-icon-theme
pkgver=1.28.0
pkgrel=3
pkgdesc='MATE icon theme'
url="https://github.com/mate-desktop/mate-icon-theme"
arch=(any)
license=('LGPL-3.0-only OR CC-BY-SA-3.0')
options=(!emptydirs)
makedepends=(git)
groups=(mate)
source=("git+https://github.com/mate-desktop/mate-icon-theme.git#tag=v$pkgver")
b2sums=(e2c30b5db8b7a7b4f1c3e072205c0735eaf36858ec417f70c58a7769657598a8598dbc0a2c6a32bd2bead0f98cfccadc3cb5fdaa22831cb9cdecdc936b146a81)

prepare() {
  cd $pkgname
  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --disable-icon-mapping
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
