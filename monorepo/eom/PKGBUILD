# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=eom
pkgver=1.28.1
pkgrel=2
pkgdesc="An image viewer for MATE"
arch=(x86_64)
url='https://github.com/mate-desktop/eom'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  dconf
  exempi
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  lcms2
  libexif
  libjpeg-turbo
  libpeas
  librsvg
  libx11
  libxml2
  mate-desktop
  zlib
)
makedepends=(
  git
  glib2-devel
  gobject-introspection
  gtk-doc
  mate-common
  yelp-tools
)
groups=(mate-extra)
source=("git+https://github.com/mate-desktop/eom.git#tag=v$pkgver")
b2sums=(67e24cc9909df842c802a67ce1fae197c0aa44ae01d39be6096aef2979182329015f1dc3aabf7ae3669fcc504a55e6c599e3bb6d9f54132fc3769dbe52062719)

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
    --enable-gtk-doc
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
