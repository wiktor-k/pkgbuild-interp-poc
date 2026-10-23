# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-elastic-transport
_pkgname=elastic-transport-python
pkgver=9.4.2
pkgrel=1
pkgdesc="Transport classes and utilities shared among Python Elastic client libraries"
arch=(any)
url="https://github.com/elastic/elastic-transport-python"
license=(Apache-2.0)
depends=(
  python
  python-aiohttp
  python-certifi
  python-httpx
  python-orjson
  python-requests
  python-sniffio
  python-urllib3
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-asyncio
  python-pytest-httpbin
  python-pytest-httpserver
  python-pytest-rerunfailures
  python-respx
  python-trio
  python-trustme
)
source=(
  "$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz"
  "$pkgname-update-debug-logging-tests.patch::$url/commit/3e100bf9dca8724a38c65f3953dcf0d103cacfde.patch"
)
b2sums=('fde26b2d8f2adff6156ada512f4a36a09cc99fb3c85d4a72b7002e3f08e73d015df29e67db97adbc304576cf499f758987a0320a250f14c184d4bb958286a80d'
        '87cb5104d1bfc54c7769ed9a8400ce1f560ad3f900ea76d5fe18270cfbe1c79965bbedad4c8909054086f3882a3be62de67b87c4ddd8acecc912c7f27030bfe4')

prepare() {
  cd $_pkgname-$pkgver
  # httpbin 0.10.3 switched from the JSONIFY_PRETTYPRINT_REGULAR setting to
  # app.json.compact, so its JSON responses are now indented.
  patch -Np1 < ../$pkgname-update-debug-logging-tests.patch
}

build() {
  cd $_pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_pkgname-$pkgver
  # The calls to httpbin.org, and as such the tests, are flaky when running on
  # build.archlinux.org.
  # OpenTelemetry not in Arch repos (yet).
  pytest --override-ini="addopts=" \
    --reruns 1 \
    --reruns-delay 1 \
    --ignore tests/test_otel.py
}

package() {
  cd $_pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
