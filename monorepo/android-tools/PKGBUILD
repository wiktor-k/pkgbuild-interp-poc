# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Contributor: Anatol Pomozov
# Contributor: 謝致邦 <Yeking@Red54.com>
# Contributor: Alucryd <alucryd at gmail dot com>

pkgname=android-tools
pkgver=37.0.0.p1
_tag="${pkgver/.p/p}"
pkgrel=1
pkgdesc='Android platform tools'
arch=(x86_64)
url='https://developer.android.com/tools'
license=('Apache-2.0 AND BSD-3-Clause AND GPL-2.0-or-later AND LGPL-2.1-or-later AND MIT AND libselinux-1.0')
depends=(abseil-cpp brotli fmt glibc libgcc libstdc++ libusb lz4 pcre2 protobuf zlib zstd)
makedepends=(corrosion cmake gtest ninja)
optdepends=('python: {mk,unpack_,repack_}bootimg and mkdtboimg support'
	    'android-udev: optional additional device udev rules')
source=(https://github.com/nmeum/android-tools/releases/download/${_tag}/android-tools-${_tag}.tar.xz)
sha256sums=('2c6de1abe16b211c7d9f59276217e14fd9e0e18b20e331d3c84358f5dffd576e')

build() {
	cd "android-tools-${_tag}"

	cmake \
		-DCMAKE_INSTALL_PREFIX=/usr \
		-DCMAKE_BUILD_TYPE=Release \
		-DCMAKE_CXX_FLAGS="$CXXFLAGS" \
		-DCMAKE_C_FLAGS="$CFLAGS" \
		-DANDROID_TOOLS_ADB_ENABLE_MDNS=ON \
		-G Ninja -S . -B build
	cmake --build build
}

package() {
	cd "android-tools-${_tag}"

	DESTDIR="$pkgdir" ninja -C build install

	install -Dm644 vendor/selinux/LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/libselinux.txt"
	install -Dm644 vendor/fmtlib/LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/MIT.txt"
	install -Dm644 vendor/avb/LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/BSD-3-Clause.txt"
}
