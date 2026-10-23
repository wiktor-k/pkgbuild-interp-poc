# Maintainer: Daniel M. Capella <polyzen@archlinux.org>

pkgname=python-aiohttp-openmetrics
_name=${pkgname#python-}
pkgver=0.0.13
pkgrel=1
pkgdesc='OpenMetrics provider for aiohttp'
arch=(any)
url=https://github.com/jelmer/aiohttp-openmetrics
license=(Apache)
depends=(
  python-aiohttp
  python-prometheus_client
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('068ed80871ad0fc7629167573d2bd728341956d9c07e033b62df105d8f2165afeb16917c5302616c7bb1bbbd4197fbda046752ea4b0a297bbbafa63ced34fd8b')
#validpgpkeys=(DC837EE14A7E37347E87061700806F2BD729A457) # Jelmer Vernooĳ <jelmer@jelmer.uk>

build() {
  cd "$_name"
  python -m build --wheel --skip-dependency-check --no-isolation
}

package() {
  cd "$_name"
  python -m installer --destdir="$pkgdir" dist/*.whl
}
