# Maintainer: Johannes Löthberg <johannes@kyriasis.com>
# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: hydro <hydro@freenet.de>

pkgbase=mediainfo
pkgname=(mediainfo mediainfo-gui)
pkgver=26.10
pkgrel=1
pkgdesc='Supplies technical and tag information about media files'
arch=(x86_64)
url='https://mediaarea.net'
license=(BSD-2-Clause)
depends=(
  "libmediainfo=$pkgver"
  libzen
  libstdc++
  libgcc
  glibc
)
makedepends=(wxwidgets-gtk3)
source=("$pkgname-$pkgver.tar.gz::https://github.com/MediaArea/MediaInfo/archive/v$pkgver.tar.gz")
sha512sums=('9233defc6d77e639537d4217745556f1ca1369395651c728213cf0a087dc35f69f2f7630134bc0ddd51d1521f600e09985167f3bc98d9237780b5aea425874c0')
b2sums=('6a731b771bcc4370f9974d1d691a1a6615a5c5d6448fe049f8537284750270a096ad241bd042be22016964fc99bf040235cc317aa984409612934ce51926e268')

prepare() {
  cd MediaInfo-$pkgver
# Install service menus in modern path
  sed -e 's|kservices5/ServiceMenus|kio/servicemenus|g' -i Project/GNU/GUI/Makefile.am
}

build() {
  cd "MediaInfo-$pkgver"

  pushd Project/GNU/CLI
  ./autogen.sh
  ./configure --prefix=/usr
  make
  popd

  pushd Project/GNU/GUI
  ./autogen.sh
  ./configure --prefix=/usr
  make
  popd
}

package_mediainfo() {
  pkgdesc+=' (CLI interface)'

  cd "MediaInfo-$pkgver/Project/GNU/CLI"

  make DESTDIR="$pkgdir" install

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" "$srcdir/MediaInfo-$pkgver/LICENSE"
}

package_mediainfo-gui() {
  pkgdesc+=' (GUI interface)'
  depends+=(
    wxwidgets-common
    wxwidgets-gtk3
    hicolor-icon-theme
  )

  cd "MediaInfo-$pkgver/Project/GNU/GUI"

  make DESTDIR="$pkgdir" install

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" "$srcdir/MediaInfo-$pkgver/LICENSE"
}
