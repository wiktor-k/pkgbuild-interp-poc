# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Contributor: Bartłomiej Piotrowski <bpiotrowski@archlinux.org>
# Contributor: naelstrof <naelstrof@gmail.com>

pkgname=slop
pkgver=7.7
pkgrel=4
pkgdesc='Utility to query the user for a selection and print the region to stdout'
arch=('x86_64')
url='https://github.com/naelstrof/slop'
license=('GPL3')
depends=('libxext' 'libx11' 'mesa' 'glew' 'libxrender')
makedepends=('cmake' 'glm')
source=("${pkgname}-${pkgver}.tar.gz::$url/archive/v${pkgver}.tar.gz")
sha256sums=('a69a6e5c41d7fff1c6aa35b367a5c5a6dc98e621fa9a1908808d6308c2b40f4e')

prepare() {
	cd "$pkgname-$pkgver"
        sed -e 's|CXX_STANDARD 11|CXX_STANDARD 17|g' -i CMakeLists.txt # Fix build with ICU 74
}

build() {
	cd "$pkgname-$pkgver"
	cmake -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_INSTALL_LIBDIR=/usr/lib .
	make
}

package() {
	make -C "$pkgname-$pkgver" DESTDIR="$pkgdir" install
}
