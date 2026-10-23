# Maintainer: Caleb Maclennan <caleb@alerque.com>

pkgname=hyprpwcenter
pkgver=0.1.2
pkgrel=9
pkgdesc='A GUI Pipewire control center'
arch=(x86_64)
url="https://github.com/hyprwm/$pkgname"
license=(BSD-3-Clause)
depends=(
         libgcc
         libstdc++
         glibc # libc.so libm.so
         hyprtoolkit libhyprtoolkit.so
         hyprutils libhyprutils.so
         libdrm # libdrm.so
         libpipewire libpipewire-0.3.so
         pixman
)
makedepends=(cmake)
optdepends=('ttf-material-symbols-variable: Recommended font for icons')
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('ab4cfd4710566b07e98973d6723c24802d95774aa9c02aca839ff03e3bf09659')

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
