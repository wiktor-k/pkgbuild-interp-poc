# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: kfgz <kfgz at interia pl>
# Contributor: Adam Wolk <netprobe at gmail dot com>

pkgname=libewf
_pkgname=libewf-legacy
pkgver=20140816
pkgrel=1
pkgdesc="A library for support of the Expert Witness Compression Format (EWF)"
arch=('x86_64')
url="https://github.com/libyal/libewf-legacy"
license=('LGPL-3.0-or-later')
depends=(
  'bzip2'
  'glibc'
  'zlib'
)
makedepends=(
  'fuse2'
  'git'
)
optdepends=('fuse2: to be able to use ewfmount')
source=("$pkgname-$pkgver.tar.gz::$url/archive/$pkgver.tar.gz")
sha256sums=('de94366eb9586a910ae536dc8b621b5d7637ee28913a7fc0e6395b2ab67e9b9b')

prepare() {
  cd $_pkgname-$pkgver
  ./synclibs.sh
  ./autogen.sh
}

build() {
  cd $_pkgname-$pkgver
  ./configure --prefix=/usr --sysconfdir=/etc --without-openssl
  make
}

check() {
  cd $_pkgname-$pkgver
  make check
}

package() {
  cd $_pkgname-$pkgver
  make DESTDIR="$pkgdir" install
}
