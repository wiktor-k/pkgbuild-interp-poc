# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Contributor: Philipp A. <flying-sheep@web.de>
# Contributor: Caleb Maclennan <caleb@alerque.com>

pkgname=resvg
pkgver=0.48.1
pkgrel=1
pkgdesc='SVG rendering library and CLI'
arch=('x86_64')
url="https://github.com/linebender/resvg"
license=('Apache-2.0' 'MIT')
depends=('gdk-pixbuf2')
optdepends=(
	'qt5-base: For the Qt backend'
	'cairo: For the cairo backend'
)
makedepends=(cargo cargo-c clang qt5-base qt5-tools cairo pango)
source=("$url/archive/v$pkgver/$pkgname-v$pkgver.tar.gz")
sha256sums=('40dafea6b4b9d01e9d28b6d49f1e912daf3e9055676ad9179a5a2db6e7386945')

prepare() {
	cd "$pkgname-$pkgver"
	cargo fetch --locked --target "$(rustc --print host-tuple)"
	mkdir -p tools/kde-dolphin-thumbnailer/build
}

build() {
	cd "$pkgname-$pkgver"

	cargo build --workspace --frozen --release --all-features

	(
		cd tools/viewsvg
		qmake PREFIX="$pkgdir/usr"
		make
	)

	cargo cbuild --frozen --release --all-features -p resvg-capi
	cargo doc --release --no-deps -p resvg-capi
}

check() {
	cd "$pkgname-$pkgver"

	cargo test --frozen --all-features
}

package() {
	cd "$pkgname-$pkgver"

	install -Dm755 -t "$pkgdir/usr/bin/" target/release/{resvg,usvg} tools/viewsvg/viewsvg
	cargo cinstall --frozen --release --all-features -p resvg-capi --destdir="$pkgdir" --prefix=/usr
	install -d "$pkgdir/usr/share/doc/resvg"
	cp -r target/doc/* "$pkgdir/usr/share/doc/resvg"
}

# vim: ts=2 sw=2 et:
