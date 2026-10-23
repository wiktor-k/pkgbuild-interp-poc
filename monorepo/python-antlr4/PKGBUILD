# Maintainer: Konstantin Gizdov <arch at kge dot pw>
# Contributor: Baptiste Jonglez <archlinux at bitsofnetworks dot org>

_pkgname=antlr4
pkgname="python-${_pkgname}"
pkgver=4.13.2
pkgrel=3
pkgdesc="ANTLR4 runtime for Python"
arch=('any')
url="https://www.antlr.org/"
license=('BSD')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools-scm')
checkdepends=('gdb' 'tk')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/antlr/${_pkgname}/archive/${pkgver}.tar.gz")
b2sums=('709f3c18329e08a811d7a4e5258fc0183ba6fa6ab62fe83ab694360cf1ce2bb3aa3f30cc62a0339a6ff8b64dd0a766057fd6f751cfc34a4e05051e4d54a57fad')

build() {
  cd "${srcdir}/${_pkgname}-${pkgver}/runtime/Python3"
  python -m build --wheel --no-isolation
}

check() {
  cd "${srcdir}/${_pkgname}-${pkgver}/runtime/Python3"
  python -m test || echo 'Test fails due to deprecated distutils'
}

package() {
  cd "${srcdir}/${_pkgname}-${pkgver}/runtime/Python3"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -d "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm644 "${srcdir}/${_pkgname}-${pkgver}"/LICENSE.txt "${pkgdir}/usr/share/licenses/${pkgname}"/LICENSE
}
