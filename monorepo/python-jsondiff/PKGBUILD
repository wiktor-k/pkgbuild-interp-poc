# Maintainer: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Guillaume Horel <guillaume.horel@gmail.com>

pkgname=python-jsondiff
_pkgname=jsondiff
pkgver=2.2.1
pkgrel=3
pkgdesc='Diff JSON and JSON-like structures in Python'
arch=(any)
url='https://github.com/xlwings/jsondiff'
license=(MIT)
depends=(python python-yaml)
makedepends=(git python-build python-installer python-setuptools python-wheel python-setuptools-scm)
checkdepends=(python-pytest python-hypothesis)
source=("git+https://github.com/xlwings/jsondiff.git#tag=$pkgver")
sha256sums=('a8a5e6340365f0e1268096e32c5c602ab7100778d0d3dc3ac6fc705102994308')

build() {
  cd $_pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pkgname
  pytest -v tests
}

package() {
  cd $_pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
