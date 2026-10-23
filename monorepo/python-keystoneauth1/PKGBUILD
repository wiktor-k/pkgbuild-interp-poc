# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Daniel Wallace <danielwallace at gtmanfred dot com>

pkgname=python-keystoneauth1
pkgver=5.18.1
pkgrel=1
pkgdesc="Authentication Library for OpenStack Identity"
arch=('any')
url="https://docs.openstack.org/keystoneauth/latest/"
license=('Apache-2.0')
depends=('python-iso8601' 'python-stevedore' 'python-os-service-types' 'python-requests' 'python-typing_extensions')
optdepends=('python-requests-kerberos: for kerberos support'
            'python-lxml: for saml2 support'
            'python-oauthlib: for oauth1 support'
            'python-betamax: for betamax support'
            'python-fixtures: for betamax support'
            'python-yaml: for betamax support')
makedepends=('git' 'python-build' 'python-installer' 'python-pbr' 'python-setuptools' 'python-wheel')
checkdepends=('python-hacking' 'python-oslo-config' 'python-oslo-utils' 'python-oslotest'
              'python-betamax' 'python-fixtures' 'python-requests-mock' 'python-stestr'
              'python-testresources' 'python-testtools' 'python-yaml'
              'python-requests-kerberos' 'python-lxml' 'python-oauthlib')
source=("git+https://github.com/openstack/keystoneauth.git#tag=$pkgver")
sha512sums=('74fe72bb4209dfde7027a71ed3acc92f40bacbb98e6017ec0b930b1e59b760b9b94bae84d760bdb83a25d601eb6c2ff1aeab5794d100315c32f07113576253d0')

build() {
  cd keystoneauth
  python -m build --wheel --no-isolation
}

check() {
  cd keystoneauth
  stestr run --exclude-regex '(.*test_keystoneauth_betamax_fixture)'
}

package() {
  cd keystoneauth
  python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim:set ts=2 sw=2 et:
