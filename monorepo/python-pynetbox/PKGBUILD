# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: gardar <aur@gardar.net>
pkgname=python-pynetbox
_name="${pkgname#python-}"
pkgver=7.9.0
pkgrel=1
pkgdesc="NetBox API client library"
arch=('any')
url="https://github.com/netbox-community/pynetbox"
license=('Apache-2.0')
depends=('python' 'python-packaging' 'python-requests')
makedepends=('python-build' 'python-installer' 'python-setuptools-scm' 'python-wheel')
checkdepends=('python-pytest' 'python-pyyaml')
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('570c503c5e73969efc01d67808962a28e7b4030cc89a7109c9e73e6e17576a491180a03dd7d9c6d601dff41d5d6f1a090b22cc007b6685e7f3000d24a88f3349')
b2sums=('a787fc2f05f051b0fef9c4fe41a4f3eb29ff697dcf3c856ec54aa110e187ca4f4175cf1824c02476116c6f0ed238da4ae777d554139f6151278476c860058361')

build() {
  cd "$_name-${pkgver}"
  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --no-isolation
}

check() {
  cd "$_name-${pkgver}"
  pytest -vv --ignore tests/integration
}

package() {
  cd "$_name-${pkgver}"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -D -m644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
