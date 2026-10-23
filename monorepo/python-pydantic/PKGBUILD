# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Filipe Laíns (FFY00) <lains@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>

_name=pydantic
pkgname=python-$_name
# WARNING: upstream pins pydantic-core down to the patch-level and using other versions breaks tests! only update in lock-step with python-pydantic-core!
pkgver=2.14.1
pkgrel=1
pkgdesc='Data parsing and validation using Python type hints'
arch=(any)
url="https://github.com/pydantic/pydantic"
license=(MIT)
depends=(
  python
  python-annotated-types
  python-pydantic-core
  python-typing_extensions
  python-typing-inspection
)
makedepends=(
  cython
  python-build
  python-installer
  python-hatchling
  python-hatch-fancy-pypi-readme
  python-wheel
)
checkdepends=(
  python-ansi2html
  python-cloudpickle
  python-dirty-equals
  python-email-validator
  python-faker
  python-greenlet
  python-hypothesis
  python-inline-snapshot
  python-jsonschema
  python-packaging
  python-pygments
  python-pytest
  python-pytest-benchmark
  python-pytest-examples
  python-pytest-mock
  python-pytest-pretty
  python-pytest-run-parallel
  python-pytest-timeout
  python-pytz
  python-sqlalchemy
  python-time-machine
)
optdepends=(
  'mypy: for type validation with mypy'
  'python-dotenv: for .env file support'
  'python-email-validator: for email validation'
  'python-hypothesis: for hypothesis plugin when using legacy v1'
)
source=(
  $url/archive/v$pkgver/$_name-v$pkgver.tar.gz
)
sha512sums=('70d4a6513f49e4644c07acc7eba2fde0e05724db8ac8fb9777414ae65dbb30f796db94d5843e869d7a1655eb0f2c09f230454a394c095b97fcf5088795cc8493')
b2sums=('7abb504992e3ec1b385eec1777dc6df7486bc004776e394cfa7a584a5e949440d621361928da6776ca676484e30a2f606958c97f834e37cd4788a950c34d8656')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    --deselect tests/test_docs.py  # we are not interested in code formatting
    --deselect tests/pydantic_core/test_docstrings.py::test_readme  # we are not interested in code formatting
    --deselect tests/pydantic_core/test_docstrings.py::test_docstrings  # we are not interested in code formatting
    --deselect tests/test_types.py::test_string_fails
    --deselect 'tests/test_networks.py::test_address_invalid[foobar-An email address must have an @-sign.]'
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH"
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
