# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-uncertainties
pkgver=3.2.4
pkgrel=1
pkgdesc="Transparent calculations with uncertainties on the quantities involved (aka error propagation); fast calculation of derivatives."
arch=('any')
license=('BSD-2-Clause')
url="https://github.com/lmfit/uncertainties"
depends=('python')
optdepends=('python-numpy: additional support for NumPy arrays and matrices')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-setuptools-scm' 'python-wheel')
checkdepends=('python-pytest' 'python-numpy' 'python-scipy')
source=("git+https://github.com/lmfit/uncertainties.git#tag=$pkgver")
sha512sums=('0741d308916392f74798123f031851aef55fb6610ea5365a65b3d96e2bcbe40788bdbcb9bd3b7f7af033f44a9987940056b7bbb96da23774f27f80449f908046')

build() {
  cd uncertainties
  python -m build -nw
}

check() {
  cd uncertainties
  python -m pytest
}

package() {
  cd uncertainties
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
