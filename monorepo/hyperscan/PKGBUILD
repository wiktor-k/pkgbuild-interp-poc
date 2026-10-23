# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Alireza Ayinmehr <alireza.darksun@gmail.com>
# Contributor: Maxwell Pray a.k.a. Synthead <synthead@gmail.com>
# Contributor: JSkier <jskier at gmail dot com>
# Contributor: Ivan Shapovalov <intelfx@intelfx.name>

pkgname=hyperscan
pkgver=5.4.2
pkgrel=3
pkgdesc='High-performance regular expression matching library'
arch=(x86_64)
url=https://www.hyperscan.io
license=(BSD-3-Clause)
depends=(
  glibc
  libgcc
  libstdc++
)
makedepends=(
  boost
  cmake
  git
  ninja
  python
  ragel
)
options=(!lto)
_tag=bc3b191ab56055e8560c7cdc161c289c4d76e3d2
source=(git+https://github.com/intel/hyperscan.git#tag=${_tag})
b2sums=('0a7de0065af033f1d40d4f31b16afa307f74bb5b994659182551ab1035356523e46aa7f0bc3521fc05895b0d09996fefecb68fb50a9c6c9b892a96d0215469fa')

pkgver() {
  cd hyperscan
  git describe --tags | sed 's/^v//'
}

build() {
  export CMAKE_POLICY_VERSION_MINIMUM=3.5
  cmake -S hyperscan -B build -G Ninja \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DCMAKE_INSTALL_LIBDIR=lib \
    -DBUILD_SHARED_LIBS=ON \
    -Wno-dev
  cmake --build build
}

check() {
  build/bin/unit-hyperscan
}

package() {
  DESTDIR="${pkgdir}" cmake --install build
  install -Dm 644 hyperscan/COPYING -t "${pkgdir}"/usr/share/licenses/hyperscan/
}

# vim: ts=2 sw=2 et:
