# Maintainer: Christian Heusel <christian@heusel.eu>
# Contributor: Sanjay Pavan <withercubes@protonmail.com>

pkgname=nsxiv
pkgver=34
pkgrel=1
pkgdesc='Neo (or New or Not) Simple (or Small or Suckless) X Image Viewer'
arch=('x86_64')
license=('GPL-2.0-or-later')
url='https://nsxiv.codeberg.page/'

depends=(
  'glibc'
  'imlib2' 'libx11'                 # core dependencies
  'libxft' 'fontconfig' 'freetype2' # status bar
  'libexif'                         # display exif images
  'hicolor-icon-theme'              # make icon
)

source=("$pkgname-$pkgver.tar.gz"::"https://codeberg.org/nsxiv/nsxiv/archive/v$pkgver.tar.gz")
sha256sums=('b68356296130924d56abc0616756900c0fe75ac9932b9bb9bed3a349915b7320')

prepare() {
  cd "$pkgname"
  [ ! -f config.h ] && cp config.def.h config.h
  # make install currently also runs make all
  # see https://codeberg.org/nsxiv/nsxiv-record/issues/237
  sed -i -e '/^install: / s|: all|:|' Makefile
}

build() {
  echo "RUSTFLAGS: $RUSTFLAGS"
  echo "Debug RUSTFLAGS: $DEBUG_RUSTFLAGS"
  make -C "$pkgname" HAVE_INOTIFY=1 HAVE_LIBFONT=1 HAVE_LIBEXIF=1
}

package() {
  make -C "$pkgname" PREFIX=/usr DESTDIR="$pkgdir" install-all
}
