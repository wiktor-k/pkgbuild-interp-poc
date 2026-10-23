# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: RocketDev <ma2014119@outlook.com>

_pyname=claripy
pkgname=python-${_pyname}
pkgver=9.2.221
pkgrel=1
pkgdesc='Abstraction layer for constraint solvers'
url='https://github.com/angr/claripy'
license=(BSD-2-Clause)
arch=(any)
depends=(
  python
  python-cachetools
  python-typing_extensions
  python-z3-solver
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-xdist
)
source=("${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('5e99fc6f93aad556f65ab95d232cea12ad2bf526882602d63a64f1d491ce78c4c287f031e8e45ab88ffae63714c89e51a839c7308848cc6b2ac590771ee50471')
b2sums=('3183bf0912edf38de317a9cfb640aafd0ee266b427f0b2b136ebf50f373bf6945563b8286707f9db7072e952dbd8494b6c898d5381be237bd38d0caded22d777')

build() {
  cd ${_pyname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pyname}-${pkgver}
  PYTHONPATH=build/lib pytest "${pytest_options[@]}"
}

package() {
  cd ${_pyname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
