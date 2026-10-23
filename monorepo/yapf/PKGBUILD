# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=yapf
pkgver=0.43.0
pkgrel=2
pkgdesc="Python style guide checker"
arch=('any')
url="https://github.com/google/yapf"
license=('Apache-2.0')
depends=('python' 'python-platformdirs')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools')
checkdepends=('python-nose')
source=("git+https://github.com/google/yapf.git#tag=v$pkgver")
sha512sums=('d60e01307bb29a1bd065b9d1f06705ca79421048b0c6adf4f1745a76bdc746c51911493cdb151a4672bc47b6d95591eb54323d71564ec8bd95891b9bf446dbc6')

build() {
  cd yapf
  python -m build -nw
}

check() {
  cd yapf
  PYTHONPATH="$PWD/third_party" nosetests3
}

package() {
  cd yapf
  python -m installer -d "$pkgdir" dist/*.whl
}
