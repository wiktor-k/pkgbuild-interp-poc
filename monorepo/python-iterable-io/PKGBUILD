# Maintainer: Caleb Maclennan <caleb@alerque.com>

_pyname=iterable-io
pkgname=python-$_pyname
pkgver=1.0.4
pkgrel=1
pkgdesc='Adapt generators and other iterables to a file-like interface'
arch=(any)
url='https://github.com/pR0Ps/iterable-io'
license=(LGPL-3.0-only)
depends=(python)
makedepends=(python-{build,installer,wheel}
             python-hatchling)
checkdepends=(python-pytest)
_archive="${_pyname/-/_}-$pkgver"
source=("https://files.pythonhosted.org/packages/source/${_pyname::1}/$_pyname/$_archive.tar.gz")
sha256sums=('63b5b394aedbc61b4409a0d213130a4d153c602d1eac9dcff860a43ae0340405')

build() {
	cd "$_archive"
	python -m build -wn
}

check() {
	cd "$_archive"
	export PYTHONPATH=src
	pytest
}

package() {
	cd "$_archive"
	python -m installer -d "$pkgdir" dist/*.whl
}
