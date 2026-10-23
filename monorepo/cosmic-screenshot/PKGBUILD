# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@proton.me>

pkgname=cosmic-screenshot
pkgver=1.10.0
pkgrel=1
epoch=1
pkgdesc='Utility for capturing screenshots via XDG Desktop Portal'
arch=(x86_64)
url=https://github.com/pop-os/cosmic-screenshot
license=(GPL-3.0-only)
groups=(cosmic)
depends=(
  cosmic-icon-theme
  libgcc
  glibc
  xdg-desktop-portal-cosmic
)
makedepends=(
  cargo
  git
  just
  lld
)
source=(
  git+https://github.com/pop-os/cosmic-screenshot.git#tag=epoch-${pkgver}
)
b2sums=('ba21ddd26c6a7c8184e6960c237cd40e33f31fe909c7497d7eb98b5ad7185a8ec6ee7f7bb1a0fdfce510afe2e3bab6970c7f49eb15a50168669ed24bf6b0fd43')

prepare() {
  cd cosmic-screenshot
  cargo fetch --locked
}

build() {
  cd cosmic-screenshot
  RUSTFLAGS+=" -C link-arg=-fuse-ld=lld"
  just build-release --frozen
}

package() {
  cd cosmic-screenshot
  just rootdir="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
