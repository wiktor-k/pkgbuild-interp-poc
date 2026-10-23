# Maintainer: David Runge <dvzrv@archlinux.org>

_name=pysolr
pkgname=python-pysolr
pkgver=3.11.0
pkgrel=2
pkgdesc="Python Solr client"
arch=(any)
url="https://github.com/django-haystack/pysolr"
license=(BSD-3-Clause)
depends=(
  python
  python-setuptools
  python-requests
)
makedepends=(
  python-build
  python-installer
  python-setuptools-scm
  python-wheel
)
optdepends=(
  'python-kazoo: for solrcloud'
  'python-simplejson: for JSON support'
  'solr: for interfacing a local solr instance'
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('09471b3d0b8b0211b66a94e4d46794ef48eea0fb904254a39755be15b97d00fa97a0ae61df335446c9cccf2b855092c572da8160f7ad433906205f899fa3b649')
b2sums=('d18f64e1188847b7e832a1e914487aa02164f3322f398f7e0ba41f7ae10b425779fdf3ef654686a10bb18052bb4db384e43f553d4addf892dfd53188f886d9d8')

build() {
  cd $_name-$pkgver
  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --no-isolation
}

# NOTE: tests vendor solr *and* zookeeper (including download)

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {AUTHORS,{CHANGELOG,README}.rst} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
