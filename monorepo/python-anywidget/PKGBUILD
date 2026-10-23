# Maintainer: Jakub Klinkovský <lahwaacz at archlinux dot org>

_name=anywidget
pkgname=python-$_name
pkgver=0.11.0
pkgrel=1
pkgdesc="Reusable widgets made easy"
arch=(any)
url="https://anywidget.dev/"
license=(MIT)
depends=(
  ipython
  python
  python-comm
  python-ipywidgets
  python-psygnal
  python-traitlets
  python-typing_extensions
)
makedepends=(
  python-build
  python-installer
  python-hatchling
  python-hatch-jupyter-builder
  pnpm
)
checkdepends=(
  python-pytest
  python-pydantic
  python-watchfiles
)
optdepends=(
  'python-watchfiles: for live reloading during development'
)
source=($_name-$pkgver.tar.gz::https://github.com/manzt/anywidget/archive/refs/tags/$_name@$pkgver.tar.gz)
b2sums=('0efaef9a6fc476345b19e93ce4107be332e025346f3f63c884647a10e804bd8bd92d855134625e09dc9fac0f362474ef0873c1321ac027a7b06a3dd4f6626c42')

prepare() {
  cd $_name-$_name-$pkgver
  # install required npm packages for building the frontend
  pnpm install --frozen-lockfile
}

build() {
  cd $_name-$_name-$pkgver
  # build system expects `vp` on PATH
  PATH="$PWD/node_modules/.bin:$PATH" \
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    -W ignore::DeprecationWarning
    # skip failures due to psygnal and other comm mocking: AssertionError: assert not {140618835380272: <MockComm id='140618839307136'>}
    --deselect tests/test_descriptor.py::test_descriptor_with_psygnal
    --deselect tests/test_descriptor.py::test_descriptor_with_pydantic
    --deselect tests/test_descriptor.py::test_descriptor_with_msgspec
    --deselect tests/test_descriptor.py::test_descriptor_with_traitlets
    --deselect tests/test_descriptor.py::test_infer_file_contents
  )

  cd $_name-$_name-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest "${pytest_options[@]}" tests
}

package() {
  cd $_name-$_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  mv "$pkgdir"/usr/etc "$pkgdir"

  install -vDm 644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
  install -vDm 644 README.md -t "$pkgdir"/usr/share/doc/$pkgname/
}
