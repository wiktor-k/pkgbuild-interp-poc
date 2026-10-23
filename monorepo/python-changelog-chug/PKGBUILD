# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: gilcu3

pkgname=python-changelog-chug
_pkgname=${pkgname#python-}
pkgver=0.0.3
pkgrel=5
pkgdesc="Parser for project Change Log documents"
url="https://git.sr.ht/~bignose/changelog-chug"
arch=('any')
license=('AGPL-3.0-or-later')
depends=('python' 'python-semver' 'python-docutils')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-testscenarios' 'python-testtools')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/release/${pkgver}.tar.gz")
sha256sums=('0777bacb46df129120d813ab198419735acb43870cdc32dae181c57e37eafaa7')

build() {
	cd "${_pkgname}-release/${pkgver}"
	python -m build --wheel --no-isolation
}

check() {
	cd "${_pkgname}-release/${pkgver}"
	python -m installer --destdir=tmp_install dist/*.whl
	local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
	export PATH="$PWD/tmp_install/usr/bin:$PATH"
	export PYTHONPATH="$PWD/tmp_install/$site_packages"
	python -m unittest discover -v
}

package() {
	cd "${_pkgname}-release/${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl
}
