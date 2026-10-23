# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Contributor: Sergey A. <murlakatamenka@disroot.org>

pkgname=rmpc
pkgver=0.11.0
pkgrel=2
pkgdesc='A beautiful and configurable TUI client for MPD'
url='https://mierak.github.io/rmpc/'
arch=('x86_64')
license=('BSD-3-Clause')
depends=('gcc-libs')
optdepends=(
  'yt-dlp: for youtube playback'
  'ffmpeg: for youtube playback'
  'python-mutagen: for youtube playback'
  'ueberzugpp: for alternative album art method'
  'cava: for visualizer')
makedepends=('cargo' 'desktop-file-utils')
source=("$pkgname-$pkgver.tar.gz::https://github.com/mierak/rmpc/archive/v$pkgver.tar.gz")
sha256sums=('930019066228d18e9530a8c0d77f10e231ab5efbbbca73b331efcd6fbb47557d')

prepare() {
  cd "$pkgname-$pkgver"
  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  cd "$pkgname-$pkgver"
  VERGEN_GIT_DESCRIBE="Arch Linux" cargo build --release --frozen
}

check() {
  cd "$pkgname-$pkgver"
  CARGO_TARGET_DIR=test_target cargo test --frozen
}

package() {
  cd "$pkgname-$pkgver"

  install -Dm 755 "target/release/$pkgname" -t "$pkgdir/usr/bin/"

  # shell completions
  install -Dm 644 "target/completions/$pkgname.bash" \
    -t "$pkgdir/usr/share/bash-completion/completions/"
  install -Dm 644 "target/completions/$pkgname.fish" \
    -t "$pkgdir/usr/share/fish/vendor_completions.d/"
  install -Dm 644 "target/completions/_$pkgname" \
    -t "$pkgdir/usr/share/zsh/site-functions/"

  # docs
  install -Dm 644 "target/man/$pkgname.1" "$pkgdir/usr/share/man/man1/rmpc.1"
  install -Dm 644 "README.md" "$pkgdir/usr/share/doc/$pkgname/README.md"

  # license
  install -Dm 644 "LICENSE" "$pkgdir/usr/share/licenses/$pkgname/LICENCE"

  # desktop file
  desktop-file-install -m 644 --dir "$pkgdir/usr/share/applications/" "assets/$pkgname.desktop"
}

# vim: ts=2 sw=2 et:
