# Maintainer: Brett Cornwall <ainola@archlinux.org>
# Contributor: Lennard Hofmann

pkgname=wf-recorder
pkgver=0.6.0
pkgrel=2
pkgdesc="Screen recorder for wlroots-based compositors such as sway"
arch=("x86_64")
url="https://github.com/ammen99/wf-recorder"
license=('MIT')
depends=(
	'libavcodec.so'
	'libavdevice.so'
	'libavfilter.so'
	'libavformat.so'
	'libavutil.so'
	'libpulse.so'
	'libpulse-simple.so'
	'libswresample.so'
	'libwayland-client.so'
)
makedepends=(
	'meson'
	'wayland-protocols'
)
optdepends=(
	'slurp: Select a region to record'
)
source=(
	"$pkgname-$pkgver.tar.gz::https://github.com/ammen99/wf-recorder/archive/v$pkgver.tar.gz"
	ffmpeg-9.patch
)
sha256sums=('52d2c952506d63708f9a8f1aacd4d6ca176287caf3507c8ff2882fa0390cb391'
            'd151f6163ee2839df0edf5bcafd82f41b88832cae66ddb6673ce54d7b8f680c6')

prepare() {
	patch -d $pkgname-$pkgver -p1 < ffmpeg-9.patch
}

build() {
	meson "$pkgname-$pkgver" build \
		-Dpulse=enabled \
		--prefix=/usr \
		--buildtype=plain
	ninja -C build
}

package() {
	DESTDIR="$pkgdir/" ninja -C build install
	cd "$pkgname-$pkgver"
	install -Dm644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
	install -Dm644 README.md -t "$pkgdir/usr/share/doc/$pkgname"
}
