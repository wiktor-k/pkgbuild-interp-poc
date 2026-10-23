# Maintainer: Jonas Witschel <diabonas@archlinux.org>
pkgname=luksmeta
pkgver=10
pkgrel=1
pkgdesc='Library for storing metadata in the LUKSv1 header'
arch=('x86_64')
url='https://github.com/latchset/luksmeta'
license=('LGPL-2.1-or-later')
depends=('cryptsetup' 'libcryptsetup.so')
makedepends=('asciidoc')
source=("$url/releases/download/v$pkgver/$pkgname-$pkgver.tar.bz2")
sha512sums=('03994598c7d573492eea571677d0e9b61413fc305f5a8efb04985c38e727534c6b484f847cd0dd96a53b3d2982399441bf2e269860dd0bc12bcf690d5829d93b')

build() {
	cd "$pkgname-$pkgver"
	./configure --prefix=/usr
	make
}

check() {
	cd "$pkgname-$pkgver"
	make check
}

package() {
	cd "$pkgname-$pkgver"
	make DESTDIR="$pkgdir" install
}
