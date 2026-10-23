# Maintainer: David Runge <dvzrv@archlinux.org>

_name=kazoo
pkgname=python-kazoo
pkgver=2.11.0
pkgrel=1
pkgdesc="A high-level Python library that makes it easier to use Apache Zookeeper"
arch=(any)
url="https://github.com/python-zk/kazoo"
license=(Apache-2.0)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
# NOTE: tests require running zookeeper
optdepends=(
  'python-eventlet: for using eventlet as networking library'
  'python-gevent: for using gevent as networking library'
  'python-pure-sasl: for SASL support'
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('4714b1ce308aaad16004feccc23fdc97eda1db491424500c5a443717e95afc9ecc6df2163c38fe254f9cf0750bc41d8891bfce41752e1bf31533b10c0533ec5f')
b2sums=('5d4d7762a920c53669eb7d2582d43835d2cfb3eb170c3e0c03271d515fbfc555935ea06025155dd2011b037a34a005970e88702d3640e133534d019e0dba5168')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGES,CONTRIBUTING,README}.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
