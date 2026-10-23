# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-blockbuster
pkgver=1.5.28
pkgrel=1
pkgdesc='Utility to detect blocking calls in the async event loop'
arch=(any)
url=https://github.com/cbornet/blockbuster
license=(Apache-2.0)
depends=(
  python
  python-forbiddenfruit
)
makedepends=(
  git
  python-build
  python-hatchling
  python-installer
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-asyncio
  python-pytest-benchmark
  python-requests
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('ed688db8a24276675d871cf992051d490b768f4eb038296e017b2ade68d8abfe652a53bc1292fed5339f263143580aa0b3c8dca40610145210f2beebf823a09d')

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  pytest
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
}
