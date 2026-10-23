# Maintainer: David Runge <dvzrv@archlinux.org>

_name=click-aliases
pkgname=python-click-aliases
pkgver=1.0.7
pkgrel=1
pkgdesc="Enable aliases for click"
arch=(any)
url="https://github.com/click-contrib/click-aliases"
license=(MIT)
depends=(
  python
  python-click
)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
  python-wheel
)
checkdepends=(
  python-pytest
)
source=(git+https://github.com/click-contrib/click-aliases.git#tag=v$pkgver)
sha512sums=('dbbeacbaff9dad45e40f74ab7e5dda5e714f88c950526530b17589586a41cee3a297e65387371073605ebe238d8a3ee47e89a9e32655c5c80595e54143967db5')
b2sums=('b1ebc02d07ca497194915c734dd86d1c45cba0dd6963ffb3698d63b46ec9b44e68623c2d1e82f9271d795c730f181ae2853b9000d8139286aeb1008354229f13')

build() {
  cd $_name
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    # disable broken tests (due to undead project): https://github.com/click-contrib/click-aliases/issues/5
    --deselect tests/test_basic.py::test_invalid
    --deselect tests/test_foobar.py::test_invalid
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH"
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
