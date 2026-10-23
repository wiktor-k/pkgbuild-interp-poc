# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Anatol Pomozov

pkgname=python-txaio
# https://github.com/crossbario/txaio/blob/master/docs/releases.rst
pkgver=26.6.1
pkgrel=1
pkgdesc='Compatibility API between asyncio/Twisted/Trollius'
arch=('any')
url="https://github.com/crossbario/txaio"
license=('MIT')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-hatchling')
# tests import asyncio.test_utils from python-tests.
checkdepends=('python-pytest' 'python-twisted' 'python-tests')
optdepends=(
  'python-twisted: Twisted support'
  'python-zope-interface: Twisted support'
)
source=("git+https://github.com/crossbario/txaio.git#tag=v$pkgver")
sha512sums=('d760aec24b5dabaa76fcfd0d0b757661da09e154342c85792b06bb26d55821f7b5326e2afb0b1010d028b1f42f2b7722baaa4da9f9a799d8c463183c91b8f96a')

build() {
  cd "$srcdir"/txaio
  python -m build --wheel --no-isolation
}

check() {
  cd "$srcdir"/txaio
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd txaio
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}
