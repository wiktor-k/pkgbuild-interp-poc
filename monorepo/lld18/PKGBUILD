# Maintainer: Evangelos Foutras <foutrelis@archlinux.org>
# Contributor: Jan "heftig" Steffens <jan.steffens@gmail.com>

pkgname=lld18
pkgver=18.1.6
pkgrel=2
pkgdesc="Linker from the LLVM project (Version 18)"
arch=('x86_64')
url="https://lld.llvm.org/"
license=('Apache-2.0 WITH LLVM-exception')
depends=('llvm18-libs' 'gcc-libs' 'zlib' 'zstd')
makedepends=('llvm18' 'cmake' 'ninja' 'python-sphinx')
checkdepends=('llvm')
_source_base=https://github.com/llvm/llvm-project/releases/download/llvmorg-$pkgver
source=($_source_base/lld-$pkgver.src.tar.xz{,.sig}
        $_source_base/llvm-$pkgver.src.tar.xz{,.sig}
        $_source_base/libunwind-$pkgver.src.tar.xz{,.sig}
        $_source_base/cmake-$pkgver.src.tar.xz{,.sig})
sha256sums=('f1f059c2bf98ffa558cd0c48ea568736c41f0c8029dabb53147d48b9efdaa802'
            'SKIP'
            'c231d0a5445db2aafab855e052c247bdd9856ff9d7d9bffdd04e9f0bf8d5366f'
            'SKIP'
            '7ed4651612a5c07aadfda38fdbd4e7e786877826790b5d74ce30edd638eb00b3'
            'SKIP'
            'a643261ed98ff76ab10f1a7039291fa841c292435ba1cfe11e235c2231b95cdb'
            'SKIP')
validpgpkeys=('474E22316ABF4785A88C6E8EA2C794A986419D8A'  # Tom Stellard <tstellar@redhat.com>
              'D574BD5D1D0E98895E3BF90044F2485E45D59042') # Tobias Hieta <tobias@hieta.se>

prepare() {
  # https://bugs.llvm.org/show_bug.cgi?id=49228
  mv libunwind{-$pkgver.src,}

  mv cmake{-$pkgver.src,}
  cd lld-$pkgver.src
  mkdir build
}

build() {
  cd lld-$pkgver.src/build

  local cmake_args=(
    -G Ninja
    -DCMAKE_BUILD_TYPE=Release
    -DCMAKE_INSTALL_PREFIX=/usr/lib/llvm18
    -DCMAKE_SKIP_RPATH=ON
    -DBUILD_SHARED_LIBS=ON
    -DLLVM_CMAKE_DIR=/usr/lib/llvm18/lib/cmake
    -DLLVM_CONFIG=/usr/lib/llvm18/bin/llvm-config
    -DLLVM_EXTERNAL_LIT=/usr/bin/lit
    -DLLVM_INCLUDE_TESTS=ON
    -DLLVM_LINK_LLVM_DYLIB=ON
    -DLLVM_MAIN_SRC_DIR="$srcdir/llvm-$pkgver.src"
  )
  cmake .. "${cmake_args[@]}"
  ninja
}

check() {
  cd lld-$pkgver.src/build
  ninja check-lld
}

package() {
  cd lld-$pkgver.src/build

  DESTDIR="$pkgdir" ninja install
  install -Dm644 ../LICENSE.TXT "$pkgdir/usr/share/licenses/$pkgname/LICENSE"

  local _libname _lib
  for _libname in COFF Common ELF MachO MinGW Wasm; do
    _lib=liblld$_libname.so.18.1
    mv "$pkgdir"/usr/lib/{llvm18/lib/,}$_lib
    ln -s ../../$_lib "$pkgdir/usr/lib/llvm18/lib/$_lib"
  done
}

# vim:set ts=2 sw=2 et:
