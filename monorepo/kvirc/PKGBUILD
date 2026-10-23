# Maintainer: Balló György <ballogyor+arch at gmail dot com>

pkgname=kvirc
pkgver=5.2.10
pkgrel=2
pkgdesc='Qt based IRC client'
arch=('x86_64')
url='https://www.kvirc.net/'
license=('GPL2')
depends=('enchant' 'hicolor-icon-theme' 'qt6-multimedia' 'qt6-webengine' 'qt6-5compat')
makedepends=('cmake' 'ninja' 'python')
optdepends=('perl: scripting support'
            'python: scripting support')
source=("https://github.com/kvirc/KVIrc/archive/$pkgver/$pkgname-$pkgver.tar.gz"
        'kvirc.metainfo.xml')
sha256sums=('6b448b08aeaf5fd3f2b120525c8c87ddad3336f5b750a557b2022aaad0dda626'
            '5f90d6161f3ed9527375d210ba30dde5f2d8e68b7c9bef9419d145b3bbbe1d24')

prepare() {
  cd KVIrc-$pkgver
  sed -i '/\/share\/pixmaps\//d' data/icons/48x48/CMakeLists.txt
}

build() {
  cmake -S KVIrc-$pkgver -B build -G Ninja -DCMAKE_INSTALL_PREFIX='/usr' \
    -DCMAKE_SKIP_RPATH=YES \
    -DWANT_KDE=OFF \
    -DWANT_PHONON=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
  install -Dm644 $pkgname.metainfo.xml "$pkgdir/usr/share/metainfo/net.kvirc.KVIrc5.metainfo.xml"
}
