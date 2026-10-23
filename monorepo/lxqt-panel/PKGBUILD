# Maintainer: Jerome Leclanche <jerome@leclan.ch>
# Maintainer: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=lxqt-panel
pkgver=2.4.1
pkgrel=1
pkgdesc="The LXQt desktop panel"
arch=("x86_64")
groups=("lxqt")
url="https://github.com/lxqt/$pkgname"
# https://github.com/lxqt/lxqt-panel/blob/2.3.3/AUTHORS#L9
license=("LGPL-2.1-or-later")
depends=(
	"libdbusmenu-lxqt" "lxqt-menu-data" "lxqt-themes" "libxtst"
	"lxqt-globalkeys" "solid" "libxcb" "libQt6Xdg.so" "layer-shell-qt"
)
optdepends=(
	"libpulse: Volume control plugin"
	"alsa-lib: Volume control plugin"
	"libstatgrab: CPU monitor and Network monitor plugins"
	"libsysstat: System Statistics plugin"
	"lm_sensors: Sensors plugin"
)
makedepends=(
	"lxqt-build-tools" "liblxqt" "libpulse" "libstatgrab" "libsysstat"
	"lm_sensors" "libxdamage" "alsa-lib"
)
source=(
	"https://github.com/lxqt/$pkgname/releases/download/$pkgver/$pkgname-$pkgver.tar.xz"{,.asc}
)
sha256sums=('7828d7e8fabf66ef1491e50b609b8eb041d74c433554f7d073ac07af26843d02'
            'SKIP')
validpgpkeys=(
	"19DFDF3A579BD509DBB572D8BE793007AD22DF7E"  # Pedram Pourang <tsujan2000@gmail.com>
	"48836EB124A1EAAE695EDFEA229CA0A00A17D258"  # the key owned by an email used by https://github.com/stefonarch/, a long-time LXQt contributor, as per keys.openpgp.org
)

build() {
	cmake -B build -S "$srcdir/$pkgname-$pkgver" \
		-DCMAKE_INSTALL_PREFIX=/usr \
		-DCMAKE_BUILD_TYPE=None
	make -C build
}

package() {
	cd build
	make DESTDIR="$pkgdir" install
}
