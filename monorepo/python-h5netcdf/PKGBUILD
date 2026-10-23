# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-h5netcdf
pkgver=1.8.1
pkgrel=1
pkgdesc="Pythonic interface to netCDF4 via h5py"
arch=('any')
url="https://github.com/h5netcdf/h5netcdf"
license=('BSD-3-Clause')
depends=(
  'python'
  'python-h5py'
  'python-numpy'
  'python-packaging'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-setuptools-scm'
  'python-wheel'
)
checkdepends=(
  'python-netcdf4'
  'python-pytest'
)
source=(
  "$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz"
  "$pkgname-exclude-tests.patch"
)
b2sums=('05f052b1349ceb7e79ca9e4d353ab64e341428ae7d4560a5f9e1ce0e97c94a2b49555b740b80a2f77d14952b751b5b1b9846cbaa2678410091dd58b6b9149fb1'
        '6d7aeb6dd9baaba13f1683f8b1adfce12f18da8e6cf80c36c1fe0d1f35e2c19d2839fd88e3b3c79d6ce8ce727a71917c4c7627a0dd60de6d245feb6bc0d14df8')

prepare() {
  cd ${pkgname#python-}-$pkgver
  patch -Np1 < ../$pkgname-exclude-tests.patch
}

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
