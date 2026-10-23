# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-termstyle
pkgver=0.1.11
pkgrel=14
pkgdesc="Console colouring for python"
url="https://github.com/timbertson/termstyle"
license=('BSD-3-Clause')
arch=('any')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=("git+https://github.com/timbertson/termstyle.git#commit=ed18c5ed4bfee21267843f758094a5620d5f7929")
sha512sums=('fa562187e0a967d4fc112bd08e400b1531e62bcde0a9e883ce47546c501ab6e00dd861f98b0c507cdbb54f7f5b21d5e9e6bff105e8d521a18ea0f9f13fc5690d')

build() {
  cd termstyle
  python -m build --wheel --no-isolation
}

check() {
  cd termstyle
  python test3.py
}

package() {
  cd termstyle
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
