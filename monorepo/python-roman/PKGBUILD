# Maintainer: Jelle van der Waa <jelle@archlinux.org>

pkgname=python-roman
pkgver=5.2
pkgrel=2
pkgdesc="Small helper library to convert arabic to roman numerals"
arch=('any')
license=('ZPL-2.1')
url="https://github.com/zopefoundation/roman"
depends=(python python-pytest)
makedepends=(python-build python-installer python-wheel python-setuptools git)
source=(git+https://github.com/zopefoundation/roman#tag=${pkgver})
validpgpkeys=("E499C79F53C96A54E572FEE1C06086337C50773E")
sha512sums=('SKIP')

build() {
  cd roman

  python -m build --wheel --no-isolation
}

check() {
  cd roman
  PYTHONPATH=src python src/tests.py
}

package() {
  cd roman

  python -m installer --destdir="$pkgdir" dist/*.whl

  install -Dm 644 LICENSE.txt "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
