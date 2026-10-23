# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=pluma
pkgver=1.28.1
pkgrel=2
pkgdesc="A powerful text editor for MATE"
arch=(x86_64)
url='https://github.com/mate-desktop/pluma'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  dconf
  enchant
  gdk-pixbuf2
  glibc
  gtk3
  gtksourceview4
  iso-codes
  libice
  libpeas
  libsm
  libx11
  libxml2
  mate-desktop
  pango
  zenity
)
makedepends=(
  git
  glib2-devel
  gobject-introspection
  gtk-doc
  mate-common
  yelp-tools
)
optdepends=('python-gobject: for External Tools, Python Console, Quick Open, Snippets plugins')
groups=(mate-extra)
source=(
  "git+https://github.com/mate-desktop/pluma.git#tag=v$pkgver"
  git+https://github.com/mate-desktop/mate-submodules.git
)
b2sums=(
  ad2789a065bd7b211b35f0f2a6a73006e774558312ffc92b3dbd69e2e1c6c58578ff3c358ba1ec11280f6536b2e4927d82ac3265c9d97057033133ccf1e046e7
  SKIP
)

prepare() {
  cd $pkgname

  git submodule init
  git config submodule.mate-submodules.url "$srcdir/mate-submodules"
  git -c protocol.file.allow=always submodule update

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --libexec="/usr/lib/$pkgname" \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-gtk-doc
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
