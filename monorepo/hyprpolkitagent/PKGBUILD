# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Adam Perkowski <adas1per@protonmail.com>

pkgname=hyprpolkitagent
pkgver=0.2.0
pkgrel=1
pkgdesc='Simple polkit authentication agent for Hyprland, written in QT/QML'
arch=(x86_64 aarch64)
url="https://github.com/hyprwm/$pkgname"
license=(BSD-3-Clause)
depends=(libgcc # libgcc_s.so
         'libstdc++' # libstdc++.so
         glibc # libc.so libm.so
         hyprgraphics
         hyprland-qt-support
         hyprtoolkit
         hyprutils
         polkit-qt6 # libpolkit-qt6-core-1.so libpolkit-qt6-agent-1.so
         qt6-base # libQt6Widgets.so libQt6Gui.so libQt6Qml.so libQt6Core.so
         qt6-declarative  # libQt6QuickControls2.so
         sdbus-cpp)
makedepends=(cmake)
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('dcb1147e28efe03f308513f9dad69ab54e3cb9eca38795c4498d968b4f5ec602')

build() {
	cd "$_archive"
	cmake -B build \
		-D CMAKE_INSTALL_PREFIX=/usr \
		-D CMAKE_INSTALL_LIBEXECDIR=lib/$pkgname \
		-D CMAKE_BUILD_TYPE=Release
	cmake --build build
}

package() {
	cd "$_archive"
	DESTDIR="$pkgdir" cmake --install build
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}


