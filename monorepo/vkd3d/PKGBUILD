# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Bruce Zhang

pkgbase=vkd3d
pkgname=(
  vkd3d
  vkd3d-docs
)
pkgver=1.19
pkgrel=1
pkgdesc="Direct3D 12 to Vulkan translation library By WineHQ"
url="https://wiki.winehq.org/Vkd3d"
arch=(x86_64)
license=(LGPL-2.1-or-later)
depends=(
  glibc
  ncurses
  spirv-tools
)
makedepends=(
  doxygen
  git
  graphviz
  mesa
  perl
  perl-json
  spirv-headers
  vulkan-headers
  vulkan-icd-loader
  wine
  xcb-util-keysyms
  xorgproto
)
source=(
  "git+https://gitlab.winehq.org/wine/vkd3d.git?signed#tag=vkd3d-$pkgver"
)
b2sums=('89c5b3521491cca9a1ee7633f8c3258f05f1b04aa53c11b8ff1a5afa0803788f8968d3ad0bc370927c9a41eabc5bd7c416e750780d81f04b52f337095ef0555a')
validpgpkeys=(
  277EB71D1BEBFC86E0E14983C380B588AA8A1D1C # Henri Verbeet <hverbeet@locutus.nl>
)

prepare() {
  cd vkd3d
  ./autogen.sh
}

build() {
  local configure_options=(
    --prefix=/usr
    --disable-doxygen-pdf
    --enable-doxygen-doc
    --enable-doxygen-dot
    --enable-doxygen-html
    --enable-tests
    --with-spirv-tools
  )

  cd vkd3d
  ./configure "${configure_options[@]}"
  make
}

check() {
  cd vkd3d
  make check
}

package_vkd3d() {
  provides=(libvkd3d{,-shader,-utils}.so)

  cd vkd3d
  make DESTDIR="$pkgdir" install
}

package_vkd3d-docs() {
  pkgdesc+=" - API documentation"
  depends=()

  cd vkd3d
  mkdir -p "$pkgdir/usr/share/doc"
  cp -aT doc/html "$pkgdir/usr/share/doc/$pkgbase"
}

# vim:set sw=2 sts=-1 et:
