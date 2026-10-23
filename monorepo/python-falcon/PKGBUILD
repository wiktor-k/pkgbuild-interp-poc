# Maintainer: David Runge <dvzrv@archlinux.org>

_name=falcon
pkgname=python-falcon
pkgver=4.4.0
pkgrel=1
pkgdesc="Web API framework for fast and reliable microservices, proxies, and app backends"
arch=(x86_64)
url="https://falconframework.org/"
_url="https://github.com/falconry/falcon"
license=(Apache-2.0)
depends=(
  glibc
  python
)
makedepends=(
  cython
  python-build
  python-installer
  python-setuptools
  python-wheel
)
# TODO: add python-daphne
checkdepends=(
  gunicorn
  hypercorn
  python-aiofiles
  python-cbor2
  python-httpx
  python-jsonschema
  python-mimeparse
  python-msgpack
  python-mujson
  python-orjson
  python-pecan
  python-pytest
  python-pytest-asyncio
  python-pyyaml
  python-rapidjson
  python-requests
  python-testtools
  python-ujson
  python-websockets
  uvicorn
)
optdepends=(
  'python-bottle: for falcon-bench'
  'python-django: for falcon-bench'
  'python-flask: for falcon-bench'
  'python-jsonschema: for JSON schema validation'
  'python-msgpack: for MessagePackHandler'
  'python-pecan: for falcon-bench'
  'python-pprofile: for falcon-bench'
)
source=($_name-$pkgver.tar.gz::$_url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('899aaaa83de502be505022f256426536399a80786e0a147a3c1ca9f50882e073831494af5db341827535947810ba0ec7e3467eddda50f981b8d1a886ca8b6904')
b2sums=('c2271326082a6e9c89969a5a4f7d62c874bde8e24344142853351cae8cb14e0c33441e600f165108f62c279265661a9ba518d007a7df5deae068eb645c056e75')

prepare() {
  cd $_name-$pkgver
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  python -m installer --destdir=test_dir dist/*.whl
  PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH" pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
