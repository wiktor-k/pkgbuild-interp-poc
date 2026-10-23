# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-power-manager
pkgver=1.28.1
pkgrel=3
pkgdesc="Power management tool for the MATE desktop"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-power-manager'
license=(GPL-2.0-or-later)
depends=(
  cairo
  dbus
  dbus-glib
  dconf
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libcanberra
  libnotify
  libsecret
  libx11
  libxext
  libxrandr
  mate-desktop
  pango
  polkit
  upower
)
makedepends=(
  git
  glib2-devel
  mate-common
  mate-panel
  yelp-tools
)
optdepends=('mate-panel: Set brightness and inhibit power management from the panel')
source=("git+https://github.com/mate-desktop/mate-power-manager.git#tag=v$pkgver")
groups=(mate-extra)
b2sums=(b32b7dda930515c2ea35b90134e2eee977b7b00734a2f6a629561e5cc3f9bf77ec96f3b9cec135e87355d388d1bc165da0b54cc39bcd4734d958123d19259914)

prepare() {
  cd $pkgname
  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sbindir=/usr/bin \
    --libexecdir="/usr/lib/$pkgname" \
    --sysconfdir=/etc \
    --localstatedir=/var
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
