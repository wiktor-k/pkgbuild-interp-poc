# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Piotr Dąbrowski <piotr_dab at op dot pl>
# Contributor: Letu Ren <fantasquex at gmail dot com>
# Contributor: workonfire <kolucki62@gmail.com>

pkgname=python-onigurumacffi
_pkgname=${pkgname#python-}
pkgver=1.5.0
pkgrel=1
pkgdesc="Python cffi bindings for the oniguruma regex engine"
url="https://github.com/asottile/onigurumacffi"
arch=('x86_64')
license=('MIT')
depends=('glibc' 'oniguruma' 'python' 'python-cffi')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-pytest')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/v${pkgver}.tar.gz")
sha256sums=('891a54cc4bcf5f8035d9a1841c751bd557d37d243296f8f8a08907384481c5ae')

build() {
	cd "${_pkgname}-${pkgver}"
	python -m build --wheel --no-isolation
}

check() {
	cd "${_pkgname}-${pkgver}"
	python -m venv --system-site-packages test-env
	test-env/bin/python -m installer dist/*.whl
	test-env/bin/python -m pytest
}

package() {
	cd "${_pkgname}-${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl

	install -Dm 644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
