# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Guillaume Horel <guillaume.horel@gmail.com>

pkgname=python-fastparquet
pkgver=2026.9.0
pkgrel=1
pkgdesc="A Python implementation of the Parquet columnar file format"
arch=(x86_64)
url="https://github.com/dask/fastparquet"
license=(Apache-2.0)
depends=(
  glibc
  python
  python-cramjam
  python-fsspec
  python-numpy
  python-packaging
  python-pandas
)
makedepends=(
  cython
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(python-pytest)
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('34eb6bb7fc9d7ea94640a9845d8f59b1d09567261e3895e130e9d30e828e00d53498d0bf8fe520431fb18c5acc87f4cc864cd888e7822cdb7acff513967457fc')

build() {
  cd ${pkgname#python-}-$pkgver
  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir=tmp_install dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  cp -r fastparquet/test test-data "$PWD/tmp_install/$site_packages"
  (cd "$PWD/tmp_install/$site_packages" && pytest)
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
