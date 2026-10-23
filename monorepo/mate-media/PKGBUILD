# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-media
pkgver=1.28.1
pkgrel=4
pkgdesc="MATE Media Tools"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-media'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  libcanberra
  libgcc
  libmatemixer
  libxml2
  mate-desktop
  pango
)
makedepends=(
  git
  mate-common
  mate-panel
)
optdepends=('mate-panel: Control volume from the panel')
groups=(mate-extra)
source=("git+https://github.com/mate-desktop/mate-media.git#tag=v$pkgver")
b2sums=(98ef780ff690ab78ec38bfcb9f843c70616628aee58c976dc5b4fb580cd971f3d4ad4bc939344b2ee72f7fa80ed205eac6257b09d28867893959b005673f5354)

prepare() {
  cd $pkgname
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
