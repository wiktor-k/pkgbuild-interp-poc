# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
pkgname=collada-dom
pkgver=2.5.0
pkgrel=5
pkgdesc="COLLADA Document Object Model (DOM) C++ Library"
arch=('x86_64')
url="https://github.com/rdiankov/collada-dom/"
license=('LicenseRef-collada')
depends=(
  'boost-libs'
  'minizip'
  'libxml2'
  'libgcc'
  'libstdc++'
  'glibc'
)
makedepends=('cmake' 'boost')
source=("$pkgname-$pkgver.tar.gz::https://github.com/rdiankov/collada-dom/archive/refs/tags/v${pkgver}.tar.gz"
        "$pkgname-fix-boost1.85.patch::https://github.com/rdiankov/collada-dom/pull/43.patch"
         boost-1.89.patch
         "$pkgname-fix-gcc16-ambiguous-overload.patch")
sha512sums=('f189d09e2396faa266734981bb7b5e91ec34b6faa9ad340206e769dae316496bf4271c129980668dc2756874dbb8c1157162197d0d3a74075e35200821875156'
             'b1f742316cecf4a932df7ba86f2d3bc7dda8a0cae0b0300c0638e8858e654ce5de8bc351c094f99bb6a9fba693f272de6969411679ac2c4bfe99bb9b95d1ed79'
             '24bddd4c9663cf763114d5371f89f569ff320a9bee757e2d701913dbdaad1327f980e0c877900beb47fdfeca08b96fae2015699810cac269bcd73b2a879b9632'
             '8f81440ec43b029aac66ffdb2a0462cc41dbe9e099bae16319d1ab0944d6be8aa51730b91515f557c951fa4c73a8de447f83c9dc4329cae0b0f6fa60f65e6b6e')

prepare() {
  cd "$pkgname-$pkgver"

  patch -Np1 -i "$srcdir/$pkgname-fix-boost1.85.patch"
  patch -p1 -i ../boost-1.89.patch
  patch -p1 -i ../"$pkgname-fix-gcc16-ambiguous-overload.patch"
  sed -i 's|<unzip.h>|<minizip/unzip.h>|' dom/include/dae/daeZAEUncompressHandler.h
  sed -i 's|<zip.h>|<minizip/zip.h>|' dom/src/modules/LIBXMLPlugin/daeLIBXMLPlugin.cpp
}

build() {
  export CMAKE_POLICY_VERSION_MINIMUM=3.5

  local cmake_options=(
    -B build
    -S $pkgname-$pkgver
    -W no-dev
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
  )
  cmake "${cmake_options[@]}"
  cmake --build build
}

package() {
  install -Dm644 "${pkgname}-${pkgver}/licenses/license_e.txt" "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"

  DESTDIR="$pkgdir" cmake --install build
}
