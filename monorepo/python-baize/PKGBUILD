# Maintainer: Robin Candau <antiz@archlinux.org>

pkgname=python-baize
_pkgname=${pkgname#python-}
pkgver=0.23.1
pkgrel=1
pkgdesc="Powerful and exquisite WSGI/ASGI framework/toolkit"
url="https://github.com/abersheeran/baize"
arch=('any')
license=('Apache-2.0')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-pdm-backend' 'python-setuptools')
checkdepends=('python-pytest' 'python-pytest-asyncio' 'python-httpx' 'python-starlette')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/refs/tags/v${pkgver}.tar.gz")
sha256sums=('bb0c3ba87d3d9440a30abb61af9a2e576631b5ba268161f3490a847656222823')

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

