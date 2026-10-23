# Maintainer: Bruno Pagani <archange@archlinux.org>

_pkg=smbprotocol
pkgname=python-${_pkg}
pkgver=1.17.0
pkgrel=1
pkgdesc="SMBv2 and v3 client"
arch=(any)
url="https://github.com/jborean93/smbprotocol"
license=(MIT)
depends=(python-cryptography python-pyspnego)
makedepends=(git python-setuptools python-build python-installer python-wheel)
checkdepends=(python-pytest python-pytest-mock python-gssapi)
optdepends=('python-gssapi: Kerberos support')
# tags are not signed, commits only
_commit=190667f41c5c8856db741efc133c27a624a7a2bf # git rev-parse v${pkgver}
#source=(git+${url}.git#commit=${_commit}?signed)
# 1.9.0 was signed using GitHub key
source=(git+${url}.git#commit=${_commit})
sha256sums=('f675c3d7da954954cb1e51e35e638cb2c4734fa3df16e17586c2f0e7cd4a8fb1')
validpgpkeys=(A03A0E13C2EFFD384B1EC39A2AAC89085FBBDAB5) # Jordan Borean <jborean93@gmail.com>

build() {
  cd ${_pkg}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pkg}
  PYTHONPATH="${PWD}"/src pytest -vv --color=yes
}

package() {
  cd ${_pkg}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "${pkgdir}"/usr/share/licenses/${pkgname}
}
