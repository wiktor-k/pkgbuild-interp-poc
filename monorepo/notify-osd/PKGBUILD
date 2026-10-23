# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Johannes Pfau <johannespfau at gmail dot com>

pkgname=notify-osd
pkgver=0.9.35+20.04.20191129
pkgrel=2
pkgdesc="Canonical's on-screen-display notification agent, implementing the freedesktop.org Desktop Notifications Specification with semi-transparent click-through bubbles"
arch=(x86_64)
url="https://launchpad.net/notify-osd"
license=(GPL-3.0-only)
depends=(
  at-spi2-core
  cairo
  dbus
  dbus-glib
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gsettings-desktop-schemas
  gtk3
  libwnck3
  libx11
  pango
  pixman
)
makedepends=(
  glib2-devel
  libnotify
)
source=(
  "https://launchpad.net/ubuntu/+archive/primary/+files/${pkgname}_$pkgver.orig.tar.gz"
  notify-osd-fix-types.patch
)
b2sums=(
  719d6c62c7bd0e14996acb03d0f1703b9a9ad7adf717a92403c9e0fd571c0d2c5fe58e1b7c2f1f47d7cf6d43524b0d3b5708875fb275486a6d286d77520f5f1d
  5c0c200b39d9bc1f5ff219ce8b536aa35bb3a0a3e514f43e176376ca15aa4ccbd54d4d8023375c1f84e6a8421e894a7cb4efaaa74d3cc09fd24926f9a7c24d1e
)

prepare() {
  # Fix types
  patch -Np1 -i notify-osd-fix-types.patch

  # Remove gnome-common dependency
  sed -i '/GNOME_/d' configure.in

  autoreconf -fiv
}

build() {
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --libexecdir=/usr/lib/$pkgname \
    --disable-schemas-compile \
    CFLAGS="$CFLAGS -std=gnu17"
  make
}

package() {
  make DESTDIR="$pkgdir" install

  # Remove readme file from non-standard location
  rm "$pkgdir/usr/share/notify-osd/icons/hicolor/scalable/status/README"
}
