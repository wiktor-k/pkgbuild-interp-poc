# Maintainer: Bert Peters <bertptrs@archlinux.org>

pkgname=python-narwhals
_name=${pkgname#python-}
pkgver=2.27.1
pkgrel=1
pkgdesc="Extremely lightweight and extensible compatibility layer between dataframe libraries"
arch=('any')
url="https://narwhals-dev.github.io/narwhals/"
license=('MIT')
depends=(python)
makedepends=(python-build python-installer python-wheel python-uv-build)
# Not included, as it's not currently packaged for [extra]: duckdb
optdepends=('python-pandas: for Pandas support'
            'python-pyarrow: for PyArrows support'
            'python-polars: for Polars support')
checkdepends=(
	python-pytest{,-env}
	python-hypothesis
	python-pandas
	python-pyarrow
	python-polars
)
source=($pkgname-$pkgver.tar.gz::https://github.com/narwhals-dev/narwhals/archive/refs/tags/v$pkgver.tar.gz)
sha256sums=('56d4165b8f7204c52ff073869202145df562a6a16246ebe4222154d945713d75')

build() {
  cd "$_name-$pkgver"
  # Upstream requires an older version of uvbuild but it's not necessary
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd "$_name-$pkgver"
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # Not all supported data frames are packaged, only run tests for those
  # that are
  test-env/bin/python -P -m pytest --constructors "pandas,pandas[pyarrow],polars[eager],pyarrow"
}

package() {
  cd "$_name-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE.md
}
