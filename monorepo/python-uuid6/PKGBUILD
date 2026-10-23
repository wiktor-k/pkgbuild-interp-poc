# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-uuid6
pkgver=2025.0.1
pkgrel=1
pkgdesc='New time-based UUID formats which are suited for use as a database key'
arch=('any')
url='https://github.com/oittaa/uuid6-python'
license=('MIT')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-setuptools-scm' 'python-wheel')
source=("git+https://github.com/oittaa/uuid6-python#tag=$pkgver")
sha512sums=('592ce9a44d49be80410397b86a5172c474a913e9399e2c3198458b4708ec8ee336d981b5c8767d2fbb44d6e6b670b093b10f75b83facf9dd417abe9d7ab9acb2')

build() {
  cd uuid6-python

  python -m build --wheel --no-isolation
}

check() {
  cd uuid6-python

  PYTHONPATH=src python -m unittest discover -s test
}

package() {
  cd uuid6-python

  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
