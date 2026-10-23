# Maintainer: Jan Alexander Steffens (heftig) <heftig@archlinux.org>

pkgbase=pam_wrapper
pkgname=(
  pam_wrapper
  pam_wrapper-docs
)
pkgver=1.1.8
pkgrel=2
pkgdesc="Tool to test PAM applications and PAM modules"
url="https://cwrap.org/pam_wrapper.html"
arch=(x86_64)
license=(GPL-3.0-or-later)
depends=(
  glibc
  pam
  python
)
makedepends=(
  cmake
  cmocka
  doxygen
  git
  graphviz
  ninja
  python-setuptools
)
source=(
  "git+https://git.samba.org/pam_wrapper.git#tag=pam_wrapper-$pkgver"
)
b2sums=('69ae8379e6134896b47cfeff1a2cb42f9b072ddbd6871991a06e1159aa987001baa9331bb242457faebf8827bf0440c258800ca243735465aaaa6792cbc93e82')

prepare() {
  cd $pkgname
}

build() {
  local cmake_options=(
    -D CMAKE_INSTALL_PREFIX=/usr
    -D CMAKE_BUILD_TYPE=None
    -D UNIT_TESTING=true
  )

  cmake -S $pkgname -B build -G Ninja "${cmake_options[@]}"
  cmake --build build
  cmake --build build --target doc
}

check() {
  ctest --test-dir build --output-on-failure --stop-on-failure -j$(nproc)
}

package_pam_wrapper() {
  provides=(
    libpam_wrapper.so
    libpamtest.so
  )

  DESTDIR="$pkgdir" cmake --install build
}

package_pam_wrapper-docs() {
  pkgdesc+=" (documentation)"
  depends=()

  mkdir -p "$pkgdir"/usr/share/{doc,man}
  cp -a build/doc/html -T "$pkgdir/usr/share/doc/$pkgname"
  cp -a build/doc/man/man3 -t "$pkgdir/usr/share/man"
}

# vim:set sw=2 sts=-1 et:
