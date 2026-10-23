# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-semver
pkgver=3.1.0
pkgrel=1
pkgdesc="Python helper for Semantic Versioning"
url="https://github.com/python-semver/python-semver"
license=('BSD-3-Clause')
arch=('any')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-setuptools-scm'
             'python-wheel')
checkdepends=('python-pytest')
source=("git+https://github.com/python-semver/python-semver.git#tag=$pkgver")
sha512sums=('9343acafb568e8aa1b47cb25309308ee47d2f9da2db782a6d1869f7f8a883ca09d68b7a2694bc787bdb1692563001a8d1f66c2e6728efe8c125462e91e5ef30e')

build() {
  cd python-semver
  python -m build -nw
}

check() {
  cd python-semver
  pytest -o addopts=''
}

package() {
  cd python-semver
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
