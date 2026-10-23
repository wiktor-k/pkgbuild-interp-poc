# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Daniel Peukert <daniel@peukert.cc>

pkgname=age-plugin-yubikey
pkgver=0.5.1
pkgrel=1
pkgdesc='Yubikey plugin for age'
arch=(x86_64 i686 arm armv6h armv7h aarch64)
url="https://github.com/str4d/$pkgname"
license=(Apache-2.0 MIT)
depends=(glibc
         libgcc
         pcsclite)
makedepends=(cargo)
optdepends=('age: for use with age'
            'rage-encryption: for use with rage')
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('51e4680ad7ad7f56535e4f3018531bd0196815659378709979617d4f17102700')

prepare() {
	cd "$_archive"
	cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
	cd "$_archive"
	cargo build --frozen --release --all-features
}

check() {
	cd "$_archive"
	cargo test --frozen --all-features
}

package() {
	cd "$_archive"
	install -Dm0755 -t "$pkgdir/usr/bin/" "target/release/$pkgname"
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE-MIT
}
