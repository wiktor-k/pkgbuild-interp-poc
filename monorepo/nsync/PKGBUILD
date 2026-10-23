# Maintainer: Torsten Keßler <tpkessler@archlinux.org>
# Contributor: Carl Smedstad <carl.smedstad at protonmail dot com>
# Contributor: Chih-Hsuan Yen <yan12125@gmail.com>

pkgname=nsync
pkgver=1.30.0
pkgrel=1
pkgdesc='A C library that exports various synchronization primitives, such as mutexes'
arch=('x86_64')
url='https://github.com/google/nsync'
license=('Apache-2.0')
depends=('gcc-libs')
makedepends=('cmake')
source=("$pkgname-$pkgver.tar.gz::$url/archive/$pkgver/$pkgname-$pkgver.tar.gz"
        "$pkgname-export.patch")
sha256sums=('883a0b3f8ffc1950670425df3453c127c1a3f6ed997719ca1bbe7f474235b6cc'
            'f59245d9b3845030d005cfb524efd01167e44b085b0872f1f9de856a0d2c7c48')

prepare() {
  cd $pkgname-$pkgver
  # Based on https://github.com/msys2/MINGW-packages/blob/master/mingw-w64-nsync/0001-nsync-mingw-w64.patch
  patch -Np1 -i ../$pkgname-export.patch
}

build() {
  local cmake_args=(
    -Wno-dev
    -B build
    -S $pkgname-$pkgver
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
    -D BUILD_SHARED_LIBS=ON
  )
  cmake "${cmake_args[@]}"
  cmake --build build
}

check() {
  ctest --test-dir build --output-on-failure
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
