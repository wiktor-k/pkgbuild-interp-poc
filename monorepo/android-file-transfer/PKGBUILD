# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: yaroslav <proninyaroslav@mail.ru>
# Contributor: Askhat Bakarov <sirocco{at}ngs{dot}ru>

pkgname=android-file-transfer
_pkgname="${pkgname}-linux"
pkgver=4.6
pkgrel=1
pkgdesc="Android MTP client with minimalistic UI"
url="https://whoozle.github.io/android-file-transfer-linux"
arch=('x86_64')
license=('LGPL-2.1-or-later')
depends=('qt6-base' 'fuse3' 'hicolor-icon-theme' 'file' 'android-udev' 'taglib')
makedepends=('cmake' 'qt6-tools' 'ninja')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/whoozle/android-file-transfer-linux/archive/v${pkgver}.tar.gz")
sha256sums=('d8bcbd1ea5854e2653661f7e953c955531eed540ca9aee966ae9bb64149f69e4')

build() {
	cmake -B build -S "${_pkgname}-${pkgver}" \
		-DCMAKE_BUILD_TYPE='None' \
		-DCMAKE_INSTALL_PREFIX='/usr' \
		-DCMAKE_CXX_FLAGS="${CXXFLAGS} -ffat-lto-objects" \
		-G Ninja \
		-Wno-dev
	ninja -C build
}

package() {
	DESTDIR="${pkgdir}" ninja -C build install
}
