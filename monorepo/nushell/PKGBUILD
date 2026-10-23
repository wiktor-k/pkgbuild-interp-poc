# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: KokaKiwi <kokakiwi+aur@kokakiwi.net>
# Contributor: Felix Golatofski <contact@xdfr.de>
# Contributor: Bumsik Kim <k.bumsik@gmail.com>

pkgname=nushell
pkgver=0.116.1
pkgrel=1
pkgdesc='A new type of shell'
arch=(x86_64)
url='https://www.nushell.sh'
license=(MIT)
depends=(
  curl
  glibc
  libgcc
  libgit2
  libssh2
  openssl
  sqlite3
  zstd
)
makedepends=(cargo git)
checkdepends=(procps-ng)
install=nushell.install
source=("git+https://github.com/nushell/nushell.git#tag=$pkgver")
sha512sums=('ca68230e1d5f94b9a9657b02a78bb28a5ce7ab01bce985bd4693a0eea6326ba12b2dd1607016ab8c19d2ae6f5721aadc9571d2cfc6feff2a65c87caaee99af84')
b2sums=('9a9aa5c38fab0eb7748c3f23db7a085ef65b42ffb376bd0c1d3585fa7c7b2d466dc592c7fb7492d65cedf6e6d2f3431c2464b744634f32f7c54aa63d3fd9b41f')

prepare() {
  cd "$pkgname"

  # https://github.com/nushell/nushell/issues/18578
  cargo update -v --precise 1.5.3 ethnum

  export CARGO_HOME="$srcdir"
  cargo fetch --locked --target host-tuple
}

build() {
  cd "$pkgname"

  export CARGO_PROFILE_RELEASE_DEBUG=2
  export CARGO_PROFILE_RELEASE_STRIP=false
  export CARGO_HOME="$srcdir"
  export CARGO_PROFILE_RELEASE_LTO=thin
  export CARGO_PROFILE_RELEASE_CODEGEN_UNITS=1
  export CARGO_PROFILE_RELEASE_OPT_LEVEL=3

  CFLAGS+=' -ffat-lto-objects'

  export LIBGIT2_NO_VENDOR=1
  export LIBSSH2_SYS_USE_PKG_CONFIG=1
  export LIBSQLITE3_SYS_USE_PKG_CONFIG=1
  export ZSTD_SYS_USE_PKG_CONFIG=1

  cargo build --release --frozen --workspace
}

#check() {
#  cd "$pkgname"
#  # Skip tests that pretend they know what the local terminal is going to be
#  local skipped=(
#    plugins::stream::echo_interactivity_on_slow_pipelines
#    plugins::stress_internals::test_exit_before_hello_stdio
#  )
#  cargo test --frozen --workspace -- ${skipped[@]/#/--skip }
#}

package() {
  cd "$pkgname"

  find target/release \
    -maxdepth 1 \
    -executable \
    -type f \
    -name "nu*" \
    -exec install -vDm755 -t "$pkgdir/usr/bin" "{}" +

  # documentation
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" README.md

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
