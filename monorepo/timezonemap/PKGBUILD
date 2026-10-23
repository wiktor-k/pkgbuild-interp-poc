# Contributor: Eli Schwartz <eschwartz@archlinux.org>

pkgname=timezonemap
pkgver=0.4.5.4
pkgrel=1
pkgdesc='GTK3 timezone map widget'
arch=(x86_64)
url='https://codeberg.org/dashea/timezonemap'
license=(GPL-3.0-only)
depends=(
  cairo
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  json-glib
  librsvg
  libsoup3
)
makedepends=(
  git
  gobject-introspection
)
source=("git+https://codeberg.org/dashea/timezonemap.git#tag=$pkgver")
b2sums=(36209220b8b68e597cb17f6abb709737a85186ec9c0dbf461c3f5ba8fbbf2493ddf12f6ce088222922a6ce6c8b2201d5f640cdcd6f50222ceb533eb588f44ec6)

prepare() {
  cd $pkgname
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
