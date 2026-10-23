# Maintainer: Peter Jung <ptr1337@archlinux.org>

pkgname=cosmic-viewer
pkgver=1.10.0
pkgrel=1
pkgdesc='Image viewer for the COSMIC desktop environment'
arch=(x86_64)
url=https://github.com/pop-os/cosmic-viewer
license=(GPL-3.0-only)
groups=(cosmic)
depends=(
  dav1d
  libgcc
  glibc
  hicolor-icon-theme
  libde265
  libjpeg-turbo
  libstdc++
  libxkbcommon
)
makedepends=(
  cargo
  clang
  cmake
  git
  just
  lld
)
source=(git+https://github.com/pop-os/cosmic-viewer.git#tag=epoch-${pkgver})
b2sums=('c0bb78f21520b1656645ac91be81a19d33639830b1a39cee5888fd2176c044b3d398a4f5afa46fd6ff001a1194b7a1974d4ab03d86000699607efc7ae80809f9')

prepare() {
  cd "$pkgname"
  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  cd "$pkgname"

  RUSTFLAGS+=" -C link-arg=-fuse-ld=lld"
  # embedded libheif is a static C++ lib; lld can't read GCC LTO bitcode
  CFLAGS+=" -ffat-lto-objects"
  CXXFLAGS+=" -ffat-lto-objects"
  # link against system libjpeg-turbo instead of building the bundled copy
  export TURBOJPEG_SOURCE=pkg-config
  export TURBOJPEG_DYNAMIC=1

  # use nice to build with lower priority
  nice just build-release --frozen
}

package() {
  cd "$pkgname"
  just rootdir="$pkgdir" install
}

# vim: ts=2 sw=2 et:
