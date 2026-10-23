# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgbase=pyqt6-networkauth
pkgname=(python-pyqt6-networkauth)
pkgver=6.11.0
pkgrel=1
pkgdesc='Python bindings for QtNetworkAuth'
arch=(x86_64)
url='https://www.riverbankcomputing.com/software/pyqtnetworkauth/intro'
license=(GPL3)
groups=(pyqt6)
depends=(glibc
         libstdc++
         python
         python-pyqt6
         qt6-base
         qt6-networkauth)
makedepends=(pyqt-builder
             sip)
source=(https://pypi.python.org/packages/source/P/PyQt6-NetworkAuth/pyqt6_networkauth-$pkgver.tar.gz)
sha256sums=('03650f0b0d709284f4ede3b31170d78b996864406bf34a4a0e76474b9144ed2a')

build() {
  cd pyqt6_networkauth-$pkgver
  sip-build \
    --no-make \
    --qmake=/usr/bin/qmake6 \
    --pep484-pyi
  cd build
  make
}

package_python-pyqt6-networkauth(){
  cd pyqt6_networkauth-$pkgver/build
  make INSTALL_ROOT="$pkgdir" install
}
