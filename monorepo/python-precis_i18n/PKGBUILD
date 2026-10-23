# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Robin Candau <antiz@archliunx.org>

_pkgname=precis_i18n
pkgname=python-precis_i18n
pkgver=1.1.2
pkgrel=2
pkgdesc='PRECIS framework (RFC 8264, RFC 8265, RFC 8266)'
url='https://github.com/byllyfish/precis_i18n'
arch=('any')
license=('MIT')
depends=('python')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-pytest')
source=(${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('7b5b12f1dad22dd90731c47a8b3fe36ed273d56fb454173284e776dc3ac19ac19be6936b9b7f70df1b1de3b7de2fa5ce8cd08f2f120440c1850ea2c5885163d2')
b2sums=('a3b313118faebf89ead8b38fe7923b6f91b032925b99f384f252e35b9b70953222c200005de4c37c712d17860210ee518081ad92dfec3cadf518e585920ad9a0')

build() {
  cd ${_pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pkgname}-${pkgver}
  pytest
}

package() {
  cd ${_pkgname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE.txt -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 README.md CHANGELOG.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
