# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Contributor: Bartłomiej Piotrowski <bpiotrowski@archlinux.org>
# Contributor: naelstrof <naelstrof@gmail.com>

pkgname=maim
pkgver=5.8.2
pkgrel=1
pkgdesc='Utility to take a screenshot using imlib2'
arch=('x86_64')
url='https://github.com/naelstrof/maim'
license=('GPL3')
depends=('freetype2' 'libpng' 'libjpeg-turbo' 'libwebp' 'libx11' 'libxcomposite' 'libxext' 'libxfixes' 'libxrandr' 'slop')
makedepends=('cmake' 'glm')
source=("${pkgname}-${pkgver}.tar.gz"::"$url/archive/v${pkgver}.tar.gz")
sha256sums=('2252c95c7eeebdee190c0a8cae3adeb1bb9c70858aef807060669fd02180204c')

prepare() {
  cd "$pkgname-$pkgver"
  sed -e 's|CXX_STANDARD 11|CXX_STANDARD 17|g' -i CMakeLists.txt
}

build() {
  cd "$pkgname-$pkgver"
  cmake -DCMAKE_INSTALL_PREFIX=/usr .
  make
}

package() {
  make -C "$pkgname-$pkgver" DESTDIR="$pkgdir" install
}
