# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Grzegorz Koperwas <admin@grzegorzkoperwas.site>

pkgname=awww
pkgver=0.12.1
pkgrel=1
pkgdesc="An Answer to your Wayland Wallpaper Woes"
arch=(x86_64)
url="https://codeberg.org/LGFae/awww"
license=(GPL-3.0-only)
depends=(
  dav1d
  glibc
  libgcc
  lz4
)
makedepends=(
  cargo
  scdoc
  wayland
  wayland-protocols
)
provides=(swww)
replaces=(swww)
options=(!lto)
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")
b2sums=('496a0893a6f6d73d0f01c7d8ca7a44f0fd659f75d752f755b089230a020c925880d9bc008a0c198f2d0deb2edba9c2663c3226348540d885def925c17ac675b5')

prepare() {
  cd $pkgname
  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  cd $pkgname
  cargo build --frozen --release --all-features

  # manpages
  ./doc/gen.sh
}

check() {
  cd $pkgname
  cargo test --frozen --all-features
}

package() {
  cd $pkgname
  install -vDm755 -t "$pkgdir/usr/bin" \
    target/release/awww \
    target/release/awww-daemon

  install -vDm644 completions/awww.bash "$pkgdir/usr/share/bash-completion/completions/awww"
  install -vDm644 completions/awww.fish "$pkgdir/usr/share/fish/vendor_completions.d/awww.fish"
  install -vDm644 completions/_awww "$pkgdir/usr/share/zsh/site-functions/_awww"
  install -vDm644 completions/awww.elv "$pkgdir/usr/share/elvish/lib/awww.elv"

  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" ./*.md
  install -vDm644 -t "$pkgdir/usr/share/man/man1" ./doc/generated/*.1
}
