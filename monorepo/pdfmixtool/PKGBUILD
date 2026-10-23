# Maintainer: George Hu <integral@archlinux.org>
# Contributor: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Marco Scarpetta <marcoscarpetta02@gmail.com>

pkgname=pdfmixtool
pkgver=1.2.2
pkgrel=2
pkgdesc="Application to split, merge, rotate and mix PDF files"
arch=('x86_64')
url="https://scarpetta.eu/${pkgname}"
license=('GPL-3.0-or-later')
depends=('hicolor-icon-theme' 'qpdf' 'imagemagick' 'qt6-svg' 'poppler-qt6')
makedepends=('cmake' 'qt6-tools')
source=("https://gitlab.com/scarpetta/${pkgname}/-/archive/v${pkgver}/${pkgname}-v${pkgver}.tar.gz")
sha256sums=('5feccbc7d797920bede35a39526cb0093052f38e2bfc17c0c609e6129aba82de')

build() {
	cmake -S "${pkgname}-v${pkgver}" \
		-B build \
		-D CMAKE_INSTALL_PREFIX=/usr \
		-D CMAKE_BUILD_TYPE=None

	cmake --build build
}

package() {
	DESTDIR="${pkgdir}" cmake --install build
}
