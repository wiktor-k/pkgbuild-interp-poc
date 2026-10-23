# Maintainer: Robin Candau <antiz@archlinux.org>

pkgname=python-pytest-instafail
_pkgname=pytest-instafail
pkgver=0.5.0
pkgrel=2
pkgdesc="py.test plugin to show failures instantly"
url="https://github.com/pytest-dev/pytest-instafail"
arch=('any')
license=('BSD-3-Clause')
depends=('python' 'python-pytest')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=("git+${url}.git#tag=v${pkgver}")
sha256sums=('2631c20de63da4d97536b99ad6932a0dbf0632fd21e5160dfd25e8984b813276')

build() {
        cd "${_pkgname}"
        python -m build --wheel --no-isolation
}

check() {
        cd "${_pkgname}"
        pytest
}

package() {
        cd "${_pkgname}"
        python -m installer --destdir="${pkgdir}" dist/*.whl

	install -Dm 644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
