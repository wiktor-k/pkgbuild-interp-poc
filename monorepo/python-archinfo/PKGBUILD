# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: RocketDev <ma2014119@outlook.com>

_pyname=archinfo
pkgname=python-${_pyname}
pkgver=9.2.221
pkgrel=1
pkgdesc='Classes with architecture-specific information useful to other projects'
url='https://github.com/angr/archinfo'
arch=('any')
license=(BSD-2-Clause)
depends=(
  python
)
makedepends=(
  python-build
  python-installer
  python-wheel
  python-setuptools
)
optdepends=(
  python-capstone
  python-unicorn
  python-keystone
  python-pyvex
  python-pypcode
)
checkdepends=(
  python-pytest
  python-pytest-xdist
  python-pypcode
)
source=("${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('d1cf2bd91bc3c7d618764ce1b511a6cdb84f0090a3901583e6c97d61845d4f83d5af88869ccceb3db746b1c6da058c31d48c80716fc19998986185aae435058a')
b2sums=('f8f0f4d1acafc210a534b732d8287b8c89c2f4d408bf3b19d40d2632c218ebf3a7e095b88c14a0e0b4fff247f25dc3ed914a51bd7effb732677ccc9b1d43fe4c')

build() {
  cd "${_pyname}-${pkgver}"
  python -m build --wheel --no-isolation
}

check() {
  cd "${_pyname}-${pkgver}"
  python -m venv --system-site-packages .venv
  .venv/bin/python -m installer dist/*.whl
  .venv/bin/python -m pytest tests
}

package() {
  cd "${_pyname}-${pkgver}"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
