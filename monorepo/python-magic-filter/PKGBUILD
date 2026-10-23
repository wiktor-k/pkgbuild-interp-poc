# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Evgeniy Filimonov <evgfilim1 at gmail dot com>

pkgname=python-magic-filter
_pkgname=${pkgname#python-}
pkgver=1.0.12
pkgrel=5
pkgdesc="A filter package based on dynamic attribute getter"
arch=(any)
url="https://github.com/aiogram/magic-filter"
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-hatchling
  python-installer
  python-wheel
)
checkdepends=(python-pytest)
source=("$url/archive/v$pkgver/$_pkgname-$pkgver.tar.gz")
sha256sums=('e1bdaddbc2090a13883956c55bd0dd5b8affd0eeb825ce19ec5e891247f7c0cf')

build() {
  cd "$_pkgname-$pkgver"
  python -m build --wheel --no-isolation
}

check() {
  cd "$_pkgname-$pkgver"
  pytest
}

package() {
  cd "$_pkgname-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
