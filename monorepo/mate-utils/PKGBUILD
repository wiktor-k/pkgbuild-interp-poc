# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-utils
pkgver=1.28.3
pkgrel=1
pkgdesc="Common MATE utilities for viewing disk usage, logs and fonts, taking screenshots, managing dictionaries and searching files"
arch=(x86_64)
url='https://github.com/mate-desktop/mate-utils'
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
  libcanberra
  libgcc
  libgtop
  libice
  libsm
  libx11
  libxext
  mate-desktop
  pango
  udisks2
  zlib
)
makedepends=(
  git
  glib2-devel
  gtk-doc
  mate-common
  mate-panel
  yelp-tools
)
optdepends=('mate-panel: Look up words in dictionaries from the panel')
groups=(mate-extra)
source=(
  "git+https://github.com/mate-desktop/mate-utils.git#tag=v$pkgver"
  git+https://github.com/mate-desktop/mate-submodules.git
  mate-utils-fix-window-size.patch
)
b2sums=(
  661eac31992be2b44909bf14acad27154ec02960dc3e1d7728022ac44deac7b5422a4cf04a1ffdeebd9b5a185ac130abc9aa75464914f079c0c5c161f5a627e4
  SKIP
  0570eca0cc99d63cc630b961b5a040e76724b4b6544d3c97f0b463d61a0a264b4632bdd0e3c36ad582c98fe9fb865de7f4a422a4339536e8373f4c317d8c7194
)

prepare() {
  cd $pkgname

  git submodule init
  git config submodule.gsearchtool/mate-submodules.url "$srcdir/mate-submodules"
  git -c protocol.file.allow=always submodule update

  # Fix window size growing with client-side decorations
  # https://github.com/mate-desktop/mate-utils/pull/397
  git apply -3 ../mate-utils-fix-window-size.patch

  # Remove OnlyShowIn from desktop files
  # https://github.com/mate-desktop/mate-utils/pull/398
  sed -i '/^OnlyShowIn=/d' \
    baobab/data/mate-disk-usage-analyzer.desktop.in.in \
    gsearchtool/data/mate-search-tool.desktop.in \
    logview/data/mate-system-log.desktop.in.in \
    mate-dictionary/data/mate-dictionary.desktop.in.in \
    mate-disk-image-mounter/data/mate-disk-image-mounter.desktop.in

  # Set NoDisplay=true in mate-disk-image-mounter.desktop
  # https://github.com/mate-desktop/mate-utils/pull/399
  echo 'NoDisplay=true' >> mate-disk-image-mounter/data/mate-disk-image-mounter.desktop.in

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-gtk-doc \
    --enable-in-process
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
