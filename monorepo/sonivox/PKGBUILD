# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=sonivox
pkgver=4.0.2
pkgrel=1
pkgdesc="Fork of the AOSP 'platform_external_sonivox' to use outside of Android"
arch=(x86_64)
url="https://github.com/EmbeddedSynth/sonivox"
license=(Apache-2.0)
depends=(
  glibc
  zlib
)
makedepends=(
  cmake
  gtest
)
provides=(libsonivox.so)
source=(
  $pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz
  https://www.ronimusic.com/sf2/Airfont_340.dls  # soundfont required for testing...
)
sha256sums=('5c513d7a6fa0ebf61afd3b96da881e3014edf53acfb8d08a63dd853fca37c1fa'
            'beb3e39e3c9fc51ef4dff36fdd8db0361471a91d244c3ee78af90f6d3c783b04')
b2sums=('4356ce1371c803d1b488fbe7f1de447bd15a27918be28629583ce7452255ae9c5a99367be631f2ce77af40fedc8fb945f62051748d241f277cb73492c0cf7f18'
        '298ea526ddd68367817ffbec77e8f338bbc167a608f3d670a96d2c517923cb72e44df9bde5aafd97faad0139789732dc2975be9f9d7bee3c7c6cecc78014bf47')

prepare() {
  # Symlink the soundfont required for unit tests into place:
  # https://github.com/pedrolcl/sonivox/issues/87
  ln -s Airfont_340.dls soundfont.dls
}

build() {
  local cmake_options=(
    -B build
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
    -S $pkgname-$pkgver
    -W no-dev
  )

  # This ensures the soundfont used in unit tests is found and not downloaded during build.
  export TEMP="$srcdir"
  cmake "${cmake_options[@]}"
  cmake --build build --verbose
}

check() {
  ctest --test-dir build --output-on-failure
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
