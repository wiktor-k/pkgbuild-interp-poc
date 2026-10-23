# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Maintainer: Peter Jung <ptr1337@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@proton.me>

pkgname=cosmic-randr
pkgver=1.10.0
pkgrel=1
epoch=1
pkgdesc='Library and utility for displaying and configuring Wayland outputs'
arch=(x86_64)
url=https://github.com/pop-os/cosmic-randr
license=(MPL-2.0)
groups=(cosmic)
depends=(
  libgcc
  glibc
  wayland
)
makedepends=(
  cargo
  git
  just
  lld
)
source=(
  git+https://github.com/pop-os/cosmic-randr.git#tag=epoch-${pkgver}
  cosmic-randr-lto.patch
)
b2sums=('ee2b37da2354c048ee034371d469a08abd308e143c179a386bd996f98b36964177cbcd9514b9e57bc85b9700cf39db832c7a0148b119dabb40b3c596e57fa578'
        '3bcdff90827243e43c4c581a8982a2739c339a98de089739c0ceaea2fdd76e555bde0a81a52f3603ee3cf7ccf8cf9e1410fd7211e445252fda7ce909ce13d776')

prepare() {
  cd cosmic-randr
  patch -Np1 -i ../cosmic-randr-lto.patch
  cargo fetch --locked
}

build() {
  cd cosmic-randr
  RUSTFLAGS+=" -C link-arg=-fuse-ld=lld"
  just build-release --frozen
}

package() {
  cd cosmic-randr
  just rootdir="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
