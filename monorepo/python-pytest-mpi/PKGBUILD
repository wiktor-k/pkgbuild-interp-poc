# Maintainer: Bruno Pagani <archange@archlinux.org>

_name=pytest-mpi
pkgname=python-$_name
pkgver=0.6
pkgrel=7
pkgdesc="pytest plugin for working with MPI"
arch=(any)
url="https://github.com/aragilar/pytest-mpi"
license=(BSD-3-Clause)
depends=(
  python
  python-mpi4py
  python-pytest
)
makedepends=(
  python-build
  python-installer
  python-setuptools
)
checkdepends=(
  python-sybil
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
b2sums=('ee71c3ee65354801de81a9af1a5b5733c96fc061af68fa34c2d88edffbe3b8b4b59c87cff3062b2780b829074dcaed85fa50f1cccf7c9d0eaee1edf9db553e60')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    -W ignore::DeprecationWarning
    -p pytester
    --runpytest=subprocess
  )

  cd $_name-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest "${pytest_options[@]}" tests
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
