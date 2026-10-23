# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Massimiliano Torromeo <massimiliano.torromeo@gmail.com>

pkgname=python-augeas
pkgver=1.2.0
pkgrel=1
pkgdesc="Python bindings for Augeas"
arch=('any')
license=('LGPL-2.1-or-later')
url="https://augeas.net"
depends=('augeas' 'python-cffi')
makedepends=('python-build' 'python-installer' 'python-setuptools')
source=("https://files.pythonhosted.org/packages/source/p/python-augeas/python-augeas-$pkgver.tar.gz")
sha512sums=('5d462263cb1857694796500ff412cf53d3800a030277a53ba9d84dda9dbcbdd7693a3ae7eec467d816aa07dfa27c5b06cc0cf6470bfa8bf72e45c56e192c431f')

build() {
  cd python-augeas-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd python-augeas-$pkgver
  python -m venv --system-site-packages testenv
  testenv/bin/python -m installer dist/*.whl
  testenv/bin/python -m unittest discover -v
}

package() {
  cd python-augeas-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
