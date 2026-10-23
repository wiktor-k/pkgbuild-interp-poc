# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Contributor: alba4k <blaskoazzolaaaron@gmail.com>

pkgname=hyprutils
pkgver=0.14.3
pkgrel=1
pkgdesc='Hyprland utilities library used across the ecosystem'
arch=(x86_64)
url="https://github.com/hyprwm/$pkgname"
license=(BSD-3-Clause)
depends=(pixman  libpixman-1.so)
makedepends=(cmake)
# checkdepends=(gtest)
provides=(libhyprutils.so)
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('87bbfc9549ed96316b0d592c16d01f58729b23a2999a5293b73eec620b40c5f5')

build() {
	cd "$_archive"
	local cmake_options=(
		-D CMAKE_BUILD_TYPE=Release
		-D CMAKE_INSTALL_PREFIX=/usr
		-D BUILD_TESTING=False
	)
	cmake -W no-dev ${cmake_options[@]} -S . -B ./build
	cmake --build build
}

package() {
	cd "$_archive"
	DESTDIR="$pkgdir" cmake --install build
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
}
