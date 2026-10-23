# Maintainer: Leonidas Spyropoulos <artafinde@archlinux.org>
# Contributor: Patrick Ziegler <p.ziegler96@gmail.com>

pkgname=polybar
pkgver=3.7.2
pkgrel=3
pkgdesc="A fast and easy-to-use status bar"
arch=('x86_64')
url="https://github.com/polybar/polybar"
backup=("etc/polybar/config.ini")
license=("MIT")
depends=("libuv" "cairo" "xcb-util-image" "xcb-util-wm" "xcb-util-xrm" "xcb-util-cursor"
         "alsa-lib" "libpulse" "libmpdclient" "libnl" "jsoncpp" "curl")
makedepends=("cmake" "python" "pkg-config" "python-sphinx" "python-packaging" "i3-wm"
             "ninja" "git")
optdepends=("i3-wm: i3 module support")
source=("${pkgname}-${pkgver}.tar.gz::${url}/releases/download/${pkgver}/${pkgname}-${pkgver}.tar.gz"
        "${pkgname}-${pkgver}.tar.gz.asc::${url}/releases/download/${pkgver}/${pkgname}-${pkgver}.tar.gz.asc"
         https://github.com/polybar/polybar/commit/e7cf47de71a7ceaf31c2093a959ba4aa08ec96a7.patch
         https://github.com/polybar/polybar/commit/455d63611a695ffe4fe372272be28f4f6d9b90d3.patch)
sha256sums=('e2feacbd02e7c94baed7f50b13bcbf307d95df0325c3ecae443289ba5b56af29'
            'SKIP'
            'd6be29e4eeb7c54b9f1fd714a7b2d5ab6639f3aeab22cfc763d8a2d9bb962c6a'
            'fbd0ff540e769284b4abcee9b4067bee7b0bad1c2ee9cca9f146e2faaa89e45c')
validpgpkeys=('1D5791352D51A228D4DDDBA4521E5E03AEBCA1A7') # Patrick Ziegler <p.ziegler96@gmail.com>

prepare() {
  cd "${pkgname}-${pkgver}"
  # Broken tests from executing inside systemd-nspawn
  sed -i 's/is_async/DISABLED_is_async/' \
    tests/unit_tests/utils/process.cpp

  # Fix build with GCC 15
  patch -p1 -i ../e7cf47de71a7ceaf31c2093a959ba4aa08ec96a7.patch

  # Fix freetype linking
  patch -p1 -i ../455d63611a695ffe4fe372272be28f4f6d9b90d3.patch
}

build() {
  cd "${pkgname}-${pkgver}"
  cmake \
    -Bbuild \
    -GNinja \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DBUILD_DOC_HTML=OFF \
    -DBUILD_TESTS=ON \
    -DPYTHON_EXECUTABLE=/usr/bin/python3
  cmake --build build --verbose
}

check() {
  cd "${pkgname}-${pkgver}"
  ninja check -C build
}

package() {
  cd "${pkgname}-${pkgver}"
  DESTDIR="${pkgdir}" cmake --install build
  install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}

# vim:set ft=sh sw=2 sts=2 et:
