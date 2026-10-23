# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Radu Potop <radu@wooptoo.com>

pkgname=python-a2wsgi
_pkgname=${pkgname#python-}
pkgver=1.10.10
pkgrel=2
pkgdesc="Converts WSGI apps to ASGI apps or the other way around"
url="https://github.com/abersheeran/a2wsgi"
arch=('any')
license=('Apache-2.0')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-pdm-backend')
checkdepends=('python-pytest' 'python-pytest-asyncio' 'python-httpx' 'python-starlette' 'python-baize')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/refs/tags/v${pkgver}.tar.gz")
sha256sums=('9b8b5594422ea243b73394a8cf00952e0a8a71e8bdd2fa8db730e6de157281f0')

build() {
	cd "${_pkgname}-${pkgver}"
	python -m build --wheel --no-isolation
}

check() {
	cd "${_pkgname}-${pkgver}"
	pytest
}

package() {
	cd "${_pkgname}-${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl
}

