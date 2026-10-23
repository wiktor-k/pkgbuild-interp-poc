# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Daniel Wallace <danielwallace at gtmanfred dot com>

pkgname=python-os-service-types
pkgver=1.9.0
pkgrel=1
pkgdesc="Python library for consuming OpenStack sevice-types-authority data"
arch=('any')
url="https://pypi.python.org/pypi/os-service-types/$pkgver"
license=('Apache-2.0')
depends=('python-pbr' 'python-typing_extensions')
makedepends=('git' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-stestr' 'python-oslotest' 'python-testscenarios' 'python-requests-mock'
              'python-keystoneauth1' 'python-six')
options=('!emptydirs')
source=("git+https://github.com/openstack/os-service-types.git#tag=$pkgver")
sha512sums=('fd1b5f5d3591b4679e53002aede3be145c0f114195b603489f3c22e887fefad6b5e5a569bb21e8c4b5686d9f7b51010736a17dd67bd36406fb982fa80f802239')

build() {
  cd os-service-types
  python -m build --wheel --no-isolation
}

check() {
  cd os-service-types
  stestr run
}

package() {
  cd os-service-types
  python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim:set ts=2 sw=2 et:
