# Maintainer: Johannes Löthberg <johannes@kyriasis.com>
# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Ivan Shapovalov <intelfx@intelfx.name>

pkgname=python-frozendict
pkgver=2.4.7
pkgrel=1
epoch=1
pkgdesc='An immutable dictionary'
url='https://pypi.python.org/pypi/frozendict/'
arch=('any')
license=('LGPL-3.0-only')
depends=('python')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-pytest')
source=("https://pypi.org/packages/source/f/frozendict/frozendict-$pkgver.tar.gz")
sha512sums=('91e77b144a85ab7ff1240b303640323274150544d5adaf25615ed8acfd5353bcf37a149d3ff9d79ca6fd54d99c4fbb43fdd38ae2cf76a77d1c70c44912d2d96d')
b2sums=('37217d8afda8e4f44bb11747a790f0fddd6560e461f3a58581801aa8c292edb2d5c8276840a0fd8a8978151e2ee65703be9a024f3e9f291359f38ea3f4d7a2c3')

build() {
	cd frozendict-$pkgver
	python -m build --wheel --no-isolation
}

check(){
	cd frozendict-$pkgver
	local python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
	PYTHONPATH="$PWD/build/lib.linux-$CARCH-cpython-$python_version" pytest
}

package() {
	cd frozendict-$pkgver
	python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim: set ts=4 sw=4 tw=0 ft=PKGBUILD :
