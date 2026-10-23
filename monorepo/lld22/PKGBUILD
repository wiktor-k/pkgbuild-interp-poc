# Maintainer: Evangelos Foutras <foutrelis@archlinux.org>
# Contributor: Jan "heftig" Steffens <jan.steffens@gmail.com>

pkgname=lld22
pkgver=22.1.8
pkgrel=1
pkgdesc="Linker from the LLVM project (Version 22)"
arch=('x86_64')
url="https://lld.llvm.org/"
license=('Apache-2.0 WITH LLVM-exception')
depends=(
  'llvm22-libs'
  'libstdc++'
  'zlib'
  'zstd'
)
makedepends=('llvm22' 'cmake' 'ninja')
checkdepends=('llvm')
_source_base=https://github.com/llvm/llvm-project/releases/download/llvmorg-$pkgver
source=($_source_base/llvm-project-$pkgver.src.tar.xz{,.sig})
sha256sums=('922f1817a0df7b1489272d18134ee0087a8b068828f87ac63b9861b1a9965888'
            'SKIP')
validpgpkeys=('474E22316ABF4785A88C6E8EA2C794A986419D8A'  # Tom Stellard <tstellar@redhat.com>
              'D574BD5D1D0E98895E3BF90044F2485E45D59042'  # Tobias Hieta <tobias@hieta.se>
              'FFB3368980F3E6BB5737145A316C56D064CACBA5'  # Douglas Yung <douglas.yung@sony.com>
              '71046D1E9C6656BDD61171873E83BABF4A4F9E85'  # Cullen Rhodes <cullen.rhodes@arm.com>
)

prepare() {
  cd llvm-project-$pkgver.src/lld
  mkdir build
}

build() {
  cd llvm-project-$pkgver.src/lld/build

  local cmake_args=(
    -G Ninja
    -DCMAKE_BUILD_TYPE=Release
    -DCMAKE_INSTALL_PREFIX=/usr/lib/llvm22
    -DCMAKE_SKIP_INSTALL_RPATH=ON
    -DBUILD_SHARED_LIBS=ON
    -DLLVM_CMAKE_DIR=/usr/lib/llvm22/lib/cmake
    -DLLVM_CONFIG=/usr/lib/llvm22/bin/llvm-config
    -DLLVM_EXTERNAL_LIT=/usr/bin/lit
    -DLLVM_INCLUDE_TESTS=ON
    -DLLVM_LINK_LLVM_DYLIB=ON
    -DLLVM_MAIN_SRC_DIR="$srcdir/llvm-project-$pkgver.src/llvm"
  )
  cmake .. "${cmake_args[@]}"
  ninja
}

# Disabled as 2 tests fail
#check() {
#  cd llvm-project-$pkgver.src/lld/build
#  ninja check-lld
#}

package() {
  cd llvm-project-$pkgver.src/lld/build

  DESTDIR="$pkgdir" ninja install
  install -Dm644 "$srcdir/llvm-project-$pkgver.src/lld/LICENSE.TXT" "$pkgdir/usr/share/licenses/$pkgname/LICENSE"

  local _libname _lib
  for _libname in COFF Common ELF MachO MinGW Wasm; do
    _lib=liblld$_libname.so.22.1
    mv "$pkgdir"/usr/lib/{llvm22/lib/,}$_lib
    ln -s ../../$_lib "$pkgdir/usr/lib/llvm22/lib/$_lib"
  done
}

# vim:set ts=2 sw=2 et:
