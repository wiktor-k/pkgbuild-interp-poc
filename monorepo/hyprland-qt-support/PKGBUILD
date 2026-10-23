# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Charles Dong <chardon_cs@proton.me>

pkgname=hyprland-qt-support
pkgver=0.1.0
pkgrel=14
pkgdesc='QML style provider for Hypr* QT apps'
arch=(x86_64 aarch64)
url="https://github.com/hyprwm/$pkgname"
license=(BSD-3-Clause)
depends=(
         libgcc
         libstdc++
         glibc # libc.so libm.so
         hyprlang libhyprlang.so
         qt6-base # libQt6Core.so libQt6Gui.so libQt6Widgets.so
         qt6-declarative # libQt6Qml.so libQt6QuickControls2.so
)
makedepends=(cmake)
provides=(libhyprland-quick-style-impl.so
          libhyprland-quick-style.so
          libhyprland-quick-style-implplugin.so
          libhyprland-quick-styleplugin.so)
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('cac1f980bd088b890097f3f999cfdf03e73ee94c53f3c92d0b3bc23baa9e7b2c')

build() {
	cd "$_archive"
	cmake -B build \
		-D CMAKE_INSTALL_PREFIX=/usr \
		-D INSTALL_QML_PREFIX=/lib/qt6/qml \
		-D CMAKE_BUILD_TYPE=Release
	cmake --build build
}

package() {
	cd "$_archive"
	DESTDIR="$pkgdir" cmake --install build
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
