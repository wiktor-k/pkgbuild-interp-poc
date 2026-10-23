# Maintainer: Jelle van der Waa <jelle@vdwaa.nl>
# Contributor: Grey Christoforo <first name [at] last name [dot] net>

pkgname=python-ijson
pkgver=3.5.0
pkgrel=1
pkgdesc="Iterative JSON parser with Pythonic interface"
arch=(x86_64)
url="https://github.com/ICRAR/ijson"
license=('BSD-3-Clause' 'ISC')
depends=('python' 'yajl')
makedepends=(python-build python-installer python-setuptools python-wheel)
checkdepends=('python-pytest')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/ICRAR/ijson/archive/v${pkgver}.tar.gz")
sha256sums=('75b7936a6dd81bdf9207aab23e9d0b70cf1bff751b120502da0181da8fad18ec')

build() {
  cd ijson-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ijson-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -P -m pytest
}

package(){
  cd ijson-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -m644 -Dt "${pkgdir}/usr/share/licenses/${pkgname}" LICENSE.txt
}

# vim:ts=2:sw=2:et:
