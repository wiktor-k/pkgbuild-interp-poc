# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Timothy Redaelli <timothy.redaelli@gmail.com>
# Contributor: Hilton Medeiros <medeiros.hilton@gmail.com>

pkgname=python-pbkdf2
pkgver=1.3
pkgrel=15
pkgdesc='Password-based key derivation function PBKDF2'
url='https://www.dlitz.net/software/python-pbkdf2/'
arch=('any')
license=('MIT')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
depends=('python')
checkdepends=('python-pytest')
optdepends=('python-crypto: to make use of PyCrypto`s HMAC and SHA')
source=("git+https://github.com/dlitz/python-pbkdf2.git#tag=v${pkgver}")
sha512sums=('a6dd129d4cc6053418834a80b552db702175ac1871be7591bfa0a65f6c53af5421a5eb869ff8302cd923c42dbebaf4c57d5a24aa4d6e900c6b32861500479517')

build() {
  cd python-pbkdf2
  python -m build --wheel --no-isolation
}

check() {
  cd python-pbkdf2
  pytest
}

package() {
  cd python-pbkdf2
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 README.txt "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
  install -Dm 644 README.txt -t "$pkgdir/usr/share/doc/$pkgname"
}

# vim: ts=2 sw=2 et:
