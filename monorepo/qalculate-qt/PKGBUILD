# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=qalculate-qt
pkgver=5.13.0
pkgrel=1
pkgdesc='Qt frontend for libqalculate'
arch=(x86_64)
url='https://qalculate.github.io/'
license=(GPL-2.0-only)
depends=(glibc
         libgcc
         libqalculate
         libstdc++
         qt6-base
         qt6-svg)
makedepends=(git
             qt6-tools)
source=(git+https://github.com/Qalculate/qalculate-qt#tag=v$pkgver)
sha256sums=('e418df85282423509f9664528744384a401221d0760e7a284f7e15dc34a180f1')

build() {
  cd $pkgname
  qmake6 \
    PREFIX=/usr
  make
}

package() {
  cd $pkgname
  make INSTALL_ROOT="$pkgdir" install
}
