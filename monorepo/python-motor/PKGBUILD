# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Butui Hu <hot123tea123@gmail.com>

pkgname=python-motor
_pkgname=${pkgname#python-}
pkgver=3.7.1
pkgrel=2
pkgdesc="Async Python driver for MongoDB and Tornado or asyncio"
arch=(any)
url="https://github.com/mongodb/motor"
license=(Apache-2.0)
depends=(
  python
  python-pymongo
)
makedepends=(
  python-build
  python-hatch-requirements-txt
  python-hatchling
  python-installer
  python-wheel
)
source=("$pkgname-$pkgver.tar.gz::$url/archive/$pkgver.tar.gz")
sha256sums=('8a1b89a73a55be08c59b6bf1bbeb4578f134cda881d1926898a0171799620404')

build() {
  cd "$_pkgname-$pkgver"
  python -m build --wheel --no-isolation
}

package() {
  cd "$_pkgname-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
}
