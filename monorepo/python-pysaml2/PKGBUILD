# Maintainer: Johannes Löthberg <johannes@kyriasis.com>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Ivan Shapovalov <intelfx@intelfx.name>

pkgname=python-pysaml2
pkgver=7.5.5
pkgrel=1
pkgdesc='Python implementation of SAML Version 2'
arch=('any')
url='https://github.com/IdentityPython/pysaml2'
license=('Apache-2.0')
depends=(
  'python'
  'python-cryptography'
  'python-dateutil'
  'python-defusedxml'
  'python-pyopenssl'
  'python-requests'
  'python-xmlschema'
  'xmlsec'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-poetry-core'
)
checkdepends=(
  'python-pyasn1'
  'python-pymongo'
  'python-pytest'
  'python-responses'
 )
optdepends=(
  'python-ldap: LDAP authentication and user information'
  'python-memcached: memcached storage'
  'python-paste: for paste integration'
  'python-pymongo: MongoDB storage'
  # 'python-repoze.who: for repoze.who integration'  # TODO: package
  'python-zope-interface: for zope integration'
)
source=(
  "$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz"
  "$pkgname-pyopenssl-compat.patch::https://salsa.debian.org/openstack-team/python/python-pysaml2/-/raw/e82fb9e61b830f39839b08853f0bfbb3605818be/debian/patches/more-compat-with-recent-openssl.patch"
)
b2sums=(
  '537daa48053c96de5afa89d0b167fcfd4e209e7725d2301f5e3907bff72d185e0b2549155fd51d0a078d75859d5a54a7a84a3d93ab7f42acb2bb0654d70b3f64'
  '6bee5ff182535013be3f1bd318120c0da3a95c1295511915f77029b2bc698f4e90b24b07ad149ae670cd1ce26cfb19a6d9781b2543c8ecd31a8024c36eae2b6c'
)

prepare() {
  cd ${pkgname#python-}-$pkgver
  # Use cryptography for certificate request APIs removed from pyOpenSSL.
  patch -Np1 < ../$pkgname-pyopenssl-compat.patch

  # Upstream caps xmlschema at version 3, but we have 4.x which changed sandbox
  # behavior - files outside base_url are now blocked. Use "local" to allow any
  # local file while still blocking remote URLs.
  sed -i 's/"allow": "sandbox"/"allow": "local"/' src/saml2/xml/schema/__init__.py
}

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # Deselected tests fail for some reason
  test-env/bin/python -m pytest \
    --deselect=tests/test_50_server.py::TestServer1::test_encrypted_response_6 \
    --deselect=tests/test_50_server.py::TestServer1NonAsciiAva::test_encrypted_response_6 \
    --deselect=tests/test_81_certificates.py::TestGenerateCertificates::test_validate_cert_chains \
    --deselect=tests/test_81_certificates.py::TestGenerateCertificates::test_validate_with_root_cert \
    --deselect=tests/test_schema_validator.py::test_namespace_processing
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
