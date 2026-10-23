# Maintainer: Bruno Pagani <archange@archlinux.org>

_pkg=entrypoint2
pkgname=python-${_pkg}
pkgver=1.1
pkgrel=8
pkgdesc="Easy to use command-line interface for python modules"
arch=(any)
url="https://github.com/ponty/entrypoint2"
license=(BSD-2-Clause)
depends=(python)
makedepends=(python-build python-installer python-setuptools python-wheel)
checkdepends=(python-pytest python-easyprocess)
source=(https://files.pythonhosted.org/packages/source/${_pkg::1}/${_pkg}/${_pkg}-${pkgver}.tar.gz)
#source=(${url}/archive/${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha256sums=('fc0b7fe7b21acdab47a585ab9407ca7e5c4f96cb6888575db6b0ceb91f0e105a')

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
  install -Dm644 LICENSE.txt -t "${pkgdir}"/usr/share/licenses/${pkgname}
}
