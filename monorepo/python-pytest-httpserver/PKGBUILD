# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Daniel Bermond <dbermond@archlinux.org>

_name=pytest-httpserver
pkgname=python-pytest-httpserver
pkgver=1.2.0
pkgrel=1
pkgdesc="Http server for pytest to test http clients"
arch=(any)
url="https://github.com/csernazs/pytest-httpserver"
license=(MIT)
depends=(
  python
  python-pytest
  python-werkzeug
)
makedepends=(
  python-build
  python-installer
  python-poetry-core
  python-wheel
)
checkdepends=(
  python-pytest
  python-requests
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha256sums=('d2758eafa784baca1f15cd370c40aeeaa721b43e5f11346bf61c6457460aaf01')
b2sums=('4de517ecf479f909d46e523a4e0fb92e7bad147f30b4e44b91e7833a759238e5b7a8e1f80641b88ccc96a2d0fd55b82941bcfbaa8c5d3ab2aa9d790a35c2ac0d')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH"
  pytest -vv --ignore tests/test_release.py
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 CHANGES.rst README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
