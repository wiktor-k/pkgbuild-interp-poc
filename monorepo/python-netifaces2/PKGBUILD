# Maintainer: Robin Candau <antiz@archlinux.org>

pkgname=python-netifaces2
_pkgname=netifaces-2
pkgver=0.0.22
pkgrel=3
pkgdesc="Portable network interface information (Supports Python 3.6 and higher)"
url="https://pypi.org/project/netifaces2"
arch=('x86_64')
license=('MIT')
depends=('python')
makedepends=('cargo' 'python-build' 'python-installer' 'python-wheel' 'python-maturin')
checkdepends=('python-pytest')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/SamuelYvon/netifaces-2/archive/refs/tags/V${pkgver}.tar.gz")
sha256sums=('1680050452f67cd2de204326db75ec040bbe5a27b2921109402de8e36679c704')

prepare() {
	cd "${_pkgname}-${pkgver}"
	cargo fetch --locked
}

build() {
	cd "${_pkgname}-${pkgver}"
	python -m build --wheel --no-isolation --config-setting="build-args=--frozen"
}

check(){
	cd "${_pkgname}-${pkgver}"
	python -m venv --system-site-packages test-env
	test-env/bin/python -m installer dist/*.whl
	PATH="${PWD}/test-env/bin:${PATH}" test-env/bin/python -m pytest
}

package() {
	cd "${_pkgname}-${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl
	install -Dm 644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
