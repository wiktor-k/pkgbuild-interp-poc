# Maintainer: Balló György <ballogyor+arch at gmail dot com>

pkgname=python-generic
pkgver=1.1.7
pkgrel=1
pkgdesc='Generic programming library for Python'
arch=(any)
url='https://github.com/gaphor/generic'
license=(BSD-3-Clause)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
  python-sphinx-furo
  python-wheel
)
checkdepends=(python-pytest)
source=("git+https://github.com/gaphor/generic.git#tag=$pkgver")
b2sums=('397856d70d52d3ff163a7faba0dafa9bb3426ee5bb8211a4f7c328ae6e1a0e1c5dbd9d85786541a4fb2e46b3a781b52fa857fc9956255c5162312cd2ad97af1b')

prepare() {
  cd generic
}

build() {
  cd generic
  python -m build --wheel --no-isolation
}

check() {
  cd generic
  PYTHONPATH="$PWD/dist/generic-$pkgver-py3-none-any.whl" \
    pytest
}

package() {
  cd generic
  python -m installer --destdir="$pkgdir" dist/*.whl

  PYTHONPATH="$PWD/dist/generic-$pkgver-py3-none-any.whl" \
    python -m sphinx -b html docs "$pkgdir/usr/share/doc/$pkgname/html"
  rm -r "$pkgdir/usr/share/doc/$pkgname/html/.doctrees"
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSES/BSD-3-Clause.txt
}
