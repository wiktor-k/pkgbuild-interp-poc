# Maintainer: Santiago Torres-Arias <santiago@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Jonas Witschel <diabonas@archlinux.org>

pkgname=python-aiohttp-socks
pkgver=0.12.0
pkgrel=1
pkgdesc="SOCKS proxy connector for aiohttp. SOCKS4(a) and SOCKS5 are supported."
arch=('any')
url="https://github.com/romis2012/aiohttp-socks"
license=('Apache-2.0')
depends=(
  'python'
  'python-aiohttp'
  'python-python-socks'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-anyio'
  'python-flask'
  'python-pytest'
  'python-pytest-asyncio'
  'python-starlette'
  'python-tiny-proxy'
  'python-trustme'
  'uvicorn'
)
source=("git+$url.git?signed#tag=v$pkgver")
b2sums=('c1dc9d91dfa94ee06a32bee5f401555f74b04bdac73ace68b155f0563f61b6627bc47cf1e4ab05d564736cbf29f74747bb29d582c97eb67baa3e996d2fbbbbf7')
validpgpkeys=('706A8935FED55F3BFB7C188B5EEAD7400A5D80ED') # Roman Snegirev (maintainer)

build() {
  cd "${pkgname//python-/}"
  python -m build --wheel --no-isolation
}

check() {
  cd "${pkgname//python-/}"
  pytest
}

package() {
  cd "${pkgname//python-/}"
  python -m installer --destdir="$pkgdir" dist/*.whl
}
