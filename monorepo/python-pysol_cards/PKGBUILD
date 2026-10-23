# Maintainer: Antonio Rojas <arojas@archlinux.org>

_pyname=pysol_cards
pkgname=python-$_pyname
pkgver=0.24.0
pkgrel=2
pkgdesc='Deal PySol FC Cards'
arch=(any)
url='https://pypi.org/project/pysol-cards/'
license=(MIT)
depends=(python)
makedepends=(python-build
             python-installer
             python-setuptools
             python-wheel)
source=(https://pypi.io/packages/source/p/$_pyname/$_pyname-$pkgver.tar.gz)
sha256sums=('a985492da81aa1588dfc056d9a7c6ca83f66254c0b90f25afc682a70713d4d4b')

build() {
  cd $_pyname-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $_pyname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -D -m644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
