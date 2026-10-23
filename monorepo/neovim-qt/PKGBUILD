# Maintainer: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Aleksandar Trifunović <akstrfn@gmail.com>
# Contributor: Iván Ruvalcaba <mario.i.ruvalcaba@gmail.com>
# Contributor: Ian Ker-Seymer <i.kerseymer@gmail.com>

pkgname=neovim-qt
pkgver=0.2.20
pkgrel=1
pkgdesc='GUI for Neovim'
arch=(x86_64)
url='https://github.com/equalsraf/neovim-qt'
license=(ISC)
depends=(hicolor-icon-theme neovim qt6-svg)
makedepends=(cmake git vulkan-headers)
source=("git+$url#tag=v$pkgver")
b2sums=('e9f2bfc9d3a5fcb6524ae0b127e9eb57c6dae0efad83440bcdd9b84e76b9568f302e6e6dbdbac4da3205e27804c00dfba578aa36b872d8b8760004cd48cb59b9')

prepare() {
  sed -i 's/msgpack-c 6.0.0/msgpack-c 7.0.0/g' $pkgname/CMakeLists.txt
}

build() {
  cmake \
    -B build \
    -D CMAKE_BUILD_TYPE=Release \
    -D CMAKE_C_FLAGS="$CFLAGS" \
    -D CMAKE_CXX_FLAGS="$CXXFLAGS" \
    -D CMAKE_EXE_LINKER_FLAGS="$LDFLAGS" \
    -D CMAKE_INSTALL_LIBDIR=lib \
    -D CMAKE_INSTALL_PREFIX=/usr \
    -D USE_SYSTEM_MSGPACK=on \
    -D WITH_QT=Qt6 \
    -S $pkgname
  make -C build
}

package() {
  DESTDIR="$pkgdir" make -C build install
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname" $pkgname/LICENSE
}
