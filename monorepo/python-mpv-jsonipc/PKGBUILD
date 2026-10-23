# Maintainer: Giovanni Harting <anonfunc@archlinux.org>

pkgname=python-mpv-jsonipc
pkgver=1.4.0
pkgrel=1
pkgdesc='Python API to MPV using JSON IPC'
arch=(any)
url='https://github.com/iwalton3/python-mpv-jsonipc'
license=(Apache-2.0)
depends=(python)
makedepends=(python-build python-installer python-wheel python-setuptools)
source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz")
b2sums=('0173b6a68fcb19cc18afda6cd1121cbeaa41a5326f824d84505011c1daf1fdf83fa25550a959712de8f7c7ee7214de26edae3eb285b2932b2e4fbd8ebb80f0cd')

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim:set ts=2 sw=2 et:
