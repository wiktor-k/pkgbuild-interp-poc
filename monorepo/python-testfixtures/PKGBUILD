# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-testfixtures
pkgver=12.4.0
pkgrel=1
pkgdesc="A collection of helpers and mock objects that are useful when writing unit tests or doc tests"
arch=('any')
license=('MIT')
url="https://github.com/Simplistix/testfixtures"
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-hatchling')
checkdepends=('python-pytest' 'python-sybil' 'python-pytest-django' 'python-click'
              'python-django' 'python-loguru' 'python-numpy' 'python-pandas'
              'python-polars' 'python-pydantic' 'python-structlog' 'python-twisted')
source=("git+https://github.com/Simplistix/testfixtures.git#tag=$pkgver")
sha512sums=('7d5dc6513fa8fd34a894cfeab8b89fd4deabd67338130500b983fa23111a113bafb77f5b2e670b893f926951acf8a405cd90c1598bad6beeda47da8c0910a681')

build() {
  cd testfixtures
  python -m build --wheel --no-isolation
}

check() {
  cd testfixtures
  PYTHONPATH="$PWD/src" pytest --ignore=build
}

package() {
  cd testfixtures
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
