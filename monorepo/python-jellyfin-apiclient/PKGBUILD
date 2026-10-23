# Maintainer: Giovanni Harting <anonfunc@archlinux.org>
# Contributor: Marius Lindvall <(firstname) {cat} varden {dog} info>

pkgname=python-jellyfin-apiclient
pkgver=1.20.1
pkgrel=1
pkgdesc='Python API client for Jellyfin'
arch=(any)
url=https://github.com/iwalton3/jellyfin-apiclient-python
license=(GPL-3.0-only)
depends=(python python-requests python-urllib3 python-websocket-client python-certifi)
makedepends=(python-build python-installer python-wheel python-setuptools)
source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz")
b2sums=('298929d61b42c57ff18c05d7e75a7450dc7c5218ce0b631cf8d168c552a84fff88a7d71ab9b50efc849148c0194c2974e5cca5d2e31a068f27f1f06546275f17')

build() {
	cd jellyfin-apiclient-python-$pkgver
	python -m build --wheel --no-isolation
}

package() {
	cd jellyfin-apiclient-python-$pkgver
	python -m installer --destdir="$pkgdir" dist/*.whl
}
