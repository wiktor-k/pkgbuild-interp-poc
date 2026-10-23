# Maintainer: Morten Linderud <foxboron@archlinux.org>
# Contibutor: Pierre Neidhardt <ambrevar@gmail.com>

pkgname=python-keyutils
pkgver=0.6
pkgrel=12
pkgdesc="A set of python bindings for keyutils"
arch=("x86_64")
url="https://github.com/sassoftware/python-keyutils"
license=("Apache-2.0")
depends=("python" "keyutils")
makedepends=("python" "python-setuptools" "python-build" "python-installer" "python-wheel" "cython")
checkdepends=("python-pytest")
#source=("$pkgname-$pkgver.tar.gz::https://github.com/sassoftware/python-keyutils/archive/$pkgver.tar.gz")
# github repo is gone
source=("https://pypi.io/packages/source/k/keyutils/keyutils-${pkgver}.tar.gz")
sha256sums=('a9cfbacc33a5ada165531ff2d6e1241ec93a9f0cabc63205d6ff73f8eff33752')

prepare() {
  # Remove dependency on pytest-runner
  sed -i '/pytest-runner/d' "keyutils-$pkgver"/setup.py
}

build(){
	cd "keyutils-$pkgver"
	cython keyutils/_keyutils.pyx
	python -m build --wheel --no-isolation
}

package(){
	cd "keyutils-$pkgver"
	python -m installer --destdir="$pkgdir" dist/*.whl
}
