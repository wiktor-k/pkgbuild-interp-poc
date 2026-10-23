# Maintainer: Antonio Rojas <arojas@archlinux.org>

_pyname=async-lru
pkgname=python-$_pyname
pkgver=2.4.0
pkgrel=1
pkgdesc='Simple LRU cache for asyncio'
arch=(any)
url='https://github.com/aio-libs/async-lru'
license=(MIT)
depends=(python)
makedepends=(git
             python-build
             python-installer
             python-setuptools
             python-wheel)
checkdepends=(python-pytest-asyncio
              python-pytest-cov
              python-pytest-timeout)
source=(git+https://github.com/aio-libs/async-lru#tag=v$pkgver)
sha256sums=('4fe430e3aa26cd01690dc6c91518462b133d0d7541837cf04cb9b50e32182f66')

build() {
  cd $_pyname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pyname
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest -v -W ignore::DeprecationWarning
}

package() {
  cd $_pyname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
