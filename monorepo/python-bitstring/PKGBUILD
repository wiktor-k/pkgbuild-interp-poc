# Maintainer: Santiago Torres-Arias <santiago@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: nblock <nblock [/at\] archlinux DOT us>
# Contributor: Thomas Conneely <tc116 at le dot ac dot uk>

pkgname=python-bitstring
pkgver=4.4.0
pkgrel=1
pkgdesc='A Python module to help you manage your bits'
arch=('any')
url="https://github.com/scott-griffiths/bitstring"
license=('MIT')
depends=(
  'python'
  'python-bitarray'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-hypothesis'
  'python-pytest'
)
source=("git+$url.git#tag=$pkgver")
b2sums=('18971e0c67bda9bf83c6a244c864c696f48057c07b26de097979ac5658477cfc95ab1c3578f597a100640afa15abe323120db9a6cf6543f9dc47d17e6bbd31f5')

build() {
  cd "${pkgname#python-}"
  python -m build --wheel --no-isolation
}

check() {
  cd "${pkgname#python-}"
  local pytest_args=(
    # Requires python-gfloat which is not yet packaged.
    # https://github.com/graphcore-research/gfloat
    --ignore=tests/test_fp8.py
    --ignore=tests/test_mxfp.py

    # Deselect to avoid dependency on python-pytest-benchmark.
    --deselect=tests/test_benchmarks.py
  )
  pytest "${pytest_args[@]}"
}

package() {
  cd "${pkgname#python-}"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
