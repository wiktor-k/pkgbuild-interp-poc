# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Maintainer: Filipe Laíns (FFY00) <lains@archlinux.org>

pkgname=python-aiomysql
pkgver=0.3.2
pkgrel=3
pkgdesc="Library for accessing a MySQL database from the asyncio"
arch=(any)
url="https://github.com/aio-libs/aiomysql"
license=(MIT)
depends=(
  python
  python-pymysql
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
# checkdepends=(
#   python-docker
#   python-ipdb
#   python-pytest
#   python-sqlalchemy
#   python-uvloop
# )
optdepends=('python-sqlalchemy: SQLAlchemy support')
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('da4d1ef5fa2fee271e200e4c9c297d7837a90486c3d5cc22349a96e539c267befe4eb93cf086e247f3cbb0e2f72bc9e12c3ee7e4ca99d3d988d47a741f411640')

build() {
  cd ${pkgname#python-}-$pkgver
  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver
  python -m build --wheel --no-isolation --skip-dependency-check
}

# Tests require a running MySQL instance.
# check() {
#   cd ${pkgname#python-}-$pkgver
#   pytest
# }

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
