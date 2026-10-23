# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: GI_Jack <GI_Jack@hackermail.com>

_pyname=minidump
pkgname=python-${_pyname}
pkgver=0.0.24
pkgrel=2
pkgdesc='Python library to parse and read Microsoft minidump file format'
url='https://github.com/skelsec/minidump'
arch=('any')
license=('MIT')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
source=("${url}/archive/refs/tags/${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('9a7652a0430b26eabdebfaf8278414c266f226634d1c268beccbf06d3961fdf4450e8acb025340353c9d4737c7bea9fd2a3afb0e1d9bf317be807d9685b9c2f9')
b2sums=('d99d54302990a9a6777a4f4771257fb4b41be212cda3afd6fd6b38b43415f159ad49b0b732da73634714acc9f0604cb0cde3878be6f5c45dc7894b77ea9ebd13')

build() {
  cd ${_pyname}-${pkgver}
  python -m build --wheel --no-isolation
}

package() {
  cd ${_pyname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
