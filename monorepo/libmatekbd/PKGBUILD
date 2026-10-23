# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=libmatekbd
pkgver=1.28.0
pkgrel=3
pkgdesc="MATE keyboard library"
arch=(x86_64)
url='https://github.com/mate-desktop/libmatekbd'
license=(LGPL-2.0-or-later)
depends=(
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  libx11
  libxklavier
  pango
)
makedepends=(
  git
  glib2-devel
  gobject-introspection
  mate-common
)
source=("git+https://github.com/mate-desktop/libmatekbd.git#tag=v$pkgver")
b2sums=(b0a71c4fdfa1477e22c5427a6693e2b8111c1feaac71773bfe8d75bf260950ccbfb5bbf387219bf6c9b5a95b47ba934e8638695f569716eab98911f3bc85f66a)

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
