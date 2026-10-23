# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Letu Ren <fantasquex at gmail dot com>

_pyname=pyxbe
pkgname=python-${_pyname}
pkgver=1.0.3
pkgrel=2
pkgdesc='Library to work with XBE files, the executable file format for the original Xbox game console'
url='https://github.com/mborgerson/pyxbe'
arch=('any')
license=('custom:MIT')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
checkdepends=('python-pytest')
source=("${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('77357e55e33fa76c4d4727debc2d47e742df9a95093fc5c7156e84e9ed032a859f170b77e46ccc72d1654c5da2e7c249851fad08482d024636dedd0ba7cb5a86')
b2sums=('6b6e50e4197c2d9e3f61b82a1bed2bbcb41ce692fc8107ea2532b972280f30834714b9a75158fd2af3e4459ce010c46dcfa7ed8d2ac5d8f621421eeaac0599b5')

build() {
  cd ${_pyname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pyname}-${pkgver}/tests
  PYTHONPATH="${PWD}/../build/lib" pytest
}

package() {
  cd ${_pyname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE.txt -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
