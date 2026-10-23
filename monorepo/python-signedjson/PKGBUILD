# Maintainer: Johannes Löthberg <johannes@kyriasis.com>
# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Ivan Shapovalov <intelfx@intelfx.name>

pkgname=python-signedjson
pkgver=1.1.4
pkgrel=5

pkgdesc='Sign JSON with Ed25519 signatures'
url='https://pypi.python.org/pypi/signedjson/'
arch=('any')
license=('Apache-2.0')

depends=('python' 'python-canonicaljson' 'python-unpaddedbase64' 'python-pynacl')
makedepends=('python-setuptools' 'python-setuptools-scm' 'python-wheel' 'python-build' 'python-installer')

source=("https://files.pythonhosted.org/packages/source/s/signedjson/signedjson-$pkgver.tar.gz")

sha256sums=('cd91c56af53f169ef032c62e9c4a3292dc158866933318d0592e3462db3d6492')

build() {
	cd signedjson-$pkgver
	python -m build --wheel --no-isolation
}

package() {
	cd signedjson-$pkgver
	python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim: set ts=4 sw=4 tw=0 ft=PKGBUILD :
