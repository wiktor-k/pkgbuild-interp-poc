# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Guillaume Horel <guillaume.horel@gmail.com>

pkgname=python-cramjam
pkgver=2.13.0
pkgrel=1
pkgdesc="Thin Python bindings to de/compression algorithms in Rust"
arch=(x86_64)
url="https://github.com/milesgranger/cramjam"
license=(MIT)
depends=(
  bzip2
  glibc
  libgcc
  python
)
makedepends=(
  cmake
  nasm
  python-build
  python-installer
  python-maturin
  python-wheel
)
checkdepends=(
  python-hypothesis
  python-numpy
  python-pytest
)
options=(!lto)
source=("$url/archive/v$pkgver/${pkgname#python-}-$pkgver.tar.gz")
sha256sums=('3cf7a15ae215b8afd3683490d28ce8f7f0c659b12dd3c1aae7e8cad3e4a7df1c')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest -vv --ignore=benchmarks/ \
    --deselect='tests/test_variants.py::test_variants_decompress_into'
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
