# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-applets
pkgver=1.28.1
pkgrel=3
pkgdesc="Applets for MATE panel"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-applets'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  cpupower
  dbus-glib
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libgtop
  libmateweather
  libnl
  libnotify
  libwnck3
  libx11
  libxml2
  mate-desktop
  mate-panel
  pango
  polkit
  upower
)
makedepends=(
  git
  glib2-devel
  mate-common
  yelp-tools
)
groups=(mate-extra)
source=("git+https://github.com/mate-desktop/mate-applets.git#tag=v$pkgver")
b2sums=(d7a0e451f38d3112168789160622e941636d097be72f0ed5ee6544b0f238d1f15738ddeb3389e88b9b7f434675211222ebecbfd96f7a2012c7b8d063d5ae1c81)

prepare() {
  cd $pkgname

  # Geyes in-process: follow the mouse properly no matter where the applet is
  # https://github.com/mate-desktop/mate-applets/pull/675
  git cherry-pick -n e8e8d00cff2ade8e034681b2a3c8f5dd83fbb331

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-in-process
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
