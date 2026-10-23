# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Letu Ren <fantasquex at gmail dot com>

_pyname=pypcode
pkgname=python-${_pyname}
pkgver=3.3.3
pkgrel=1
pkgdesc="Python bindings to Ghidra's SLEIGH library for disassembly and lifting to p-code IR"
url='https://github.com/angr/pypcode'
license=(BSD-2-Clause)
arch=(x86_64)
depends=(
  gcc-libs
  glibc
  nanobind
  python
  python-cffi
)
makedepends=(
  cmake
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
)
source=("${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('f68d4a86d2c2f9f70fcebae7fabeddbc8c83a6378fdcc96392b7f648f38eaab9aaf8c8e69a95a31010080c50cb9898f10633dc1d91c45a6ffd5e5ec08b05cdc3')
b2sums=('ade8d8870ac861e2f341cf042ecdafe14fb4afb4f1e1c658b5e0ed22348c59761685833c45f304c4d0a5dcf892920446e842eb578276e048793f0d27e20c69a5')

build() {
  cd ${_pyname}-${pkgver}
  python -m build --wheel --no-isolation --skip-dependency-check
}

package() {
  cd ${_pyname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE.txt -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
