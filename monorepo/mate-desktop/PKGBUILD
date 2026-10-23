# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-desktop
pkgver=1.28.2
pkgrel=3
pkgdesc="Library with common API for various MATE modules"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-desktop'
license=('GPL-2.0-or-later AND LGPL-2.0-or-later')
depends=(
  at-spi2-core
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  iso-codes
  libgcc
  libx11
  libxrandr
  startup-notification
)
makedepends=(
  git
  gobject-introspection
  gtk-doc
  mate-common
)
groups=(mate)
source=("git+https://github.com/mate-desktop/mate-desktop.git#tag=v$pkgver")
b2sums=(6a333d1a6f671c5a1cc262d04db38fc7c9b4c86529c7c6e831e7eaf41bff470b7543e3a2905080de226e76810378d0c7e0510920f3280258aa19a7f6e94f1723)

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
