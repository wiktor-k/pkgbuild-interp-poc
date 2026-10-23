# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Jonathan Liu <net147@gmail.com>

pkgname=python-pyicu
pkgver=2.16.2
pkgrel=1
pkgdesc="Python binding for ICU"
arch=('x86_64')
license=('MIT')
url="http://pyicu.osafoundation.org/"
depends=('python' 'icu')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-fonttools' 'python-pytest' 'python-six')
source=("https://pypi.io/packages/source/p/pyicu/pyicu-$pkgver.tar.gz")
sha512sums=('779331e0d7ec1003c9b2754631f619c523d8c5e86e3f47b7ce5bcf5d6a8308d07a97f912fe3872b802e6031afdf184ef19969dbc80f96dc57f5b1f3849cbf0fe')

build() {
  cd pyicu-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd pyicu-$pkgver
  local python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
  PYTHONPATH="$PWD/build/lib.linux-$CARCH-cpython-$python_version" pytest
}

package() {
  cd pyicu-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -D -m644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}
