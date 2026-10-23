# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Hugo Osvaldo Barrera <hugo@barrera.io>

pkgname=python-aiostream
_name=${pkgname#python-}
pkgver=0.8.1
pkgrel=1
pkgdesc='Generator-based operators for asynchronous iteration'
arch=(any)
url=https://github.com/vxgmichel/aiostream
license=(GPL-3.0-only)
depends=(python-typing_extensions)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-asyncio
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('6931db01b7581e7ab1f956f824cc5ed5cc8d0daaaffcc0833de4a3c283e1d15c86371a6b84c2966c85efa8fe8d394b2e5d19e4f31365ff68abe394c7dc9e7885')

build() {
  cd "$_name"
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd "$_name"
  # Override addopts as they invoke coverage testing
  PYTHONPATH="$PWD/$_name:$PYTHONPATH" pytest --override-ini="addopts=" tests --strict-markers
}

package() {
  cd "$_name"
  python -m installer --destdir="$pkgdir" dist/*.whl
}
