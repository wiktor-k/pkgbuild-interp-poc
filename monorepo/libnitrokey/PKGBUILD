# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: David Runge <dvzrv@archlinux.org>
# Contributor: Étienne Deparis <etienne@depar.is>

pkgname=libnitrokey
pkgver=3.8
pkgrel=4
pkgdesc='Communicate with Nitrokey stick devices in a clean and easy manner'
url='https://github.com/Nitrokey/libnitrokey'
arch=(x86_64)
license=(LGPL-3.0-or-later)
depends=(
  glibc
  libgcc
  libstdc++
  nitrokey-udev-rules
)
makedepends=(
  catch2
  cmake
  hidapi
  libusb
  qt5-tools
)
provides=(libnitrokey.so)
source=(
  $url/releases/download/v$pkgver/$pkgname-v$pkgver.tar.gz{,.sig}
  $pkgname-3.8-catch2_3.patch
  $pkgname-3.8-do-not-install-udev-rules.patch
)
sha512sums=('f986824f642f36d3cd3ebad6fd354843f6045cf0b619f971394f3917f1b447ca72ddae81bdba1886ac95797f07a90b702441b81cbc47618ee629364a348f96ff'
            'SKIP'
            '4e553aa186a9b58e407edcd383250d62bf5cf786d7b2ed19096c755889fe85a7b7dd5a38ffa816aa43a5e5eaf54246959684cd9b66a483e2d48ba1ad3bc3b667'
            '81b3b3a47fa6bae19dd395b8f2b2a7ebac8566882c7c1dc66e1c19df57c4bcf33dad44273ab7b65b8828b24af756599f09c9d94b583635a9174918686f79dd58')
b2sums=('02166746e5607bad6d89af195ca08a9c780801e5416553a290ac5e8875aa769afdcec4ad02da8ab85f66ea85772780ac74eacde06c0be5a094247c7421104262'
        'SKIP'
        'e0c968aa7181faeabfd2c4f80912f5c5a15ca4a8514228fe9b3d77b1bfa836e4d249e51ffe04a5988f5ee183de45c959575278a4c72f60c5d4c78c67efe70ebb'
        '7ac4a0de9afc7f8f984c58c0dbc2ce6a1672c0875d39bc6da365128f6ae249398bfdddb79eecccedf8a0a2c3d429fa68d8af5032bdb3f3f27cb642b208c26aec')
validpgpkeys=('868184069239FF65DE0BCD7DD9BAE35991DE5B22') # Szczepan Zalega <szczepan.zalega@gmail.com>

prepare() {
  patch -Np1 -d $pkgname-v$pkgver -i ../$pkgname-3.8-catch2_3.patch
  patch -Np1 -d $pkgname-v$pkgver -i ../$pkgname-3.8-do-not-install-udev-rules.patch
}

build() {
  local cmake_options=(
    -B build
    -D BUILD_SHARED_LIBS=ON
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
    -D CMAKE_INSTALL_LIBDIR=lib
    -D CMAKE_POLICY_VERSION_MINIMUM=3.5
    -D COMPILE_OFFLINE_TESTS=ON
    -S $pkgname-v$pkgver
    -W no-dev
  )

  cmake "${cmake_options[@]}"
  cmake --build build --verbose
}

check() {
  ctest --test-dir build --output-on-failure
}

package() {
  depends+=(
    hidapi libusb libhidapi-libusb.so
  )

  DESTDIR="$pkgdir" cmake --install build
}

# vim: ts=2 sw=2 et:
