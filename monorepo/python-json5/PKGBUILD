# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=python-json5
pkgver=0.17.3
pkgrel=1
pkgdesc='A Python implementation of the JSON5 data format'
arch=(any)
url='https://pypi.org/project/json5/'
license=(Apache-2.0)
depends=(python)
makedepends=(git
             python-build
             python-installer
             python-setuptools
             python-wheel)
checkdepends=(python-pytest)
source=(git+https://github.com/dpranke/pyjson5#tag=v$pkgver)
sha256sums=('eec14f8cdd2bd557cdd984c8162bcff4c0c5c4b7fb83dd95f8e94a7cd376acbd')

build() {
  cd pyjson5
  python -m build --wheel --no-isolation
}

check() {
  cd pyjson5
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest -v
}

package() {
  cd pyjson5
  python -m installer --destdir="$pkgdir" dist/*.whl
}
