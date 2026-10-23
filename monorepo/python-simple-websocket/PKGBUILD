# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-simple-websocket
pkgver=1.1.0
pkgrel=2
pkgdesc='Simple WebSocket server and client for Python'
url='https://github.com/miguelgrinberg/simple-websocket'
arch=(any)
license=(MIT)
depends=(
  python
  python-wsproto
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz"
         fix-asyncio-pytest.patch)
sha512sums=('89c2c580902af77e8708bf2ae5c40358f3f9fe3052ee58d42bc518a68042c8b0bda1ad0cf3e304b67189590e75ee1bfe085d12f9d18755c443731c0884b1cb92'
            'f8bc1cc2b0bce4aadd4532513b0a9bf8e629c10b9fad0786eb4ea1947fe771b6e058f51f15e7916ddbbaf93c147371449f6f5cdc76673ca4a0b73fa0f3a6c0fd')

prepare() {
  cd ${pkgname#python-}-$pkgver
  patch -p1 -i ../fix-asyncio-pytest.patch
}

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
