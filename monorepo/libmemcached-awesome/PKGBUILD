# Maintainer: Evangelos Foutras <foutrelis@archlinux.org>

pkgname=libmemcached-awesome
_name=${pkgname/-awesome}
pkgver=1.1.4
pkgrel=4
pkgdesc="C/C++ client library and tools for the memcached server"
arch=('x86_64')
url="https://awesomized.github.io/libmemcached/"
license=(BSD-3-Clause)
depends=(
  glibc
  libevent
  libgcc
  libsasl
  libstdc++
  openssl
)
makedepends=('cmake' 'memcached' 'python-sphinx' 'python-sphinx_rtd_theme')
provides=('libmemcached')
conflicts=('libmemcached')
replaces=('libmemcached')
options=('!debug') # some interaction with FIXTURES_OPTIONS_CONFIG makes it
                   # look in the wrong directory (/usr/src/..) for the file
source=(
  https://github.com/awesomized/libmemcached/archive/$pkgver/$pkgname-$pkgver.tar.gz
  $pkgname-1.1.4-memcapable-test.patch::https://github.com/awesomized/libmemcached/commit/664fada99d0260fc07c77b8514979fda2175e20c.patch
)
sha256sums=('c477e1f6510e1dc698e84f3717ce690a8f65b94c616ecaa62306cce0f5e3116a'
            '0cb231caf77e3a94c357643317fcb0ebe98582632ed19a61a8aa72c98bcdec02')

prepare() {
  # Fix issue with bin/memcapable test: https://github.com/awesomized/libmemcached/issues/161
  patch -Np1 -d $_name-$pkgver -i ../$pkgname-1.1.4-memcapable-test.patch
}

build() {
  local cmake_options=(
    -B build
    -D BUILD_DOCS=ON
    -D BUILD_DOCS_MANGZ=OFF
    -D BUILD_TESTING=ON
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
    -D ENABLE_HASH_FNV64=ON
    -D ENABLE_HASH_HSIEH=ON
    -D ENABLE_HASH_MURMUR=ON
    -D ENABLE_MEMASLAP=ON
    -D ENABLE_OPENSSL_CRYPTO=ON
    -D ENABLE_SASL=ON
    -S libmemcached-$pkgver
  )

  CFLAGS+=' -ffat-lto-objects'
  CXXFLAGS+=' -ffat-lto-objects'

  cmake "${cmake_options[@]}"
  cmake --build build --verbose
}

check() {
  local ctest_flags=(
    --test-dir build
    --output-on-failure
    --parallel "$(nproc)"
  )

  ctest "${ctest_flags[@]}"
}

package() {
  DESTDIR="$pkgdir" cmake --install build
  install -Dm644 libmemcached-$pkgver/LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
}

# vim:set ts=2 sw=2 et:
