# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
pkgname=tlottie
pkgver=1.0.6
pkgrel=1
pkgdesc="Rust library for drawing Lottie animations"
arch=('x86_64')
url="https://github.com/dkaraush/tlottie"
license=('MIT')
depends=('glibc' 'libgcc')
makedepends=('rust')
source=("$pkgname-$pkgver.tar.gz::https://github.com/dkaraush/tlottie/archive/refs/tags/v1.0.6.tar.gz")
sha512sums=('101116a4739bd4526985ff844ee159236b0d93d3abfa54ee75ac70477bbeefa111ad152cef51fbf3fb57d41ac4602cb61a5ac0277810040fe8a59286acf163f9')

prepare() {
  cd "$pkgname-$pkgver"

  cargo fetch --locked --target host-tuple
}

build() {
  cd "$pkgname-$pkgver"

  cargo build --release --frozen --features c-api
}

# check() {
#   cd "$pkgname-$pkgver"
#
#   cargo test --frozen --features c-api
# }

package() {
  cd "$pkgname-$pkgver"

  install -Dm755 -t "$pkgdir/usr/lib" target/release/libtlottie.so
  install -Dm644 -t "$pkgdir/usr/include" include/tlottie.h
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
