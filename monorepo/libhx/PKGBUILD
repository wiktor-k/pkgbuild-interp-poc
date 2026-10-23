# Maintainer: George Hu <integral@archlinux.org>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: Max Roder <maxroder@web.de>
# Contributor: Nathan Owe <ndowens.aur at gmail dot com>

pkgname=libhx
_srcname=libHX
pkgver=5.4
pkgrel=1
pkgdesc='A library providing queue, tree, I/O and utility functions'
arch=(x86_64)
url="https://inai.de/projects/${pkgname}/"
license=(GPL-3.0-only LGPL-2.1-or-later)
depends=(glibc)
source=(https://inai.de/files/${pkgname}/${_srcname}-${pkgver}.tar.{zst,asc})
sha256sums=('3b169f4b0d3ddd888033b5816f54e7438b991c465197640c37d9f9dd5d926595'
            'SKIP')
validpgpkeys=('2BDA563E9DDE46AF71CC171DAEB6941A2F9C9944')

prepare() {
	cd "${_srcname}-${pkgver}/"
	autoreconf -vi
}

build() {
	cd "${_srcname}-${pkgver}/"
	./configure --prefix=/usr
	make
}

check() {
	make -C "${_srcname}-${pkgver}" check
}

package() {
	DESTDIR="${pkgdir}" make -C "${_srcname}-${pkgver}" install
}
