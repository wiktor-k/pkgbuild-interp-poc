# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Filipe Nascimento <flipee at tuta dot io>
# Contributor: Julien Nicoulaud <julien DOT nicoulaud AT gmail DOT com>

pkgname=python-cachelib
pkgver=0.17.0
pkgrel=1
pkgdesc="A collection of cache libraries in the same API interface. Extracted from werkzeug"
arch=(any)
url="https://github.com/pallets/cachelib"
license=(BSD-3-Clause)
depends=(python)
makedepends=(
  python-build
  python-flit-core
  python-installer
  python-wheel
)
checkdepends=(
  memcached
  python-boto3
  python-pylibmc
  python-pymongo
  python-pytest-xprocess
  python-redis
  valkey
)
optdepends=(
  'python-boto3: for S3 backend'
  'python-pylibmc: for memcached backend'
  'python-pymongo: for MongoDB backend'
  'python-redis: for Redis backend'
)
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('d0ced891e0779215ed58e40b9d2da1c88287973c31cc00aa1a13ad7ce438a7bbe5d3092afdb033032307599db64ad4992b1ba5d0014685a660c20a104c4128f8')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # The Python valkey client is not packaged in the repositories.
  test-env/bin/python -m pytest \
    --deselect tests/test_dynamodb_cache.py \
    --deselect tests/test_mongodb_cache.py \
    --deselect tests/test_valkey_cache.py
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
