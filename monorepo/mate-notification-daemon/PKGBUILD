# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-notification-daemon
pkgver=1.29.0
pkgrel=3
pkgdesc="Notification daemon for MATE"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-notification-daemon'
license=(GPL-2.0-or-later)
groups=(mate)
depends=(
  at-spi2-core
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk-layer-shell
  gtk3
  hicolor-icon-theme
  libcanberra
  libnotify
  libwnck3
  libx11
  libxml2
)
makedepends=(
  git
  glib2-devel
  mate-common
  mate-panel
)
optdepends=('mate-panel: Panel applet')
source=("git+https://github.com/mate-desktop/mate-notification-daemon.git#tag=v$pkgver")
b2sums=(362719e89592ca256aeacd8a7a84597c98b85710a0cb6fd0f47282725dff3ae64796b215237f7074c61d0a68db1f2378248d12175c2aff8d742bf4b22d05c925)

prepare() {
  cd $pkgname

  # Position notification history popup
  # https://github.com/mate-desktop/mate-notification-daemon/pull/242
  git cherry-pick -n 416d6366374971898aa3255b4f24b2d23c246262

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --libexecdir="/usr/lib/$pkgname" \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-in-process
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
