# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Lex Black <autumn-wind@web.de>
# Contributor: Alexander F Rødseth <xyproto@archlinux.org>
# Contributor: Thomas S Hatch <thatch45@gmail.com>
# Contributor: Corrado 'bardo' Primier <corrado.primier@mail.polimi.it>
# Contributor: Jochem Kossen <j.kossen@home.nl>
# Contributor: Daniel J Griffiths <ghost1227@archlinux.us>

pkgname=scite
pkgver=5.6.7
pkgrel=1
pkgdesc='Editor with facilities for building and running programs'
arch=(x86_64)
url='https://www.scintilla.org/SciTE.html'
license=(HPND)
depends=(
  at-spi2-core
  cairo
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libgcc
  libstdc++
  pango
)
backup=(usr/share/scite/SciTEGlobal.properties)
source=(
  "https://www.scintilla.org/${pkgname}${pkgver//./}.tgz"
  SciTE.appdata.xml
)
b2sums=(
  c4b9512298f57cd42a5c882224dd9a8cf625f0c05c7f82b13e688b0fd06ea112e86395ce67c2aee14696579374c4d593c4c6acef001c9536e91da0dcaa014430
  684e8072881bb507dcdb47159ce69b0c9de017512f80d2bce1efc721ff73afdd503dac69c9e8f8ba9c55ce7dac68a6cf7815b4d7c2638c1a18506ad51edfd58d
)

prepare() {
  sed -i 's/xdg-open/gio open/' scite/src/*.properties
  sed -i 's|pixmapdir=$(datadir)/pixmaps|pixmapdir=$(datadir)/icons/hicolor/48x48/apps|
          s|pixmapdir)/Sci48M.png|pixmapdir)/SciTE.png|' scite/gtk/makefile
  sed -i 's/Sci48M/SciTE/' scite/gtk/{SciTEGTK.cxx,SciTE.desktop}
}

build() {
  make -C lexilla/src
  make -C scintilla/gtk GTK3=1
  make -C scite/gtk GTK3=1
}

package() {
  make -C scite/gtk DESTDIR="$pkgdir" GTK3=1 install
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname/" scite/License.txt
  install -Dm644 -t "$pkgdir/usr/share/metainfo/" SciTE.appdata.xml
  ln -s SciTE "$pkgdir/usr/bin/scite"
}
