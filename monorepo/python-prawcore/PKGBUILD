# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-prawcore
pkgver=4.0.0
pkgrel=1
pkgdesc="Low-level communication layer for PRAW 4+."
arch=('any')
license=('BSD-2-Clause')
url="https://github.com/praw-dev/prawcore"
depends=('python' 'python-requests')
makedepends=('git' 'python-build' 'python-installer' 'python-hatchling')
checkdepends=('python-pytest' 'python-vcrpy')
source=("git+https://github.com/praw-dev/prawcore.git#tag=v$pkgver")
sha512sums=('8ff448c32fad66b3b17bea3292331935fba6101f67ba07aa1cdb1607381d3d8f0df42349cccec8d125900bc1f75bfc995d4e44347e9ef6db27ce6852da3caa40')

build() {
  cd prawcore
  python -m build -nw
}

check() {
  cd prawcore
  pytest
}

package() {
  cd prawcore
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
