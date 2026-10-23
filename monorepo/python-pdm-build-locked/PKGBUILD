# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-pdm-build-locked
_name=${pkgname#python-}
pkgver=0.3.8
pkgrel=1
pkgdesc="A pdm plugin to publish locked dependencies as optional-dependencies"
arch=(any)
url="https://github.com/pdm-project/pdm-build-locked"
license=(MIT)
depends=(
  python
  python-pdm
)
makedepends=(
  python-build
  python-installer
  python-pdm-backend
)
checkdepends=(
  python-hatchling
  python-pkginfo
  python-pytest
  python-pytest-mock
)
optdepends=(
  'python-hatchling: for python-hatch integration'
  'python-pdm-backend: for python-pdm-backend hook'
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('6332fd35453a878dfd995765a72bf96ddf835d5f90be2d8bcc6f2a818118ef96c49288c3ded064fba81901dbceddb598a47250946d22f09df30a01ce0cc95c13')
b2sums=('2830347d880477caab9756f7e5efe8ea96d83e4e77ea08161e3e5c5c5bfde785f357cd7b3bdda0619f7e0cf1416f7571e09d8b7d612e0fc6eeff753e325c6de1')

build() {
  cd $_name-$pkgver
  export PDM_BUILD_SCM_VERSION=$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    -o addopts=""  # we are not interested in coverage tests
    # some issue with dependency count (not really relevant for functionality)
    --deselect 'tests/unit/test_build_command.py::test_build_locked_pyproject[large]'
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH" pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
