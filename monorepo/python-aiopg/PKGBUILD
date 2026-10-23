# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Maintainer: Filipe Laíns (FFY00) <lains@archlinux.org>

pkgname=python-aiopg
pkgver=1.4.0
pkgrel=7
pkgdesc="Library for accessing a PostgreSQL database from the asyncio"
arch=(any)
url="https://github.com/aio-libs/aiopg"
license=(BSD-2-Clause)
depends=(
  python
  python-psycopg2
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
# checkdepends=(
#   python-docker
#   python-pytest
#   python-sqlalchemy
# )
optdepends=('python-sqlalchemy: SQLAlchemy support')
source=(
  "$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz"
  "$pkgname-remove-dep-async-timeout.patch::$url/commit/b5d4a510a0de774085d7bf29297c5abb4ff3b9bc.patch"
)
sha256sums=('1b1c2ae3449640fa8fbaa95fe893c83df6ee5195b7018ce5c1b4427f95979509'
            '3d8d5f3d8c097b5008b9d938459673ca75d0bab566b1e8308e9ed1298f77cc59')

prepare() {
  cd ${pkgname#python-}-$pkgver
  patch -Np1 -F3 < ../$pkgname-remove-dep-async-timeout.patch
}

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

# All tests require a running PostgreSQL Docker container
# check() {
#   cd ${pkgname#python-}-$pkgver
#   python -m installer -d tmp_install dist/*.whl
#   local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
#   PYTHONPATH="$PWD/tmp_install/$site_packages" pytest
# }

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
