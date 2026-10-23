# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-validators
pkgver=0.36.0
pkgrel=1
pkgdesc="Python Data Validation for Humans"
url="https://python-validators.github.io/validators"
license=('MIT')
arch=('any')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools')
checkdepends=('python-pytest' 'python-eth-hash')
optdepends=('python-eth-hash: for eth_address validator')
source=("git+https://github.com/python-validators/validators.git#tag=$pkgver")
sha512sums=('db242203efddca51f215831bae9e56c71d324393da496806b295a71432ac933a93bc0d448005d0d2fc34be994d1ec3cd40aea525c55fe6c9072cd8476a1a3ac1')

build() {
  cd validators
  python -m build -nw
}

check() {
  cd validators
  python -m pytest
}

package() {
  cd validators
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
