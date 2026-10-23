# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-lazy-object-proxy
pkgver=1.13.0
pkgrel=1
pkgdesc='A fast and thorough lazy object proxy'
arch=('x86_64')
url="https://github.com/ionelmc/python-lazy-object-proxy"
license=('BSD-2-Clause')
depends=("python")
makedepends=('git' 'python-setuptools-scm' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-pytest-benchmark')
source=("git+https://github.com/ionelmc/python-lazy-object-proxy.git#tag=v$pkgver")
sha512sums=('7e1534015fd923777b75834cbb6c073c93c52316745dfdfb1691586b33a28e6fa2d4b7cd6037017c3d6b8e482daca74e68251137c879eeb01da4c3686341da12')

build() {
  cd python-lazy-object-proxy
  python -m build --wheel --no-isolation
}

check() {
  cd python-lazy-object-proxy
  local python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
  PYTHONPATH="$PWD/build/lib.linux-$CARCH-cpython-${python_version}" pytest
}

package() {
  cd python-lazy-object-proxy
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
