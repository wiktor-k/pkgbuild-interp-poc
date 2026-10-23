# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-jsonlines
pkgver=4.0.0
pkgrel=6
pkgdesc="Library with helpers for the jsonlines file format"
url="https://github.com/wbolster/jsonlines"
license=('BSD-3-Clause')
arch=('any')
depends=('python-attrs')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-pytest')
source=("git+https://github.com/wbolster/jsonlines.git#tag=$pkgver")
sha512sums=('4c3324884d0bf63e7da988ca93061cd98d0ef0f34781dff6eb5d0de57d31760d03e8f9239cbd2cb348a881a89826b68a43712a68bac03f23433e60a16cbb171c')

build() {
  cd jsonlines
  python -m build --wheel --no-isolation
}

check() {
  cd jsonlines
  python -m pytest
}

package() {
  cd jsonlines
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE.rst -t "$pkgdir"/usr/share/licenses/$pkgname/
}
