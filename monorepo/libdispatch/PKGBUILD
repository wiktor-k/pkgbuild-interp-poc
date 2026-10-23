# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Adrian Perez de Castro <aperez@igalia.com>
pkgname=libdispatch
pkgver=6.2.2
pkgrel=1
pkgdesc='Comprehensive support for concurrent code execution on multicore hardware'
arch=('x86_64')
url=https://apple.github.io/swift-corelibs-libdispatch
license=('Apache-2.0')
depends=('glibc')
makedepends=('git' 'clang' 'cmake' 'ninja')
provides=('libblocksruntime')
source=("${pkgname}::git+https://github.com/apple/swift-corelibs-libdispatch.git#tag=swift-${pkgver%.0}-RELEASE"
        'libdispatch-remove-werror.patch')
sha512sums=('595ad53e4fe97d1084881fe7917f92b96d2d9cfc1ad86d490f2e0ab27d0d9a9004d336bf50823f3186f0ac26eb81280f802fff8d4e424bc9b8da511f8299d4b5'
            '41f32b71b5f87c1395f11b716e2ffd706299c14302453850758b0d56a94bdc60d642d79d3f486cfebd71162644ab2f7e43f8bcc2fbed4262c9c78f3c1d46be73')

prepare () {
  cd "${pkgname}"
  patch -Np1 -i "${srcdir}/libdispatch-remove-werror.patch"
}

build () {
  cd "${pkgname}"
  export CC=clang
  export CXX=clang
  cmake \
    -Bbuild \
    -GNinja \
    -DCMAKE_BUILD_TYPE=RelWithDebInfo \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DBlocksRuntime_INCLUDE_DIR=/usr/include \
    -DBlocksRuntime_LIBRARIES=/usr/lib/libBlocksRuntime.so
  ninja -C build
}

check () {
  cd "${pkgname}"
  ninja -C build test
}

package () {
  cd "${pkgname}"
  DESTDIR="${pkgdir}" ninja -C build install
}
