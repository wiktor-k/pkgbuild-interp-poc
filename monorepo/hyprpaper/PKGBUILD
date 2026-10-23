# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: ThatOneCalculator <kainoa@t1c.dev>
# Contributor: Aaron Blasko <blaskoazzolaaaron [at] gmail.com>

pkgname=hyprpaper
pkgver=0.8.4
pkgrel=9
pkgdesc='a blazing fast wayland wallpaper utility with IPC controls'
arch=(x86_64)
url="https://github.com/hyprwm/$pkgname"
license=(BSD-3-Clause)
depends=(file libmagic.so
         glibc # libc.so libm.so
         libgcc libgcc_s.so
         libstdc++ libstdc++.so
         hyprlang libhyprlang.so
         hyprtoolkit libhyprtoolkit.so
         hyprutils libhyprutils.so
         hyprwire libhyprwire.so
         libglvnd libGLESv2.so
         wayland) # libwayland-cursor.so libwayland-client.so
makedepends=(cmake
             hyprwayland-scanner
             wayland-protocols
             xorgproto)
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('4912c7aede303918f38a3fc6b294429e2d541b639ebba67bcab58d1d367ea360')

build() {
	cd "$_archive"
	cmake -B build \
		-D CMAKE_INSTALL_PREFIX=/usr \
		-D CMAKE_BUILD_TYPE=Release
	make protocols
	cmake --build build
}

package() {
	cd "$_archive"
	DESTDIR="$pkgdir" cmake --install build
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
