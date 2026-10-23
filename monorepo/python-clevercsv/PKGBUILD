# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-clevercsv
pkgver=0.8.5
pkgrel=1
pkgdesc="A Python package for handling messy CSV files"
url="https://github.com/alan-turing-institute/CleverCSV"
license=('MIT')
arch=('x86_64')
depends=('python-chardet' 'python-regex' 'python-packaging')
makedepends=('git' 'python-setuptools' 'python-build' 'python-installer' 'python-wheel')
optdepends=('python-pandas' 'python-tabview' 'python-wilderness' 'python-faust-cchardet')
checkdepends=('python-pytest' 'python-pandas' 'python-tabview' 'python-wilderness'
              'python-faust-cchardet' 'python-termcolor')
source=("git+https://github.com/alan-turing-institute/CleverCSV.git#tag=v$pkgver")
sha512sums=('fdd25aefd2b624cf5deba04fc10a14528e2fbc41a7552a29083d0b0a15517b84cb7a41b4cd42f4185076eed48c5f0bc16a7de2705f0f5bd319d861613130aa8f')

build() {
  cd CleverCSV
  python -m build --wheel --no-isolation
}

check() {
  cd CleverCSV
  local python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
  # chardet 6 can still return None for tiny Latin-1 samples after upstream's relaxed expectations.
  PYTHONPATH="build/lib.linux-x86_64-cpython-${python_version}" pytest \
    --deselect tests/test_unit/test_console.py::ConsoleTestCase::test_standardize_target_encoding2 \
    --deselect tests/test_unit/test_encoding.py::EncodingTestCase::test_encoding_chardet
}

package() {
  cd CleverCSV
  python -m installer --destdir="$pkgdir" dist/*.whl
  mkdir "$pkgdir"/usr/share
  mv "$pkgdir"/usr/man "$pkgdir"/usr/share/man 

  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
