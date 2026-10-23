# Maintainer: Jonas Witschel <diabonas@archlinux.org>

pkgname=mpv-mpris
pkgver=1.3
pkgrel=1
pkgdesc='MPRIS plugin for mpv'
arch=('x86_64')
url='https://github.com/hoyon/mpv-mpris'
license=('MIT')
depends=('glib2' 'libavformat.so' 'mpv')
checkdepends=('jq' 'playerctl' 'socat' 'sound-theme-freedesktop' 'xorg-server-xvfb')
source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz")
sha512sums=('73397284d0180826f324398be4362920f8eb47cc38c271013b8e754163cb34bfc18b0a8a642b3b5b1fdf29eb3a689cc340ccb70bda4450db29c992e4526aff4a')

build() {
	cd "$pkgname-$pkgver"
	make
}

check() {
	cd "$pkgname-$pkgver"
	make test
}

package() {
	cd "$pkgname-$pkgver"
	make PREFIX=/usr DESTDIR="$pkgdir" install
	install -Dm644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
}
