# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-django-guardian
_name="${pkgname#python-}"
pkgver=3.5.2
pkgrel=1
pkgdesc="Per object permissions for Django"
arch=(any)
url="https://github.com/django-guardian/django-guardian"
license=(
  BSD-2-Clause
  CC-BY-4.0
)
depends=(
  python
  python-django
  python-django-environ
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-django
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('09a802200e0cfd2298f175edefcd3223d6d218cb9ca43d73f569a65d74942fca8986015c0e156ae9df2ca150d887a35e3fd8ff4929bb3773a7c73a42e5f9dffc')
b2sums=('25b58bcf9836c5c2f9153b30bf4e36938b7571b30ecfa48329ca9d976462087785db564fba7cbf6200d9a840a1b337bdd1a8378c49b59ae213e8efdf66915222')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  export PYTHONPATH="build:$PYTHONPATH"
  pytest -v
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {README,SECURITY}.md -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
