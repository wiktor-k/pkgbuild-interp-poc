# Maintainer: Caleb Maclennan <caleb@alerque.com>

pkgname=hyprshutdown
pkgver=0.1.1
pkgrel=8
pkgdesc='A graceful shutdown utility for Hyprland '
arch=(x86_64)
url="https://github.com/hyprwm/$pkgname"
license=(BSD-3-Clause)
depends=(glibc # libc.so libm.so
         hyprtoolkit libhyprtoolkit.so
         hyprutils libhyprutils.so
         libdrm # libdrm.so
         libgcc libgcc_s.so
         libstdc++ libstdc++.so
         pixman)
makedepends=(cmake
             glaze)
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('32adb385b7bfe22398d45cd5325416bf0ac3ae4a5bd89678353dae96f92ba638')

build() {
	cd "$_archive"
	local cmake_options=(
		-D CMAKE_BUILD_TYPE=None
		-D CMAKE_INSTALL_PREFIX=/usr
	)
	cmake -B build -W no-dev ${cmake_options[@]}
	cmake --build build
}

package() {
	cd "$_archive"
	DESTDIR="$pkgdir" cmake --install build
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
}
