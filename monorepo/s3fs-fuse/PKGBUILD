# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: David Roheim <david dot roheim at gmail dot com>
# Contributor: Chris Heien <chris.h.heien@gmail.com>
# Contributor: Jonny Gerold <jonny@fsk141.com>

pkgname=s3fs-fuse
pkgver=1.97
pkgrel=1
pkgdesc="FUSE-based file system backed by Amazon S3"
arch=('x86_64')
url="https://github.com/s3fs-fuse/s3fs-fuse/wiki"
license=('GPL-2.0-or-later')
depends=('fuse3' 'curl' 'libxml2')
conflicts=('s3fs')
provides=('s3fs')
replaces=('s3fs')
source=("$pkgname-$pkgver.tar.gz::https://github.com/s3fs-fuse/s3fs-fuse/archive/v$pkgver.tar.gz")
sha512sums=('24b70877a5c162252f08e55176a9e9fa2fd380ed3bbe8f51306f0e1489038a80cbd2a07699689431e3294661fa8d0f2434e7fc94cd14e6c9b4b8fa8be2e1f8a7')

build() {
  cd "$srcdir/$pkgname-$pkgver"

  ./autogen.sh
  ./configure --prefix=/usr
  make
}

check() {
  make -C "$pkgname-$pkgver/src" check
}

package() {
  cd "$srcdir/$pkgname-$pkgver"
  make DESTDIR="$pkgdir/" install
}
