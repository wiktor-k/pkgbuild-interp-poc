# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Daniel Wallace <danielwallace at gtmanfred dot com>

pkgname=python-cliff
pkgver=4.17.0
pkgrel=1
pkgdesc="Command Line Interface Formulation Framework"
arch=('any')
url="https://docs.openstack.org/cliff/latest/"
license=('Apache-2.0')
depends=('python-autopage' 'python-cmd2' 'python-prettytable' 'python-stevedore' 'python-yaml')
makedepends=('git' 'python-build' 'python-installer' 'python-wheel' 'python-pbr')
checkdepends=('python-fixtures' 'python-stestr' 'python-sphinx' 'python-testtools')
options=('!emptydirs')
source=("git+https://github.com/openstack/cliff.git#tag=$pkgver")
sha512sums=('d1235296f1c8d29e16164e50b5030131d69f96c2679bcba0c24906c0d5587d1a7e18fecedea7492cf7b2a5e6608ac33ffb4a00a2a1cde550cfc4073edf114cf8')

build() {
  cd cliff
  python -m build --wheel --no-isolation
}

check() {
  cd cliff
  PYTHONPATH="$PWD/src" stestr run
}

package() {
  cd cliff
  python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim:set ts=2 sw=2 et:
