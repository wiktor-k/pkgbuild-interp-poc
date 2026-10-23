# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-terminal
pkgver=1.28.3
pkgrel=1
pkgdesc="The MATE Terminal Emulator"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-terminal'
license=(GPL-3.0-or-later)
depends=(
  at-spi2-core
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  libice
  libsm
  libx11
  mate-desktop
  pango
  perl
  vte3
)
makedepends=(
  git
  glib2-devel
  mate-common
  yelp-tools
)
groups=(mate-extra)
source=(
  "git+https://github.com/mate-desktop/mate-terminal.git#tag=v$pkgver"
  git+https://github.com/mate-desktop/mate-submodules.git
)
b2sums=(
  559b6909c9d40f017ffedac0a925b360a786c362d15cc1175786fc2275a29915645992745f4a7a05d3d8c1f801b9d29b80247fefa033667c94d45d8a9fc84724
  SKIP
)

prepare() {
  cd $pkgname

  git submodule init
  git config submodule.src/mate-submodules.url "$srcdir/mate-submodules"
  git -c protocol.file.allow=always submodule update

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

check() {
  cd $pkgname
  make check
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
