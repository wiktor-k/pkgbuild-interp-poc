# Maintainer: David Runge <dvzrv@archlinux.org>

_name=truststore
pkgname=python-truststore
pkgver=0.10.5
pkgrel=1
pkgdesc="Verify certificates using OS trust stores"
arch=(any)
url="https://github.com/sethmlarson/truststore"
license=(MIT)
depends=(
  python
)
makedepends=(
  python-build
  python-flit-core
  python-installer
  python-wheel
)
checkdepends=(
  python-aiohttp
  python-flaky
  python-httpx
  python-pyopenssl
  python-pytest
  python-pytest-asyncio
  python-pytest-httpserver
  python-requests
  python-trustme
  python-urllib3
)
optdepends=('python-urllib3: when injecting into SSLContext via urllib3')
source=(
  $pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz
  $pkgname-0.10.5-update-flit-core.patch
)
sha512sums=('b618a1d0b5ef9ed1359658054d5a5739602858d48cbac331287148e98f164da686a79d1d3d818be82fafdf14319763b7eb76f37dec0afc6eba069a40615153be'
            'eaf62b0f407122fd78f1e0dbfcd0601b65fe2d46ea52f77e1dacfd2b58cfe2c1a9a755033786474d31b102c19fde2ae6bec383f06534af95b22ff4a9263990b0')
b2sums=('50884bc56754d9d91846829ef4d09d2905a33a6f764ac1e224efdbe8e8a9cf93e587f631c7c0c79a357da0cac28875df7e1a32700a5a192ea103928bd19061a2'
        '5bbcdf3aa1b53fcc98e937acbd7d47282a9e59c8ed90c09d091d9b8f4b33e80e0196f2598601e4f5ba2868a194b29d0ba62b187856d39380e8c20928a974f25d')

prepare() {
  # update build dependency on flit-core: https://github.com/sethmlarson/truststore/pull/249
  patch -Np1 -d $_name-$pkgver -i ../$pkgname-0.10.5-update-flit-core.patch
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    # disable failing test: https://github.com/sethmlarson/truststore/issues/124
    --deselect tests/test_api.py::test_trustme_cert_loaded_via_capath
    # do not run tests that require network access
    # NOTE: not available in 0.8.0 yet
    # -m 'not internet'
  )
  cd $_name-$pkgver
  PYTHONPATH=src pytest -vv
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
