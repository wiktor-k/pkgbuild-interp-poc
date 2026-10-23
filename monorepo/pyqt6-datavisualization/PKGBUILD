# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgbase=pyqt6-datavisualization
pkgname=(python-pyqt6-datavisualization)
pkgver=6.11.0
pkgrel=1
pkgdesc='Python bindings for QtDataVisualization'
arch=(x86_64)
url='https://www.riverbankcomputing.com/software/pyqtdatavisualization/intro'
license=(GPL3)
groups=(pyqt6)
depends=(glibc
         libstdc++
         python
         python-pyqt6
         qt6-base
         qt6-datavis3d)
makedepends=(pyqt-builder
             qt6-declarative
             sip)
source=(https://pypi.python.org/packages/source/P/PyQt6-DataVisualization/pyqt6_datavisualization-$pkgver.tar.gz)
sha256sums=('866f2f77ce162d296fc6957de30b59476190e8505dfdb42b77288d14e60086ee')

build() {
  cd pyqt6_datavisualization-$pkgver
  sip-build \
    --no-make \
    --qmake=/usr/bin/qmake6 \
    --pep484-pyi
  cd build
  make
}

package_python-pyqt6-datavisualization(){
  cd pyqt6_datavisualization-$pkgver/build
  make INSTALL_ROOT="$pkgdir" install
}
