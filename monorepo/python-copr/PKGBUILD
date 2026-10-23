# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-copr
pkgver=2.7
pkgrel=1
pkgdesc="Python interface for Copr"
url="https://github.com/fedora-copr/copr"
license=('GPL-2.0-or-later')
arch=('any')
depends=(
  'python'
  'python-filelock'
  'python-munch'
  'python-requests'
  'python-requests-gssapi'
  'python-requests-toolbelt'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=('python-pytest')
source=("git+$url.git#tag=python-copr-$pkgver-1")
b2sums=('eabb29f5041e0469ae3b0b490d5edcce20960f1d142fb5c07b1164968ab347aa61ba0b4f2a88ff4796dddd8317f41edeb06646243fc5111897bd664ff5a4cec9')

build() {
  cd copr/python
  python -m build --wheel --no-isolation
}

check() {
  cd copr/python
  pytest
}

package() {
  cd copr/python
  python -m installer --destdir="$pkgdir" dist/*.whl
}
