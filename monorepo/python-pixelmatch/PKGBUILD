# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Carlos Aznarán <caznaranl@uni.pe>

_pyname=pixelmatch-py
pkgname=python-pixelmatch
pkgdesc='A pixel-level image comparison library'
pkgver=0.4.0
pkgrel=1
arch=(any)
url='https://github.com/whtsky/pixelmatch-py'
license=(ISC)
depends=(python
         python-pillow)
makedepends=(git
             python-build
             python-installer
             python-poetry-core)
checkdepends=(libpng
              python-pytest
              python-pytest-benchmark)
source=(git+https://github.com/whtsky/pixelmatch-py#tag=v$pkgver)
sha256sums=('8a6d5a2a3687f150c5d095a03fc1cf997cbff9cecd1d553eceb021d7338b254d')

build() {
  cd $_pyname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pyname
#  Test data is broken
#  pytest -v
}

package() {
  cd $_pyname
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
