# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Mantas Mikulėnas <grawity@gmail.com>

pkgname=python-requests-gssapi
pkgver=1.4.0
pkgrel=2
pkgdesc="GSSAPI (HTTP Negotiate) authentication for Python-Requests"
arch=(any)
url="https://github.com/pythongssapi/requests-gssapi"
license=(ISC)
depends=(
  python
  python-gssapi
  python-requests
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('e4d3f5ad36d309239e2e7659e55f208981b97ee6da2433f69749fd71bcb16a16')

build() {
  cd "${pkgname#python-}-$pkgver"
  python -m build --wheel --no-isolation
}

check() {
  cd "${pkgname#python-}-$pkgver"
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd "${pkgname#python-}-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
