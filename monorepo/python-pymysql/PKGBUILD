# Maintainer: David Runge <dvzrv@archlinux.org>

_name=PyMySQL
pkgname=python-pymysql
pkgver=1.2.3
pkgrel=1
pkgdesc="Pure Python MySQL Client"
arch=(any)
url="https://github.com/PyMySQL/PyMySQL/"
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
# NOTE: tests require a running instance of mariadb
optdepends=(
  'mariadb: for using a local MariaDB instance'
  'python-cryptography: for RSA cryptography'
  'python-pynacl: for ed25519 cryptography'
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('a0703c9cc04c985cb26cd0d954ccc4f7b9ebed51ccd2f96ba03a0951061129026801f66b4fd16c663cebcf3c69df00bc111448b5770321d349ee97487ecf2476')
b2sums=('a1d7928c67028a23268b5cca3bbd2606c73f699c4c8ebcb135f2f8736f392f5900385439b6e8076f4f11d6d71b172eb0d1c687aa4e70bd5b1ec0a59e72dcfb2e')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGELOG,README}.md -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
