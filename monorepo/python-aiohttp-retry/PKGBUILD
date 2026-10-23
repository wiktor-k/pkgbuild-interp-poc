# Maintainer: Balló György <ballogyor+arch at gmail dot com>

pkgname=python-aiohttp-retry
pkgver=2.9.1
pkgrel=4
pkgdesc='Simple retry client for aiohttp'
arch=(any)
url='https://github.com/inyutin/aiohttp_retry'
license=(MIT)
depends=(
  python
  python-aiohttp
  python-yarl
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-aiohttp
  python-pytest-asyncio
)
source=("git+https://github.com/inyutin/aiohttp_retry.git#tag=v$pkgver")
b2sums=(24efc26daf47ba318a078813233c6d6357eee7684458b97a70b0ec87dcbed65779d2d2e4e7b513661dd572b6d3b3edecff7d5945b829ee766559455da8dd8fa0)

build() {
  cd aiohttp_retry
  python -m build --wheel --no-isolation
}

check() {
  cd aiohttp_retry
  pytest
}

package() {
  cd aiohttp_retry
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
}
