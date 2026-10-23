# Maintainer: Justin Kromlinger <hashworks@archlinux.org>
# Contributor: Gaetan Bisson <bisson@archlinux.org>

pkgname=xournalpp
pkgver=1.3.8
pkgrel=1
pkgdesc='Handwriting notetaking software with PDF annotation support'
url='https://github.com/xournalpp/xournalpp'
license=('GPL-2.0-or-later')
arch=('x86_64')
# https://github.com/xournalpp/xournalpp/blob/v1.3.0/readme/LinuxBuild.md#for-arch-linux
# Upstream still explicitly requests lua 5.3, but this is outdated information
# See https://github.com/xournalpp/xournalpp/discussions/6875
depends=('gtk3' 'poppler-glib' 'libxml2' 'portaudio' 'libsndfile' 'lua'
         'libzip' 'gtksourceview4' 'qpdf')
optdepends=('texlive-latexextra: LaTeX package')
makedepends=('git' 'cmake' 'ninja' 'help2man' 'gtest')
source=("git+https://github.com/xournalpp/xournalpp.git#tag=v${pkgver}")
sha256sums=('08ca88eede3f10864163f43b0d50134af8923ff4dd3cd923fb4f804c0a21c8ae')

# https://github.com/xournalpp/xournalpp/blob/v1.3.0/readme/Compile.md

prepare() {
	mkdir -p "${srcdir}/${pkgname}/build"
}

build() {
	cd "${srcdir}/${pkgname}/build"

	# There seems to be some incompatibility with our gtest,
	# building the tests fails if DOWNLOAD_GTEST is off
	cmake \
		-G 'Ninja' \
		-DCMAKE_BUILD_TYPE='None' \
		-DCMAKE_INSTALL_PREFIX='/usr/' \
		-Wno-dev \
		-DDOWNLOAD_GTEST=on \
		-DENABLE_GTEST=ON ..

	cmake --build .
	cmake --build . --target test-units
}

check() {
	cd "${srcdir}/${pkgname}/build"
	# We call `test-units` directly instead of `cmake --build . --target test` because
	# this would also call test-gtk-integration, which we skip due to the X11 dep
	./test/test-units
}

package() {
	cd "${srcdir}/${pkgname}/build"
	DESTDIR="${pkgdir}/" cmake --build . --target install
}
