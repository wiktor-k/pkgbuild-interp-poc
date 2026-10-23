# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Bartłomiej Piotrowski <bpiotrowski@archlinux.org>
# Contributor: Behem0th <grantipak@gmail.com>

pkgname=libvdpau-va-gl
pkgver=0.4.2
pkgrel=5
pkgdesc='VDPAU driver with OpenGL/VAAPI backend'
arch=(x86_64)
url=https://github.com/i-rinat/libvdpau-va-gl
license=(MIT)
depends=(
  glibc
  libgcc
  libgl
  libstdc++
  libva
  libx11
)
makedepends=(
  cmake
  git
  ninja
)
source=(git+https://github.com/i-rinat/libvdpau-va-gl.git#tag=v${pkgver})
b2sums=('b751043761da48a16c6ecaeefd9026b3bfad55335ad45c4fe17f98d8761d6d3fbc44f77388a1b0e19be6a1622550ecf49ba3e6c7dce48ce96fe82989961e0dca')

build() {
  cmake -S libvdpau-va-gl -B build -G Ninja \
    -DCMAKE_BUILD_TYPE=None \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DCMAKE_POLICY_VERSION_MINIMUM=3.5
  cmake --build build
}

package() {
  DESTDIR="${pkgdir}" cmake --install build
  install -Dm 644 libvdpau-va-gl/LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}/"
}
