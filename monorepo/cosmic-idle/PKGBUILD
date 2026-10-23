# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Maintainer: Peter Jung <ptr1337@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@proton.me>

pkgname=cosmic-idle
pkgver=1.10.0
pkgrel=1
epoch=1
pkgdesc='Cosmic idle daemon'
arch=(x86_64)
url=https://github.com/pop-os/cosmic-idle
license=(GPL-3.0-only)
groups=(cosmic)
depends=(
  libgcc
  glibc
  libxkbcommon
  wayland
)
makedepends=(
  cargo
  clang
  git
  just
  lld
)
source=(
  git+https://github.com/pop-os/cosmic-idle.git#tag=epoch-${pkgver}
)
b2sums=('4bdd89487aa752631de0268b53bc1b5c0784d845578c3f3ed08304da0abfeb2d214a393b256416d65cf635f32eb8647b883751daf693b45b9d749059a21242f4')

prepare() {
  cd $pkgname
  cargo fetch --locked
}

build() {
  cd $pkgname
  RUSTFLAGS+=" -C link-arg=-fuse-ld=lld"
  CC=clang just build-release --frozen
}

package() {
  cd $pkgname
  just rootdir="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
