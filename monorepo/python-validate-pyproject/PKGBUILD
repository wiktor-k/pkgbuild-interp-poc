# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Jelle van der Waa <jelle@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-validate-pyproject
pkgver=0.26
pkgrel=1
pkgdesc="Validation library and CLI tool for checking on 'pyproject.toml' files using JSON Schema"
arch=('any')
url="https://github.com/abravalheri/validate-pyproject"
license=('MPL-2.0')
depends=(
  'python'
  'python-fastjsonschema'
  'python-packaging'
  'python-trove-classifiers'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools-scm'
  'python-wheel'
)
checkdepends=('python-pytest')
source=("git+$url.git#tag=$pkgver")
b2sums=('e48adc8783d0f6912e829403bcd059635de5e54a8683d9e77af28f8d8b91d4e965a79027863c07688084413ebf2f5792d0a69e786af5a6540d467c349cb46694')

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest --override-ini="addopts="
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
}
