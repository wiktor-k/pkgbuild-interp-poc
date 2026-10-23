# Maintainer: Laurent Carlier <lordheavym@gmail.com>
# Contributor: Luca Bennati <lucak3 AT gmail DOT com>
# Contributor: Glaucous <glakke1 at gmail dot com>

pkgname=apitrace
pkgver=14.0
pkgrel=1
pkgdesc="Graphics API Tracing"
arch=('x86_64')
url="https://github.com/apitrace/apitrace"
license=('MIT')
depends=(
  'brotli'
  'glibc'
  'python'
  'python-numpy'
  'python-pillow'
  'libpng'
  'libprocps'
  'libx11'
  'zlib'
)
makedepends=(
  'cmake'
  'git'
  'mesa'
  'qt5-base'
)
optdepends=('qt5-base: GUI support')
source=("$pkgname::git+https://github.com/apitrace/apitrace.git#tag=${pkgver}"
        "git+https://github.com/apitrace/gltrim-tests.git"
        "git+https://github.com/google/brotli.git"
        "git+https://github.com/google/googletest.git"
        "git+https://github.com/ianlancetaylor/libbacktrace.git"
        "git+https://github.com/apitrace/libpng.git"
        "git+https://github.com/google/snappy.git"
        "git+https://github.com/madler/zlib.git"
        "git+https://github.com/microsoft/DirectXMath.git")
sha256sums=('ae711c7a2152b2ada5e319f96c58aabdd531370ed12acd32fc9f3153c82d3149'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP')

prepare() {
  cd $pkgname

  git submodule init
  git config submodule.frametrim/tests.url "$srcdir/gltrim-tests"
  git config submodule.thirdparty/brotli.url "$srcdir/brotli"
  git config submodule.thirdparty/gtest.url "$srcdir/googletest"
  git config submodule.thirdparty/libbacktrace.url "$srcdir/libbacktrace"
  git config submodule.thirdparty/libpng.url "$srcdir/libpng"
  git config submodule.thirdparty/snappy.url "$srcdir/snappy"
  git config submodule.thirdparty/zlib.url "$srcdir/zlib"
  git config submodule.thirdparty/directxmath.url "$srcdir/DirectXMath"
  git -c protocol.file.allow=always submodule update
}

build() {
  local cmake_options=(
    -B build
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
    -D ENABLE_GUI=TRUE
    -S $pkgname
    -W no-dev
  )

  cmake "${cmake_options[@]}"
  cmake --build build --verbose
}

check() {
  ctest --test-dir build --output-on-failure
}

package() {
  DESTDIR="$pkgdir" cmake --install build
  install -vDm 644 $pkgname/LICENSE -t "${pkgdir}/usr/share/licenses/apitrace/"
}
