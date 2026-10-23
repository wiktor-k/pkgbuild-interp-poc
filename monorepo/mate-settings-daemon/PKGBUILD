# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-settings-daemon
pkgver=1.28.0
pkgrel=3
pkgdesc="The MATE Settings daemon"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-settings-daemon'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  dbus
  dbus-glib
  dconf
  fontconfig
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libcanberra
  libgcc
  libmatekbd
  libmatemixer
  libnotify
  libpulse
  libx11
  libxi
  libxklavier
  mate-desktop
  nspr
  nss
  polkit
)
makedepends=(
  git
  glib2-devel
  mate-common
)
groups=(mate)
source=("git+https://github.com/mate-desktop/mate-settings-daemon.git#tag=v$pkgver")
b2sums=(a47dd2ba4efc3043de52d858228e5084bd0e240628e4c10566818eaa2ae8e22e54c0a43f2e2df0411f78acf6ed276897adab20648322b004ff95e1ca49b3d839)

prepare() {
  cd $pkgname
  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --libexecdir="/usr/lib/$pkgname" \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-pulse
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
