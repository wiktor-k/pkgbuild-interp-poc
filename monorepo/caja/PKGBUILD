# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=caja
pkgver=1.28.0
pkgrel=5
pkgdesc="File manager for the MATE desktop"
arch=(x86_64)
url='https://github.com/mate-desktop/caja'
license=('GPL-2.0-or-later AND LGPL-2.0-or-later')
depends=(
  at-spi2-core
  cairo
  dconf
  exempi
  gdk-pixbuf2
  glib2
  glibc
  gtk-layer-shell
  gtk3
  gvfs
  hicolor-icon-theme
  libexif
  libgcc
  libice
  libnotify
  libsm
  libx11
  libxml2
  mate-desktop
  pango
)
makedepends=(
  git
  glib2-devel
  gobject-introspection
  gtk-doc
  mate-common
)
optdepends=('gstreamer: automatic media playback when mouseover')
groups=(mate)
source=(
  "git+https://github.com/mate-desktop/caja.git#tag=v$pkgver"
  git+https://github.com/mate-desktop/mate-submodules.git
  caja-csd-window-position.patch
)
b2sums=(
  0d45b556f962d9bbfe9803257af5c8114bf8d2ec26d3fb46e05c53ed09c9dd65ef19bd28b290b79e0bba8f21042d1913e1a2451b0eeb7e9e3e62b5b3059e7cce
  SKIP
  1a0971fda0fef3e6c77ea07facda5cbdb112990e8e20218d965136896a9acfc4e14d3c2702e470a4c52232475c663265f9812ab69fa5e3577a4cbbd5be5ddc26
)

prepare() {
  cd $pkgname

  git submodule init
  git config submodule.mate-submodules.url "$srcdir/mate-submodules"
  git -c protocol.file.allow=always submodule update

  # Ensure windows can be moved if compositor is using CSD
  # https://github.com/mate-desktop/caja/pull/1787
  git cherry-pick -n d55e854dade057bb954778d12acdbdc77d8c0422

  # Ensure desktop window is never decorated
  # https://github.com/mate-desktop/caja/pull/1816
  git cherry-pick -n 1cc61bef81c28fba1277cd996c32621b441da068 \
                     2445887b3947eb99bd4d430b4f1dff9ed2467d73

  # Fix desktop window position with client-side decorations
  # https://github.com/mate-desktop/caja/pull/1912
  git apply -3 ../caja-csd-window-position.patch

  # Remove OnlyShowIn from caja-browser.desktop
  # https://github.com/mate-desktop/caja/pull/1913
  sed -i '/^OnlyShowIn=/d' data/caja-browser.desktop.in.in

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-gtk-doc \
    --disable-update-mimedb \
    --disable-self-check # self-check only checks for eel in make check, requires Xorg server
  make
}

check() {
  cd $pkgname
  make check
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
