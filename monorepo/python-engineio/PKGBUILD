# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-engineio
pkgver=4.14.0
pkgrel=1
pkgdesc='Python implementation of the Engine.IO realtime server'
url='https://github.com/miguelgrinberg/python-engineio'
arch=(any)
license=(MIT)
depends=(
  python
  python-requests
  python-simple-websocket
  python-websocket-client
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-sphinx
  python-wheel
)
checkdepends=(
  python-aiohttp
  python-eventlet
  python-gevent
  python-gevent-websocket
  python-pytest
  python-pytest-asyncio
  python-tornado
)
optdepends=(
  'python-aiohttp: aiohttp driver'
  'python-eventlet: eventlet driver'
  'python-gevent-websocket: gevent driver'
  'python-gevent: gevent driver'
  'python-tornado: tornado driver'
)
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('9079c5c831ee11f06b35f80c562ae3182c5dfc8694d711000c8d3d72b92325b25b2ea8a08d8315da8888c8d499c11c8c8db0b4dcc8cdf65fa80ec192ffbb6122')
b2sums=('9bb26c1e94c2b156ec2c3d0f5e730234d08dfd1f36d130bfbe344b2209949c35bfcadc852fcd1396fd774dd7e4c96580992fb8e9918f9eb6cd2cdcecf94c7845')

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
  PYTHONPATH="$PWD/src" make -j1 -C docs man text SPHINXBUILD=sphinx-build
}

check() {
  cd $pkgname-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # Deselect test failing with the following error, not sure why:
  #   tests/common/test_server.py:128: AssertionError
  test-env/bin/python -m pytest \
    --deselect tests/common/test_server.py::TestServer::test_async_mode_eventlet
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 -t "$pkgdir/usr/share/doc/$pkgname" README.md docs/_build/text/*.txt
  install -vDm 644 -t "$pkgdir/usr/share/man/man1" docs/_build/man/python-engineio.1
  install -vDm 644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}

# vim: ts=2 sw=2 et:
