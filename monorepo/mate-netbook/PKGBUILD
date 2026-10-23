# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-netbook
pkgver=1.27.0
pkgrel=4
pkgdesc="A simple window management tool"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-netbook'
license=(GPL-3.0-only)
depends=(
  at-spi2-core
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  libfakekey
  libwnck3
  libx11
  mate-panel
  pango
)
makedepends=(
  git
  mate-common
)
groups=(mate-extra)
source=(
  "git+https://github.com/mate-desktop/mate-netbook.git#tag=v$pkgver"
  mate-netbook-fix-undecorating.patch
  mate-netbook-window-activate.patch
)
b2sums=(
  3c7cc791797f8ba10cce4d84e26f72af7b02d089c28a626616cd5d78e110677822930505234d93b1ee9e0ab87a8b46bd70d8d5a4260d8129d73b23e4cf314bb4
  a6a4ce049d8c312d2208dbf6800e75d004aede785837d38d9b44adb51818b1724153df14bc1022ce8b2375497404c67d2cdab5ce09df74ce21b9e8f318943255
  245ece8cdddad29d0705bb0631be5162aafdcacc2241b19b0e49abfeb56f118cde53c18c7176101f4d6971c8aed567c4200080902dc59ec09da26bffbb1a65f6
)

prepare() {
  cd $pkgname

  # Fix undecorate on maximise
  # https://github.com/mate-desktop/mate-netbook/pull/67
  git cherry-pick -n 868f5a402a8c494c7dcc2dd1b8ef1e662818265d

  # Fix handling of client-side decorated windows
  # https://github.com/mate-desktop/mate-netbook/pull/76
  git cherry-pick -n 9c3eb2866e4a4d6f59264db0a7248233935b661a

  # Fix undecorating maximized windows
  # https://github.com/mate-desktop/mate-netbook/pull/77
  git apply -3 ../mate-netbook-fix-undecorating.patch

  # Fix window activation with stale event timestamps
  # https://github.com/mate-desktop/mate-netbook/pull/78
  git apply -3 ../mate-netbook-window-activate.patch

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
}
