# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=satty
pkgver=0.22.0
pkgrel=1
pkgdesc="Modern screenshot annotation tool, inspired by Swappy and Flameshot"
arch=(x86_64)
url="https://github.com/gabm/satty"
license=(MPL-2.0)
depends=(
  fontconfig
  gdk-pixbuf2
  glib2
  glibc
  gtk4
  hicolor-icon-theme
  libadwaita
  libgcc
  pango
)
makedepends=(
  cargo
  git
)
source=("$pkgname::git+$url.git#tag=v$pkgver")
b2sums=('06b136248eab0f123d7bfab34243a2ed227e6e0a4c350c172c7b942d4a42a750ce7ce5535839782b3e9fa740c7da8dafc0347c9854febb56d5a6dc0b28fe4fa3')

prepare() {
  cd $pkgname
  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  cd $pkgname
  cargo build --frozen --release --all-features
}

package() {
  cd $pkgname
  install -vDm755 -t "$pkgdir/usr/bin" target/release/satty

  install -vDm644 completions/_satty \
    "$pkgdir/usr/share/zsh/site-functions/_satty"
  install -vDm644 completions/satty.bash \
    "$pkgdir/usr/share/bash-completion/completions/satty"
  install -vDm644 completions/satty.fish \
    "$pkgdir/usr/share/fish/vendor_completions.d/satty.fish"
  install -vDm644 completions/satty.nu \
    "$pkgdir/usr/share/nushell/completions/satty.nu"
  install -vDm644 completions/satty.elv \
    "$pkgdir/usr/share/elvish/lib/satty.elv"

  install -vDm644 -t "$pkgdir/usr/share/applications" satty.desktop
  install -vDm644 -t "$pkgdir/usr/share/icons/hicolor/scalable/apps" assets/satty.svg

  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" README.md
  cp -va -t "$pkgdir/usr/share/doc/$pkgname" assets
}
