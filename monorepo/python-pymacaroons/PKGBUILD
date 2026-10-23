# Maintainer: Johannes Löthberg <johannes@kyriasis.com>
# Contributor: Ivan Shapovalov <intelfx100@gmail.com>

pkgname=python-pymacaroons
pkgver=0.13.0
pkgrel=15

pkgdesc="Python implementation of Macaroons. They’re better than cookies!"
url='https://github.com/ecordell/pymacaroons'
arch=('any')
license=('MIT')

depends=('python'
         'python-pynacl')
makedepends=('python-build'
             'python-installer'
             'python-setuptools'
             'python-wheel')

conflicts=('python-pymacaroons-pynacl')

source=("pymacaroons-$pkgver.tar.gz::https://github.com/ecordell/pymacaroons/archive/v$pkgver.tar.gz"
        'pymacaroons-drop-python2-mock-six.patch::https://github.com/ecordell/pymacaroons/commit/68e096119d5aceee33fb817fd9ee5434c5857a78.patch')

sha256sums=('780c67643126afe56f57fdc0f82b952d08e01a4df7eefaa929766dfb6edf6580'
            'a720c58d679149601bfdb38f519313dd4fc7bb07433e170a2c99237136b75bf7')

prepare() {
	cd pymacaroons-$pkgver
	patch --forward --strip=1 --input=../pymacaroons-drop-python2-mock-six.patch || true
}

build() {
	cd pymacaroons-$pkgver
	python -m build --wheel --skip-dependency-check --no-isolation
}

package() {
	local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
	install -d "$pkgdir"/usr/share/licenses/$pkgname
	ln -s "$site_packages"/pymacaroons-$pkgver.dist-info/licenses/LICENSE \
		"$pkgdir"/usr/share/licenses/$pkgname

	cd pymacaroons-$pkgver
	python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim: set ts=4 sw=4 tw=0 ft=PKGBUILD :
