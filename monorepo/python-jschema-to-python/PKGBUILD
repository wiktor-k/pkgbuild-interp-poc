# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=python-jschema-to-python
pkgver=1.2.3
pkgrel=9
pkgdesc='Generate source code for a set of Python classes from a JSON schema'
arch=(any)
url='https://github.com/microsoft/jschema-to-python'
license=(MIT)
depends=(
  python
  python-attrs
  python-jsonpickle
  python-pbr
)
makedepends=(
  python-build
  python-installer
  python-wheel
)
checkdepends=(python-pytest)
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('e61f38470c7c18fd5fbc4b6fc97b891416de8c04d31dccb2a79b88aae892a29aa42d91b89153ac3806a518f7cad6df6fc544362c64ba83f67683c811ee68fefc')

build() {
  cd ${pkgname#python-}-$pkgver
  export PBR_VERSION=$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  pytest
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
