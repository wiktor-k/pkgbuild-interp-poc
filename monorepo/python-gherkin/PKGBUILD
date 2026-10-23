# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-gherkin
pkgver=42.0.1
pkgrel=1
pkgdesc="A parser and compiler for the Gherkin language"
arch=('any')
url="https://github.com/cucumber/gherkin"
license=('MIT')
depends=(
  'python'
  'python-typing_extensions'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-uv-build'
  'python-wheel'
)
checkdepends=('python-pytest')
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('69c399cf3873946eb7bb432a2f8689bcf74adfdcff3cfccb2f7d235b003c7279b4751b5fb75dba8bcdd78c372b897ef96b288dd8c731d33a963d3867ec1134a3')

build() {
  cd "${pkgname#python-}-$pkgver/python"
  python -m build --wheel --no-isolation
}

check() {
  cd "${pkgname#python-}-$pkgver/python"
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd "${pkgname#python-}-$pkgver/python"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
