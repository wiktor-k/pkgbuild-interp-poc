# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: renek <aur@spaceshore.net>

pkgname=python-multidict
pkgver=7.0.0
pkgrel=1
pkgdesc='Asyncio-based multidict implementation for Python'
arch=('x86_64')
url='https://github.com/aio-libs/multidict'
license=('Apache-2.0')
depends=(
  'glibc'
  'python'
)
makedepends=(
  'cython'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-coverage'
  'python-objgraph'
  'python-perf'
  'python-psutil'
  'python-pytest'
  'python-pytest-codspeed'
)
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('7521ab5a512138aaeaacebeecb9f5c016071157392886750a284b86bb23553e528e818ebcc3fe51d6c53f4fce05408366db38ebc5965aa964b32b345af67333c')
b2sums=('1d89bf9ec9c28487598e1bbaaa12d451cc6c05f0bd62ddbabeb4beea5e11b6a5515cdf0188d0bc01df6b5a2f73c5c360a7c6e8d5266354a3a50c441ddccf1cbb')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  cd tests
  ../test-env/bin/python -m pytest -v --override-ini="addopts="
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
