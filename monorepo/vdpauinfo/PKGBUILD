# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Timothy Redaelli <timothy.redaelli@gmail.com>
# Contributor: Aaron Plattner <aplattner@nvidia.com>
# Contributor: Bart Verhoeven <bartverhoeven@live.com>

pkgname=vdpauinfo
pkgver=1.5
pkgrel=3
pkgdesc='Command line utility for querying the capabilities of a VDPAU device'
arch=(x86_64)
url=https://gitlab.freedesktop.org/vdpau/vdpauinfo
license=(MIT)
depends=(
  glibc
  libgcc
  libstdc++
  libvdpau
  libx11
)
makedepends=(git)
source=(git+https://gitlab.freedesktop.org/vdpau/vdpauinfo.git#tag=${pkgver})
validpgpkeys=(BD68A042C603DDAD9AA354B0F56ACC8F09BA9635) # Aaron Plattner <aplattner@nvidia.com>
b2sums=('edb8729486946f2b8ea4b2b3e582797f0ea4dfd2269ecba7f1a249576fb666e459b707f1c7f2ea425898a23b28e55719a2359794753b18e236a0f0390ee1048f')

prepare() {
  cd vdpauinfo
  autoreconf -fiv
}

build() {
  cd vdpauinfo
  ./configure \
    --prefix=/usr
  make
}

package() {
  make DESTDIR="${pkgdir}" -C vdpauinfo install
  install -Dm 644 vdpauinfo/COPYING -t "$pkgdir"/usr/share/licenses/vdpauinfo/
}
