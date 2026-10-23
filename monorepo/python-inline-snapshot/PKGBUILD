# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-inline-snapshot
pkgver=0.36.0
pkgrel=1
pkgdesc="Create and update inline snapshots in your python tests"
arch=(any)
url="https://github.com/15r10nk/inline-snapshot"
license=(MIT)
depends=(
  python
  python-asttokens
  python-executing
  python-pytest
  python-rich
)
makedepends=(
  python-build
  python-hatchling
  python-installer
  python-wheel
)
checkdepends=(
  python-black
  python-dirty-equals
  python-hypothesis
  python-isort
  python-pydantic
  python-pytest-freezer
  python-pytest-mock
  python-pytest-subtests
  python-pytest-xdist
  python-time-machine
)
optdepends=(
  'python-black: retains black formatting'
  'python-dirty-equals: snapshot comparison using dirty-equals'
)
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('299132b41a274e3cfc524ac16715c793fd25c91d94685b8f816f28c3281e3fa3')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir=tmp_install dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$PWD/tmp_install/$site_packages"
  pytest \
    --deselect tests/test_formating.py::test_format_command_fail \
    --deselect tests/test_typing.py
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
