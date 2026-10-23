# Maintainer: Bruno Pagani <archange@archlinux.org>

_pkg=QDarkStyle
pkgname=python-${_pkg,,}
pkgver=3.2.3
pkgrel=6
pkgdesc="Dark style sheet for QtWidgets application"
arch=(any)
url="https://github.com/ColinDuquesnoy/QDarkStyleSheet"
license=(MIT CCPL)
depends=(python-helpdev python-qtpy)
makedepends=(python-setuptools python-build python-installer)
checkdepends=(python-pytest python-pytest-qt qt6-tools)
# Pypi tarballs do not provide tests files anymore
#source=(https://files.pythonhosted.org/packages/source/${_pkg::1}/${_pkg}/${_pkg}-${pkgver}.tar.gz)
source=(${url}/archive/v.${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha256sums=('5c7412cbe2f348ac55284f9869736fb9918e78f00fcd929bfd40e6fd66237224')

build() {
  cd ${_pkg}Sheet-v.${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pkg}Sheet-v.${pkgver}
  python -m qdarkstyle.example --test --offscreen
}

package() {
  cd ${_pkg}Sheet-v.${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE.rst -t "${pkgdir}"/usr/share/licenses/${pkgname}
}
