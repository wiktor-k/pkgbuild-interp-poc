# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=atril
pkgver=1.28.7
pkgrel=3
pkgdesc="A document viewer for MATE"
arch=(x86_64)
url='https://github.com/mate-desktop/atril'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libarchive
  libgepub
  libice
  libsecret
  libsm
  libsynctex
  libtiff
  libxml2
  mate-desktop
  pango
  poppler-glib
  webkit2gtk-4.1
)
makedepends=(
  caja
  djvulibre
  git
  glib2-devel
  gobject-introspection
  gtk-doc
  libgxps
  libspectre
  mate-common
  texlive-bin
  yelp-tools
)
optdepends=(
  'caja: Document Properties extension'
  'djvulibre: Djvu support'
  'libgxps: XPS support'
  'libspectre: DVI and PS support'
  'texlive-bin: DVI support'
)
groups=(mate-extra)
source=(
  "git+https://github.com/mate-desktop/atril.git#tag=v$pkgver"
  git+https://github.com/mate-desktop/mate-submodules.git
  atril-rotation-pixbuf-cache.patch
)
b2sums=(
  2a75a37d772a8dee7e4ef2f761d1bff0542201c1e9a3901a39c698fe23cb64bd38b45ba6b0cf24413669dc58189ff43bdedb80fbe76cfc752bf75d72cd9a722b
  SKIP
  c1b6ca09db1cf4a17fcf6fca820507d0dc8222e668c749dbdaa9c6f5569e84c9b2436fa7e4d221133ec807928051e4c123fd6d86bc51e74260fc068a9fa1d47e
)

prepare() {
  cd $pkgname

  git submodule init
  git config submodule.mate-submodules.url "$srcdir/mate-submodules"
  git -c protocol.file.allow=always submodule update

  # Implement text interface for djvu
  # https://github.com/mate-desktop/atril/pull/726
  git cherry-pick -n a049ebfdfdeeac2ebb611c2b5db697987f02c6dd

  # Check pixbuf cache before clearing selection
  # https://github.com/mate-desktop/atril/pull/729
  git apply -3 ../atril-rotation-pixbuf-cache.patch

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --libexec="/usr/lib/$pkgname" \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-pixbuf \
    --enable-introspection \
    --enable-gtk-doc
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
