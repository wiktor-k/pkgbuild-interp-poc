# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-backrefs
pkgver=8.0
pkgrel=1
pkgdesc="Wrapper around re or regex that adds additional back references"
arch=('any')
url="https://github.com/facelessuser/backrefs"
license=('MIT')
depends=('python')
makedepends=(
  'python-build'
  'python-hatchling'
  'python-installer'
  'python-wheel'
)
checkdepends=(
  'python-pytest'
  'python-regex'
)
optdepends=('python-regex: use regex instead of re')
source=("$url/archive/$pkgver/${pkgname#python-}-$pkgver.tar.gz")
b2sums=('e64730d1f82a75b46f9c635e0be5330b99bee41fec6eb19670fb0537ac5f656c0782527cc567ebfdd73c103594af4df00e0da939bf902e6b3bfb5e4b5733f85a')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  pytest
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.md
}
