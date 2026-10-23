# Maintainer: David Runge <dvzrv@archlinux.org>

_name=django-filter
pkgname=python-django-filter
pkgver=26.2
pkgrel=1
pkgdesc="A generic system for filtering Django QuerySets based on user selections"
arch=(any)
url="https://github.com/carltongibson/django-filter"
license=(BSD-3-Clause)
depends=(
  python
  python-django
)
makedepends=(
  python-build
  python-flit-core
  python-installer
  python-wheel
)
checkdepends=(
  python-django-crispy-forms
  python-django-rest-framework
  python-pytest
)
optdepends=(
  'python-django-crispy-forms: for enhanced filter form presentation'
  'python-django-rest-framework: for filter integration using DRF'
)
source=(
  $_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz
)
sha512sums=('520530d0c6ffb39e47b8b5aa544a581858fbce9d0ad17d45cfe9dd471c5ba26e5b36e1e16cee02937a8857546dc0d0167b8d77a3d3a5ed0c13595e4db91f4d17')
b2sums=('14acf826c18c955c3e77f139683faca47dad4ef9678ecaa689c58251081620ac847dc14826f4886de03a1f2e82c42243fffd2447f48cc4ced372d7677de3c1d1')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  python runtests.py
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {AUTHORS,{CHANGES,README}.rst} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
