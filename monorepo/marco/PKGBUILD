# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=marco
pkgver=1.29.0
pkgrel=2
pkgdesc="MATE default window manager"
arch=(x86_64)
url='https://github.com/mate-desktop/marco'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  libcanberra
  libgtop
  libice
  libsm
  libx11
  libxcomposite
  libxcursor
  libxdamage
  libxext
  libxfixes
  libxinerama
  libxpresent
  libxrandr
  libxrender
  libxres
  mate-desktop
  mate-settings-daemon
  pango
  startup-notification
  zenity
)
makedepends=(
  git
  mate-common
  yelp-tools
)
groups=(mate)
source=(
  "git+https://github.com/mate-desktop/marco.git#tag=v$pkgver"
  marco-zenity-fix.patch
)
b2sums=(
  a7c96ed0a679ba9d6ed178b6e452d24537a76f71e73d221b06a85e9f5328dc9c4b6f04cd1074c5637dfd9f82f2dd511a25e7f0b6d14a68ef5feef62a563d152e
  850c0c63ca1a14e8a0c112d70bde9f7dd8905ffdecfd8a8c5b206d61515f540e5beecbc7a042080e33b7d7800175ec6a9966a6069abdbc083c4d839235eb1fbe
)

prepare() {
  cd $pkgname

  # Change default icon from preferences-desktop-theme to application-x-executable
  # https://github.com/mate-desktop/marco/pull/836
  git cherry-pick -n 2db9ab174696a9ea88ad1bfd0043fdf6ca5db489

  # Remove unavailable zenity options
  # https://github.com/mate-desktop/marco/pull/845
  git apply -3 ../marco-zenity-fix.patch

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
