# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: gilcu3
# Contributor: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Contributor: Anatol Pomozov
# Contributor: Massimiliano Torromeo <massimiliano.torromeo@gmail.com>

pkgname=python-daemon
pkgver=3.1.2
pkgrel=6
pkgdesc="Python library to implement a well-behaved Unix daemon process"
url="https://pagure.io/python-daemon/"
arch=('any')
license=('Apache-2.0')
depends=('python' 'python-filelock')
makedepends=('python-build' 'python-changelog-chug' 'python-docutils' 'python-installer' 'python-wheel')
checkdepends=('python-testscenarios' 'python-testtools')
source=("https://files.pythonhosted.org/packages/source/${pkgname::1}/${pkgname//-/_}/${pkgname//-/_}-${pkgver}.tar.gz"
        'python-daemon-3.1.2-lock-file-conversion.patch')
sha256sums=('f7b04335adc473de877f5117e26d5f1142f4c9f7cd765408f0877757be5afbf4'
            '6f2a7fc9df4b07c15552b1dce60f939d6b58ea0b7d462fc68032052a4c2b4aa2')

prepare() {
	cd "${pkgname//-/_}-${pkgver}"
	# Replace deprecated lockfile dependency by filelock
	# https://gitlab.archlinux.org/archlinux/packaging/packages/python-lockfile/-/issues/1#note_388414
	# https://pagure.io/python-daemon/issue/42
	# https://pagure.io/python-daemon/pull-request/81
	patch -Np1 -i "${srcdir}/python-daemon-3.1.2-lock-file-conversion.patch"
}

build() {
	cd "${pkgname//-/_}-${pkgver}"
	python -m build --wheel --no-isolation
}

check() {
	cd "${pkgname//-/_}-${pkgver}"
	python -m installer --destdir=tmp_install dist/*.whl
	local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
	export PATH="$PWD/tmp_install/usr/bin:$PATH"
	export PYTHONPATH="$PWD/tmp_install/$site_packages"
	python -m unittest discover -v
}

package() {
	cd "${pkgname//-/_}-${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl
}
