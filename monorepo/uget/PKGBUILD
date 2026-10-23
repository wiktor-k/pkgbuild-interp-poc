# Maintainer: Balló György <ballogyor+arch at gmail dot com>

pkgname=uget
pkgver=2.2.3
pkgrel=8
pkgdesc='GTK download manager featuring download classification and HTML import'
arch=(x86_64)
url='https://sourceforge.net/projects/urlget/'
license=(LGPL-2.1-or-later)
depends=(
  cairo
  curl
  gdk-pixbuf2
  glib2
  glibc
  gstreamer
  gtk3
  hicolor-icon-theme
  libappindicator
  libnotify
  openssl
  pango
)
makedepends=(intltool)
optdepends=('aria2: alternative backend')
source=(
  "https://downloads.sourceforge.net/urlget/$pkgname-$pkgver-1.tar.gz"
  uget-disable-rss.patch
  uget-gcc10.patch
  uget-gcc16.patch
  uget-download-dir.patch
  uget-gtk.appdata.xml
)
b2sums=(
  1b599c9c5ba4cb5929c211c35382b9dbc6940e52bc995f150632dfc52f1429d26e15d330f53d1da9bd1ef296368cf3877673ee3fb5ab3f33d5ded4cf86b7b3d2
  bf4358ee0de68e4869a279e94885478c10149868738604162b11e6381868e7b5a4ac21778f87b1e472688bebc46ceed723305c9ded4d251e32e52f0941ab6d64
  069a6339274d4cc2532e2d7cb752e21e8e3b0ba3566231c8993a11b956d35773e9afa4b32336f6a8085736a44a261601a065df5fcdb8bbbacd6526b82e62e0ce
  71183036a1050a054e783afb9b3c99b90134bc32128735bf6aeba62fa9642ebdd07e2512dac482787234b31acc844bc552eba9041efd55b182e9474ba089f69b
  36c095096671e437bcb639dd5c14f614b7942724000430de9edcd21438d37e918e8998e4eda610157c05654103b9c0d22d103ca90c043b0129be5f08de339a20
  613579e179a0918dfeba1ea8c5f5da724ae2f20deb4e7af6bddc38497e14403084a9eb69f496810854d337a07e86feb65b6acfb82281961b13f25e06d5964ffa
)

prepare() {
  cd $pkgname-$pkgver

  # Disable checking nonexistent RSS feeds
  patch -Np1 -i ../uget-disable-rss.patch

  # Fix build with GCC 10
  patch -Np1 -i ../uget-gcc10.patch

  # Fix build with GCC 16
  patch -Np1 -i ../uget-gcc16.patch

  # Use download dir as default
  patch -Np1 -i ../uget-download-dir.patch

  # Don't force x11 backend (FS#70837)
  sed -i 's/env GDK_BACKEND=x11 //' uget-gtk.desktop
}

build() {
  cd $pkgname-$pkgver
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --disable-rss-notify
  make
}

package() {
  cd $pkgname-$pkgver
  make DESTDIR="$pkgdir" install
  install -Dm644 ../uget-gtk.appdata.xml "$pkgdir/usr/share/metainfo/uget-gtk.appdata.xml"
}
