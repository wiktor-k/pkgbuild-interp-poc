# Maintainer: Antonio Rojas <arojas@archlinux.org>

_pyname=pycrdt-websocket
pkgname=python-$_pyname
pkgver=0.16.5
pkgrel=1
pkgdesc='CRDTs based on Yrs'
arch=(any)
url='https://github.com/y-crdt/pycrdt-websocket'
license=(MIT)
depends=(python
         python-anyio
         python-pycrdt
         python-pycrdt-store)
makedepends=(git
             python-build
             python-hatchling
             python-installer)
checkdepends=(hypercorn
              python-httpx-ws
              python-pytest
              python-pytest-asyncio
              python-trio
              python-websockets
              uvicorn)
source=(git+https://github.com/y-crdt/pycrdt-websocket#tag=$pkgver)
sha256sums=('3f38372fe81a37a0a6a3db669cc158fccdcc21c51e690b6c51f59b6723d384a2')

build() {
  cd $_pyname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pyname
  python -m venv --system-site-packages test-env
# hack to allow importing pycrdt/store from the venv
  cp -r /usr/lib/python*/site-packages/pycrdt/ test-env/lib/python*/site-packages
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest -v -k 'not test_pycrdt_yjs.py'
}

package() {
  cd $_pyname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
