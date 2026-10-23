# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

# WARNING: this package needs to be kept in sync with the requirements of deltachat-desktop
_upstream=deltachat-core-rust
_name=core
pkgname=deltachat-rpc-server
pkgver=2.62.0
pkgrel=1
pkgdesc="A JSON-RPC 2.0 interface to DeltaChat over standard I/O"
arch=(x86_64)
url="https://github.com/deltachat/deltachat-core-rust/tree/main/deltachat-rpc-server"
_url="https://github.com/deltachat/deltachat-core-rust"
license=(MPL-2.0)
depends=(
  glibc
  libgcc
  openssl
  sqlcipher
)
makedepends=(
  rust
)
source=("$_url/archive/v$pkgver/$_upstream-$pkgver.tar.gz")
sha512sums=('3b42fb93a27bb37ac98f2b410f36ede3df2914db5eeef7924739cf5243661247eb8a3147f77b6b4a3e2594985a065da2270c8e4b24cf9338384bfde05b872ac8')
b2sums=('f688f76b99a2ec082ea18a06bce20edccc72dcbb784c66d6150cdbcd753855d57aacdba0b69d85ca8c1bedb9e24334396212c940699be52d6286377f24c35aaa')

prepare() {
  cd $_name-$pkgver
  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  cd $_name-$pkgver
  export CFLAGS+=" -ffat-lto-objects"
  export OPENSSL_NO_VENDOR=1
  export LIBSQLITE3_SYS_USE_PKG_CONFIG=1
  cargo build --frozen --release --package $pkgname
}

check() {
  cd $_name-$pkgver
  export CFLAGS+=" -ffat-lto-objects"
  export OPENSSL_NO_VENDOR=1
  export LIBSQLITE3_SYS_USE_PKG_CONFIG=1
  cargo test --frozen --package $pkgname
}

package() {
  cd $_name-$pkgver
  install -vDm 755 target/release/$pkgname -t "$pkgdir/usr/bin/"
  install -vDm 644 $pkgname/README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
