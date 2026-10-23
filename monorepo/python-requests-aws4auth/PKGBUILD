# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Bruno Galeotti <bgaleotti at gmail dot com>

_pkgname=requests-aws4auth
pkgname=python-requests-aws4auth
pkgver=1.4.0
pkgrel=1
pkgdesc='Amazon Web Services version 4 authentication for the Python Requests module'
url='https://github.com/sam-washington/requests-aws4auth'
arch=('any')
license=('MIT')
depends=('python' 'python-requests')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-pytest' 'python-httpx')
optdepends=('python-httpx')
source=(${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('406263b74485638c61cbe47946782c2c3298df89c083a1399c742045b3a5ef89e1d340c2771da435ecfb2cd0aa4b9eab5b42a4938ba01705e80caf3e8cc22df4')
b2sums=('c4ea5ce854703762de3c9f6560290e34140ed196a60dab2694208dfee1faccb57b0dc919cce8708a20d19a7e0fba429acbbcd617af43432ede14243c616e478a')

prepare() {
  cd ${_pkgname}-${pkgver}
  sed '/./{H;$!d} ; x ; s/package_data=.*\n.*]},\n//' -i setup.py
}

build() {
  cd ${_pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pkgname}-${pkgver}
  python -m pytest
}

package() {
  cd ${_pkgname}-${pkgver}
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
