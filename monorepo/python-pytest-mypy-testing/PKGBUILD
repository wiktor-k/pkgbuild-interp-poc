# Maintainer:

_pyname=pytest-mypy-testing
pkgname=python-$_pyname
pkgver=0.2.0
pkgrel=1
pkgdesc='Plugin to test mypy output with pytest'
arch=(any)
license=(Apache
         MIT)
url='https://github.com/davidfritzsche/pytest-mypy-testing'
depends=(mypy
         python
         python-pytest)
makedepends=(git
             python-build
             python-flit-core
             python-installer)
checkdepends=(python-pytest)
source=(git+https://github.com/davidfritzsche/pytest-mypy-testing#tag=v$pkgver)
sha256sums=('bad69defa3ffe4d99a8cf748af7e86d738fc9be352ea1892e5e6c98664329ba7')

build() {
  cd $_pyname
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd $_pyname
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest -v
}

package() {
  cd $_pyname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSES/* -t "$pkgdir"/usr/share/licenses/$pkgname
}
