# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-pprofile
_name="${pkgname#python-}"
pkgver=2.2.0
pkgrel=3
pkgdesc="Line-granularity, thread-aware deterministic and statistic pure-python profiler"
arch=(any)
url="https://github.com/vpelletier/pprofile"
license=(GPL-2.0-or-later)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=(
  $pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz
)
sha512sums=('fc53187876f9aa739afcb069ca21a59e9dbdcee20db74fd6a059916a0916f1ef1b79a50a6e9d6378d704edc974f7461a0fbe9c29fc4c5a85981a0470bf50c761')
b2sums=('da74b43239b6a99b18be9ca0b3db9d90c2021be521149eaedd72d083da9a4575efccea34e36571d36a619b59335bd1ac8654d453d564f2391f185b418a1d4f9c')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
