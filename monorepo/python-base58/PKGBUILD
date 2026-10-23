# Maintainer: Jakub Klinkovský <lahwaacz at archlinux dot org>

_name=base58
pkgname=python-$_name
pkgver=2.1.1
pkgrel=3
pkgdesc="Base58 and Base58Check implementation compatible with what is used by the bitcoin network"
arch=(any)
url="https://github.com/keis/base58"
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=($pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz)
b2sums=('445bc3a44509c1de86b1131a02de6871e7bc4e58ec74ee1b2e3deff1f9227f2f359aeb25bcc4a46a728078bc062ee06f7dba7cf77542c1f21e836bfb582f8745')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 COPYING -t "$pkgdir"/usr/share/licenses/$pkgname/
}
