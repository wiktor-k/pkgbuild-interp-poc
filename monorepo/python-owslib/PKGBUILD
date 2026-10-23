# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Doug Newgard <scimmia at archlinux dot info>

pkgname=python-owslib
pkgver=0.36.0
pkgrel=1
pkgdesc='Python package for client programming with Open Geospatial Consortium (OGC) web service interface standards, and their related content models'
arch=(any)
url='https://geopython.github.io/OWSLib/'
license=(BSD)
depends=(python-dateutil python-lxml python-requests python-yaml)
makedepends=(python-build python-installer python-setuptools python-wheel)
source=(https://github.com/geopython/OWSLib/archive/$pkgver/$pkgname-$pkgver.tar.gz)
sha256sums=('1aa07c28742dafbb8da070e8f6b2fbd5d2181ae86eb7dadf0bde2dae1b8c7515')

prepare() {
  cd OWSLib-$pkgver
  sed -e 's|setuptools<69|setuptools|' -i pyproject.toml
}

build() {
  cd OWSLib-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd OWSLib-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
