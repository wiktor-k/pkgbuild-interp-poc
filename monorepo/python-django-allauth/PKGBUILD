# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-django-allauth
_name=${pkgname#python-}
pkgver=65.19.4
pkgrel=1
pkgdesc="Authentication, registration, account management and 3rd party account authentication"
arch=(any)
url="https://codeberg.org/allauth/django-allauth"
license=(MIT)
# NOTE: Upstream seems to be of the opinion that indirectly depending on a set of libraries is just fine.
# So we have a growing list of false-positive dependencies that we are not supposed to add as direct dependencies...
# python-django-rest-framework: https://codeberg.org/allauth/django-allauth/issues/4431
# python-django-ninja: https://codeberg.org/allauth/django-allauth/issues/4432
depends=(
  python
  python-asgiref
  python-django
  python-pyyaml  # implicitly required but not specified
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-cryptography
  python-django-ninja
  python-django-rest-framework
  python-fido2
  python-oauthlib
  python-openid
  python-pillow
  python-pytest
  python-pytest-asyncio
  python-pytest-django
  python-pyjwt
  python-psycopg2
  python-qrcode
  python-requests
  python-requests-oauthlib
  python-saml
)
optdepends=(
  'python-cryptography: for social account support'
  'python-fido2: for MFA support'
  'python-openid: for OpenID and Steam support'
  'python-oauthlib: for OIDC IdP support'
  'python-pyjwt: for social account and OIDC IdP support'
  'python-qrcode: for MFA support'
  'python-requests: for social account support'
  'python-requests-oauthlib: for social account support'
  'python-saml: for SAML support'
)
source=("$pkgname::git+https://codeberg.org/allauth/django-allauth.git#tag=${pkgver}")
sha512sums=('975e8500a2cec466c9bf13d374349a299bee95afe0ac8c58a07518aa953439100058ac5045a0dbd56b126834d53559eb0edc93634d0d07c825e4872e0517190e')
b2sums=('622d234716b17c094ccf55da9fa8bab0386f7f7f3149ba2d7e159099829198c32a5b62f7d391abb9d80c50b4bd64facceef188fd1e4508868aadcf42bc8ae8f3')

build() {
  cd $pkgname
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd $pkgname
  pytest -vv
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
}
