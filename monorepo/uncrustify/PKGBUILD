# Maintainer: Santiago Torres <sangy@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Pierre Neidhardt <ambrevar@gmail.com>
# Contributor: Auguste Pop <auguste [at] gmail [dot] com>

pkgname=uncrustify
pkgver=0.83.0
pkgrel=1
pkgdesc="A source code beautifier"
arch=(x86_64)
url="https://github.com/uncrustify/uncrustify"
license=(GPL-2.0-only)
depends=(
  glibc
  libgcc
  libstdc++
)
makedepends=(
  cmake
  python
)
source=("$url/archive/$pkgname-$pkgver.tar.gz")
b2sums=('317d9e3b0bb4c21bc8ad6a9d3ed9acd6daf24c361db97b7e6a58e8f6496675cd8f02a4c15bace6656e0e95409f44b6e3a23ad8bad69144503561f9ac51bcf038')

build() {
  cmake -S $pkgname-$pkgname-$pkgver -B build \
    -DCMAKE_BUILD_TYPE=None \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -Wno-dev
  cmake --build build
}

check() {
  ctest --test-dir build --output-on-failure
}

package() {
  DESTDIR="$pkgdir" cmake --install build
  cd $pkgname-$pkgname-$pkgver
  install -vDm644 -t "$pkgdir/usr/share/uncrustify" etc/*.cfg
}
