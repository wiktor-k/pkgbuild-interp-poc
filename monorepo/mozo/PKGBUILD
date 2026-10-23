# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mozo
pkgver=1.28.0
pkgrel=6
pkgdesc="MATE menu editing tool"
arch=(any)
url='https://github.com/mate-desktop/mozo'
license=(LGPL-2.1-or-later)
depends=(
  gdk-pixbuf2
  glib2
  gtk3
  hicolor-icon-theme
  mate-menus
  mate-panel
  python
  python-gobject
)
makedepends=(git)
groups=(mate-extra)
source=("git+https://github.com/mate-desktop/mozo.git#tag=v$pkgver")
b2sums=(f5022b6df01f999cf34075b128dc3dc54f15c11afaa4dcd2f333459f83843bb3451d3af5f94dea4a655fbaa979db5dcc79daf3383ca9890156b1de3def7918fa)

prepare() {
  cd $pkgname

  # Do not leave behind stale .desktop files
  # https://github.com/mate-desktop/mozo/pull/99
  git cherry-pick -n de820bfd5669af5ff890756444459590a7106267

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
