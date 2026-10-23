# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Jonas Witschel <diabonas@archlinux.org>

pkgname=python-python-socks
pkgver=3.1.1
pkgrel=1
pkgdesc='Core proxy client (SOCKS4, SOCKS5, HTTP) functionality for Python'
arch=('any')
url='https://github.com/romis2012/python-socks'
license=('Apache-2.0')
depends=('python')
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-anyio'
  'python-curio'
  'python-flask'
  'python-pytest'
  'python-pytest-asyncio'
  'python-pytest-trio'
  'python-starlette'
  'python-tiny-proxy'
  'python-trio'
  'python-trustme'
  'python-yarl'
  'uvicorn'
)
optdepends=(
  'python-anyio: anyio support'
  'python-curio: curio support'
  'python-trio: trio support'
)
provides=('python-python_socks')
replaces=('python-python_socks')
source=("git+$url.git?signed#tag=v$pkgver")
b2sums=('b0410cb38def76de0ddb8fdd7369b0b68252bdf23f0ac35f9f434e4000efde8f792a407dd7831636865632409a86442e1d850d01503ec862fa44816b29040a6c')
validpgpkeys=('706A8935FED55F3BFB7C188B5EEAD7400A5D80ED') # Roman Snegirev <rsng@mail.ru>

build() {
  cd "${pkgname#python-}"
  python -m build --wheel --no-isolation
}

check() {
  cd "${pkgname#python-}"
  pytest
}

package() {
  cd "${pkgname#python-}"
  python -m installer --destdir="$pkgdir" dist/*.whl
}
