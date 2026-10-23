# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-fqdn
pkgver=1.6.0
pkgrel=1
pkgdesc='RFC-compliant FQDN validation and manipulation for Python'
arch=(any)
url='https://github.com/ypcrts/fqdn'
license=(MPL-2.0)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-wheel
  python-setuptools
  python-setuptools-scm
)
checkdepends=(python-pytest)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('93cd770cf204905b323cb2801feb2d5bdf7b667287311cd5ef16aeb653646a27773226d9a7f5ec7120fbd6bbe4c7eb55e360dd58eb4da125fe035cdf5615c2a3')
b2sums=('fea451d0dcac167e1b791ce0252848e2eabfca9ef14d1293f1029082eb774fdc63af39a1c87b6cab9e5697268f2d7615e274633810a80ba85ecb1b767c2f115d')

build() {
  cd "$pkgname"

  export SETUPTOOLS_SCM_PRETEND_VERSION="$pkgver"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  pytest -v 
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # documentation
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" README.md
}
