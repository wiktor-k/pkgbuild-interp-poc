# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-django-haystack
_name="${pkgname#python-}"
pkgver=3.4.0
pkgrel=1
pkgdesc="Modular search for Django"
arch=(any)
url="https://github.com/django-haystack/django-haystack"
license=(BSD-3-Clause)
depends=(
  python
  python-django
  python-packaging
)
makedepends=(
  python-build
  python-installer
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-dateutil
  python-elasticsearch
  python-geopy
  python-pysolr
  python-requests
  python-whoosh
)
optdepends=(
  'python-elasticsearch: interface with an elasticsearch instance as search backend'
  'python-geopy: for distance information'
  'python-pysolr: interface with a solr instance as search backend'
  'python-whoosh: use whoosh as search backend'
  'python-xapian-haystack: use python-xapian as search backend'
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('18ecda9d683f94cd4c0ca993bbddf76c04d2bc45f96d5b5fd315ff90f5d4c15a07a75c5a4c92545ab0fe7a78f8ab1d0fe911aee600e25e1f7a8bdb42b8d64310')
b2sums=('e8f38653874cd6489723d389eddc0e5ed2ffc95d0146b9adce069826adc95d2f0d293c452b6005054e7d5471ccedc74cdf69a9fec7d0eed84dac9d689f519818')

build() {
  cd $_name-$pkgver
  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver/
  python test_haystack/run_tests.py
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {AUTHORS,CONTRIBUTING.md,README.rst} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
