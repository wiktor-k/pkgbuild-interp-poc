# Maintainer: George Hu <integral@archlinux.org>
# Contributor: Ronald van Haren <ronald@archlinux.org>
# Contributor: Tobias Powalowski <tpowa@archlinux.org>

pkgname=libnova
pkgver=0.16
pkgrel=3
pkgdesc="A general purpose, double precision, celestial mechanics, astrometry and astrodynamics library."
url="http://libnova.sourceforge.net/"
license=('LGPL-2.0-or-later')
arch=('x86_64')
depends=('glibc')
makedepends=('git')
source=("git+https://git.code.sf.net/p/libnova/libnova#tag=v${pkgver}")
sha512sums=('783fcc8668b5c62fe90ab08c00b6c4ef57386febff3c532bfba9743bb4d1e64aa5dbe5d5bc6b71fb5608adce718585b1fc3eeab593ce833408a622613347e87f')

prepare() {
	cd "${pkgname}/"
	autoreconf -fvi
}

build() {
	cd "${pkgname}/"
	./configure --prefix=/usr
	make
}

package() {
	DESTDIR="${pkgdir}" make -C "${pkgname}" install
}
