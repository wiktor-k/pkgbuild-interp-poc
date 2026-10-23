# Maintainer: Maxime Gauduin <alucryd@archinux.org>
# Contributor: xiretza <xiretza+aur@xiretza.xyz>

pkgname=python-cmake-build-extension
pkgver=0.6.1
pkgrel=3
pkgdesc='Setuptools extension to build and package CMake projects'
arch=(any)
url=https://github.com/diegoferigo/cmake-build-extension
license=(MIT)
depends=(
  cmake
  python
  python-gitpython
  python-setuptools
  python-setuptools-scm
  python-wheel
  ninja
)
makedepends=(
  git
  python-build
  python-installer
)
_tag=7a5deb54ee165aa157e33937420c933d0c2c41c6
source=(git+https://github.com/diegoferigo/cmake-build-extension.git#tag=${_tag})
b2sums=('f0cae38c14b038bfb50cfb0ed2231a536ae59efd3ef5f2ef92c0e9b7335d7ed8bf40a9d396ec7ab4df5e072b5ff7cef9dee7944e23e21cd320690668b7bed73a')

pkgver() {
  cd cmake-build-extension
  git describe --tags | sed 's/^v//'
}

build() {
  cd cmake-build-extension
  python -m build -wn
}

package() {
  python -m installer --destdir="${pkgdir}" cmake-build-extension/dist/*.whl
  install -Dm 644 cmake-build-extension/LICENSE -t "${pkgdir}"/usr/share/licenses/python-cmake-build-extension/
}

# vim: ts=2 sw=2 et
