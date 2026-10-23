# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Jiachen Yang <farseerfc@gmail.com>
# Contributor: Sam Stuewe <halosghost at archlinux dot info>
# Contributor: Chaoswizard <chaoswizard@gmail.com>

pkgname=python-feedgenerator
pkgver=2.2.1
pkgrel=2
pkgdesc="Standalone version of Django's feedgenerator module"
url="https://github.com/getpelican/feedgenerator"
arch=('any')
license=('BSD-3-Clause')
depends=('python')
makedepends=(
  'python-build'
  'python-hatchling'
  'python-installer'
  'python-wheel'
)
checkdepends=('python-pytest')
source=("$url/archive/$pkgver/${pkgname#python-}-$pkgver.tar.gz")
sha512sums=('83f49e1aabfdc72bb95b2134a55ce8dbae1cd8bd75a7bf48d50cd3ca4afb5c6021eb04835e6e0d04e131acd9ce3bb6901e5ada42cf327a540ffe1dd514f1d21e')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  pytest --override-ini="addopts="
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
