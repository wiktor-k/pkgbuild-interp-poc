# Maintainer: Bruno Pagani <archange@archlinux.org>

pkgname=osm-gps-map
pkgver=1.2.2
pkgrel=1
pkgdesc="Gtk Widget for Displaying OpenStreetMap tiles"
arch=(x86_64)
url="https://github.com/nzjrs/osm-gps-map"
license=(GPL-2.0-or-later)
depends=(gtk3 libsoup3)
makedepends=(gtk-doc gobject-introspection)
source=(https://github.com/nzjrs/osm-gps-map/releases/download/${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha256sums=('051383588f0fdb60f59047c6559110c35025ea72b6eca700af532c020a6b55ec')

prepare() {
  cd ${pkgname}-${pkgver}

  autoreconf -fiv
}

build() {
  cd ${pkgname}-${pkgver}
  ./configure --prefix=/usr --enable-gtk-doc
  make
}

package() {
  cd ${pkgname}-${pkgver}
  make DESTDIR="${pkgdir}" install
}
