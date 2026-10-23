# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-electrum-ecc
pkgver=0.0.8
pkgrel=1
pkgdesc="Python library for libsecp256k1, using ctypes"
arch=('any')
url="https://github.com/spesmilo/electrum-ecc"
license=('MIT')
depends=(
  'libsecp256k1'
  'python'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=('python-pytest')
source=("git+$url.git#tag=$pkgver?signed")
b2sums=('7f447d4219cc895396ddfb2a7f598868fed04869230ad23ff9952ced4b7dfce0dd7b9129bf8c14a363285d47ba0f3ac42537480594f0b7cecd9e580197dbe881')
validpgpkeys=('4AD64339DFA05E20B3F6AD51E7B748CDAF5E5ED9') # SomberNight <somber.night@protonmail.com>

build() {
  cd ${pkgname#python-}
  export ELECTRUM_ECC_DONT_COMPILE=1
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
