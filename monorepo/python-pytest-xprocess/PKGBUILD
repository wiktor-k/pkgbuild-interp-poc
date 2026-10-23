# Maintainer: David Runge <dvzrv@archlinux.org>
# Contributor: Hugo Osvaldo Barrera <hugo@barrera.io>
# Contributor: Neil Santos <nsantos16+aur@gmail.com>

_name=pytest-xprocess
pkgname=python-pytest-xprocess
pkgver=1.0.2
pkgrel=1
pkgdesc='Pytest plugin to manage external processes across test runs.'
arch=(any)
url="https://github.com/pytest-dev/pytest-xprocess"
license=(MIT)
depends=(
  python
  python-psutil
  python-pytest
)
makedepends=(
  python-build
  python-installer
  python-setuptools-scm
  python-wheel
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('e072995a63ac19086268471a708f3576dc588a2a4700d19b53511039f49e3fb3229eadfebba35da21f1d1dd524ea20a4cb79e57844f44d56138d08b486c16690')
b2sums=('0b7ddbe27eb81614f1af67437e76911e1c383b09af22054f503f2d3bca1be65f767617676f43a0e64dfc24fe08bda0c6c61461931317a827abf65148cca927e5')

build() {
  cd $_name-$pkgver
  SETUPTOOLS_SCM_PRETEND_VERSION="$pkgver" python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH"
  # issues with some tests at least since 0.18.1: https://github.com/pytest-dev/pytest-xprocess/issues/95
  pytest -vv -k "not test_interruption_cleanup and not test_interruption_does_not_cleanup"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 {CHANGELOG,README}.rst -t "$pkgdir/usr/share/$pkgname/"
}
