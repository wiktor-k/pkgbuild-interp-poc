# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@proton.me>

pkgname=pop-launcher
pkgver=1.2.7
pkgrel=2
pkgdesc='Modular IPC-based desktop launcher service'
arch=(x86_64)
url=https://github.com/pop-os/launcher
license=(MPL-2.0)
depends=(
  dbus
  fd
  glibc
  libegl
  libgcc
  libqalculate
  libxkbcommon
  pop-icon-theme
  sh
  xdg-utils
)
makedepends=(
  cargo
  clang
  git
  just
  lld
)
source=(git+https://github.com/pop-os/launcher.git#tag=${pkgver})
b2sums=('a6598288a7f807cfa682e02c7ade9fd6bfb2c3574b408cffa1b7a764bd9a458bc47f8ec1ad4741f7654a7b9e39f679e38f50f118c1b49b17bb98967e3ac046a4')

prepare() {
  cd launcher
  cargo fetch --locked
  sed 's/lto = "fat"/lto = "thin"/g' -i Cargo.toml
}

build() {
  cd launcher
  RUSTFLAGS+=" -C link-arg=-fuse-ld=lld"
  CC=clang just build-release --frozen
}

check() {
  cd launcher
  just check
}

package() {
  cd launcher
  just rootdir="${pkgdir}" install
}
