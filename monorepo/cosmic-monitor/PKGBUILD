# Maintainer: Peter Jung <ptr1337@archlinux.org>
# Contributor: LiamillionSS <liamillion-dev at protonmail dot com>
pkgname=cosmic-monitor
pkgver=1.10.0
pkgrel=1
epoch=1
pkgdesc="COSMIC System Monitor"
arch=(x86_64)
url="https://github.com/pop-os/cosmic-monitor"
license=('GPL-3.0-only')
groups=(cosmic)
depends=(
  libxkbcommon
)
makedepends=(
  'cargo'
  'git'
  'just'
  'mold'
)
source=(git+https://github.com/pop-os/cosmic-monitor.git#tag=epoch-${pkgver})
sha256sums=('51c602a589e9ef299a7fe3180ed1efdaa4f78e96437e12fb423f4c68831c3531')

prepare() {
  cd "${pkgname}"
  export RUSTUP_TOOLCHAIN=stable
  cargo fetch --locked --target "$(rustc -vV | sed -n 's/host: //p')"
}

build() {
  cd "${pkgname}"
  # use mold instead of lld to speed up build
  RUSTFLAGS+=" -C link-arg=-fuse-ld=mold"
  # use nice to build with lower priority
  nice just build-release --frozen
}

package() {
  cd "${pkgname}"
  just rootdir="$pkgdir" install
}
