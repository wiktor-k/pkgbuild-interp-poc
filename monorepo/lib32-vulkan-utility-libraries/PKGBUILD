# Maintainer: Jan Alexander Steffens (heftig) <heftig@archlinux.org>
# Maintainer: Laurent Carlier <lordheavym@gmail.com>
# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Bartosz Taudul <wolf@nereid.pl>

pkgname=lib32-vulkan-utility-libraries
pkgver=1.4.363.0
pkgrel=1
pkgdesc="Vulkan Utility Libraries (32-bit)"
url="https://www.vulkan.org/"
arch=(x86_64)
license=(Apache-2.0)
depends=(
  lib32-glibc
  lib32-vulkan-icd-loader
  vulkan-utility-libraries
)
makedepends=(
  cmake
  git
  ninja
  python
  vulkan-headers
)
options=(
  # We are producing static libraries
  !lto
)
source=("git+https://github.com/KhronosGroup/Vulkan-Utility-Libraries#tag=vulkan-sdk-$pkgver")
b2sums=('22851420b0291e30a909c6b57186170d2d76cc34176198531b3178b260af5e93cb8999f3e617650af0a3fd6f21fc60fb3bd60187ac5fb7b46defa00016786a6f')

build() {
  local cmake_options=(
    -D CMAKE_BUILD_TYPE=Release
    -D CMAKE_INSTALL_LIBDIR=lib32
    -D CMAKE_INSTALL_PREFIX=/usr
    -D CMAKE_INSTALL_SYSCONFDIR=/etc
    -D CMAKE_SKIP_INSTALL_RPATH=ON
  )

  export ASFLAGS+=" --32"
  export CFLAGS+=" -m32"
  export CXXFLAGS+=" -m32"
  export PKG_CONFIG="i686-pc-linux-gnu-pkg-config"

  cmake -S Vulkan-Utility-Libraries -B build -G Ninja "${cmake_options[@]}"
  cmake --build build
}

check() {
  ctest --test-dir build --output-on-failure --stop-on-failure
}

package() {
  DESTDIR="$pkgdir" cmake --install build
  rm -r "$pkgdir/usr/include"
}

# vim:set sw=2 sts=-1 et:
