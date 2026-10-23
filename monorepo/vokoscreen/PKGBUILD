# Maintainer: George Hu <integral@archlinux.org>
# Contributor: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Elrondo46 TuxnVape <elrond94@hotmail.com>
# Contributor: Ivan Fonseca <ivanfon@riseup.net>
# Contributor: Alfredo Ramos <alfredo dot ramos at yandex dot com>
# Contributor: Giacomo <giacomogiorgianni at gmail dot com>

pkgname=vokoscreen
_pkgname=vokoscreenNG
pkgver=4.11.0
pkgrel=1
pkgdesc="Easy to use screencast creator"
arch=('x86_64')
url="https://linuxecke.volkoh.de/${pkgname}/${pkgname}.html"
license=('GPL-2.0-or-later')
depends=(
	'glib2'
	'glibc'
	'gst-plugins-bad'
	'gst-plugins-base'
	'gst-plugins-base-libs'
	'gst-plugins-good'
	'gstreamer'
	'hicolor-icon-theme'
	'libgcc'
	'libpulse'
	'libstdc++'
	'libx11'
	'qt6-base'
	'qt6-multimedia'
	'qt6-webengine'
	'qt6-websockets'
	'wayland'
)
makedepends=('qt6-tools')
optdepends=('gst-plugin-pipewire: Wayland support')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/vkohaupt/${_pkgname}/archive/${pkgver}.tar.gz"
	"vokoscreenNG.appdata.xml")
sha256sums=('5313f029af3d7e8dfa08344d30ef619d405fa4e6f6db2ba6e021fc0db07b1842'
            '7ca66de57de1bb0eb15794656ae340864f0ce7d377cf2902466cc8a976aeddfe')

build() {
	cd "${_pkgname}-${pkgver}/"
	qmake6 PREFIX=/usr src/vokoscreenNG.pro
	make
}

package() {
	install -Dm644 "${_pkgname}.appdata.xml" -t "${pkgdir}/usr/share/metainfo/"

	cd "${_pkgname}-${pkgver}/"
	install -Dm755 "${_pkgname}" -t "${pkgdir}/usr/bin/"

	cd src/applications
	install -Dm644 "${_pkgname}.png" -t "${pkgdir}/usr/share/icons/hicolor/256x256/apps/"
	install -Dm644 "${_pkgname}.desktop" -t "${pkgdir}/usr/share/applications/"
}
