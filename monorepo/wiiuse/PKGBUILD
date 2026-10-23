# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=wiiuse
pkgver=0.16.0
pkgrel=1
pkgdesc="A library written in C that connects with several Nintendo Wii remotes"
arch=(x86_64)
url="https://github.com/wiiuse/wiiuse"
license=(GPL-3.0-or-later)
depends=(
  bluez-libs
  glibc
)
provides=(libwiiuse.so)
makedepends=(cmake)
source=($url/archive/$pkgver/$pkgname-$pkgver.tar.gz)
sha512sums=('ca40c11c4e6a466b6d9fdf0446c22f5c8ff87c693c0a6dd8e5c3a35d18e597fc7d1758e4e3c06f6583f5d6c6881e10ae91ed4286ea5a5f8f1bbad1a3b529576a')
b2sums=('b3e75fdcd6302790757869dba44c063f68c4caa4fe5fa44c0434ddb0e3223d9578262925425612ca530c952bf1798c17bf89e8d1c84e18e3da71c384470a5e15')

build() {
  local cmake_options=(
    -B build
    -D BUILD_SHARED_LIBS=ON
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_LIBDIR=lib
    -D CMAKE_INSTALL_PREFIX=/usr
    -S $pkgname-$pkgver
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
  install -vDm 644 $pkgname-$pkgver/{{CHANGELOG,README}.mkd,RELEASE.md} -t "$pkgdir/usr/share/doc/$pkgname/"
}
