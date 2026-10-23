# Maintainer: Bruno Pagani <archange@archlinux.org>

_pkg=whatthepatch
pkgname=python-${_pkg}
pkgver=1.0.7
pkgrel=2
pkgdesc="A Python patch parsing library"
arch=(any)
url="https://github.com/cscorley/whatthepatch"
license=(MIT)
depends=(python)
makedepends=(python-build python-installer python-setuptools python-wheel)
checkdepends=(python-pytest)
source=(https://files.pythonhosted.org/packages/source/${_pkg::1}/${_pkg}/${_pkg}-${pkgver}.tar.gz)
#source=(${url}/archive/v${pkgver}rel/${pkgname}-${pkgver}.tar.gz)
sha256sums=('9eefb4ebea5200408e02d413d2b4bc28daea6b78bb4b4d53431af7245f7d7edf')

build() {
  cd ${_pkg}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pkg}-${pkgver}
  PYTHONPATH="${PWD}"/build/lib/ pytest -vv --color=yes
}

package() {
  cd ${_pkg}-${pkgver}
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm644 LICENSE -t "${pkgdir}"/usr/share/licenses/${pkgname}
}
