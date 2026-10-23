# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-geoip2
pkgver=5.3.0
pkgrel=1
pkgdesc="MaxMind GeoIP2 API"
url="https://github.com/maxmind/GeoIP2-python"
license=('Apache')
arch=('any')
depends=(
  python
  python-aiohttp
  python-maxminddb
  python-requests
  python-setuptools
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools-scm
  python-uv-build
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-httpserver
)
source=(
  git+https://github.com/maxmind/GeoIP2-python.git#tag=v$pkgver
  git+https://github.com/maxmind/MaxMind-DB.git
)
sha512sums=('d2fd9cd7c053e3342738ead97c21dc4df97498138a338e939e97ff899914284f94f877d8ed9230efa3c1f89ed3439ac2234dad56f143ed0810133377793be7a4'
            'SKIP')

prepare() {
  cd GeoIP2-python
  git submodule init
  git config submodule.tests/data.url "$srcdir/MaxMind-DB"
  git -c protocol.file.allow=always submodule update
}

build() {
  cd GeoIP2-python
  python -m build --wheel --no-isolation
}

check() {
  cd GeoIP2-python

  python -m installer -d test_dir dist/*.whl

  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH" pytest \
    --deselect tests/webservice_test.py::TestAsyncClient
}

package() {
  cd GeoIP2-python
  python -m installer --destdir="$pkgdir" dist/*.whl
}
