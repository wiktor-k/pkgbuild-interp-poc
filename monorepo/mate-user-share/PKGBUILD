# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-user-share
pkgver=1.28.0
pkgrel=3
pkgdesc="User level public file sharing via WebDAV for MATE"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-user-share'
license=(GPL-2.0-or-later)
depends=(
  dconf
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libx11
  mod_dnssd
)
makedepends=(
  caja
  dbus-glib
  git
  glib2-devel
  libcanberra
  libnotify
  mate-common
  yelp-tools
)
optdepends=('caja: File sharing extension')
groups=(mate-extra)
source=("git+https://github.com/mate-desktop/mate-user-share.git#tag=v$pkgver")
b2sums=(1a3c3adc09ca732eb25991433282ef16cdc0be1e27c628e6362694ac2ce50dbbaec200e5b431e58110f7088249fedf2ac13d1384c3207a48b2699c90216a4735)

prepare() {
  cd $pkgname
  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --libexec="/usr/lib/$pkgname" \
    --sysconfdir=/etc \
    --localstatedir=/var
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
  rm "$pkgdir"/etc/xdg/autostart/mate-user-share-obex*.desktop
}
