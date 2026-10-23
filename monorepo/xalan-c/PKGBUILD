# Maintainer: Roman Kyrylych <Roman.Kyrylych@gmail.com>
# Contributor: William Rea <sillywilly@gmail.com>

pkgname=xalan-c
epoch=1
pkgver=1.12
pkgrel=12
pkgdesc="A XSLT processor for transforming XML documents"
arch=('x86_64')
url="https://xalan.apache.org/"
license=('Apache-2.0')
makedepends=('cmake')
depends=('xerces-c')
provides=('libxalan-c.so' 'libxalanMsg.so')
validpgpkeys=('F33D281D470AB4436756147C07B3C8BC4083E800') # Roger Leigh <rleigh@debian.org>
source=("https://downloads.apache.org/xalan/xalan-c/sources/xalan_c-$pkgver.tar.gz"{,.asc}
        '0001-xalan_libicu76.patch'
        '0002-cmake_min_version.patch'
        '0003-gcc_15_fixes.patch')
sha256sums=('ee7d4b0b08c5676f5e586c7154d94a5b32b299ac3cbb946e24c4375a25552da7'
            'SKIP'
            '4ad23c14264d6de40d7d9a93dd91ff7138d8474c5fa4d7e37de59ddf8845b3cb'
            '9c437c197b06be3021b9ea885b91a4599359b49da48af3f9c56baf66d3f02803'
            '41e547def3ead09ab83dff2083b8262d6ab46b8bbfa4168283f9a608e14139df')

prepare() {
  cd xalan_c-"${pkgver}"

  patch -Np1 < ../0001-xalan_libicu76.patch
  patch -Np1 < ../0002-cmake_min_version.patch
  patch -Np1 < ../0003-gcc_15_fixes.patch
}

build() {
  mkdir build
  cd build

  cmake \
    "../xalan_c-$pkgver" \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DCMAKE_BUILD_TYPE=Release
  make
}

package() {
  cd build

  make DESTDIR="$pkgdir" install
}
