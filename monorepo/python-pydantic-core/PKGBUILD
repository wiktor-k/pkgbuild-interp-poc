# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>

_name=pydantic-core
pkgname=python-pydantic-core
# WARNING: this package is pinned down to the patch-level version in python-pydantic and should only be updated in lock-step with it
pkgver=2.50.1
pkgrel=1
epoch=3
pkgdesc="Core validation logic for pydantic written in rust "
arch=(x86_64)
url="https://github.com/pydantic/pydantic-core"
license=(MIT)
depends=(
  glibc
  libgcc
  python
  python-typing-inspection
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-maturin
)
checkdepends=(
  python-dirty-equals
  python-hypothesis
  python-inline-snapshot
  python-pytest
  python-pytest-benchmark
  python-pytest-examples
  python-pytest-mock
  python-pytest-run-parallel
  python-pytest-timeout
)
options=(!lto)
source=(
    "pydantic::git+https://github.com/pydantic/pydantic.git#tag=core-v${pkgver}"
)

sha512sums=('7bd832af38a5b98e0863204887f25694b9354eb25d0132553e12a35e4eefa64228c7d6518c2c1dd45d15942291c607c2e1993cb7aff475de1982a7bf41779039')
b2sums=('7f325b25dc69f135f7c54f15facbff4693b1174117dff07cb50fe193ac7503219c1ef7819cb0a473b7002d0133e8f4b54e116ad803e67f4b1420e84bd5d5dd14')

build() {
  cd "pydantic/$_name"
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    --ignore tests/test_docstrings.py  # we are not interested in linting/ formatting with ruff
    --ignore tests/test_tzinfo.py  # requires python-pydantic
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd "pydantic/$_name"
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH"
  HYPOTHESIS_PROFILE=slow pytest "${pytest_options[@]}"
}

package() {
  cd "pydantic/$_name"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
