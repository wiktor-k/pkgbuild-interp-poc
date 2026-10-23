# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Blackleg <blacklg@openaliasbox.org>
# Contributor: Nidhogg
# Contributor: juantascon

pkgname=python-peewee
pkgver=4.5.2
pkgrel=1
pkgdesc="A small, expressive orm"
arch=('x86_64')
url="https://github.com/coleifer/peewee"
license=('MIT')
depends=(
  'glibc'
  'python'
  'sqlite'
)
makedepends=(
  'cython'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-aiosqlite'
  'python-apsw'
  'python-flask'
  'python-gevent'
  'python-greenlet'
  'python-psycopg2'
)
optdepends=(
  'python-aiomysql: for async MySQL database support'
  'python-aiosqlite: for async SQLite database support'
  'python-apsw: for APSW SQLite database support'
  'python-asyncpg: for async PostgreSQL database support'
  # 'python-cysqlite: for Cython SQLite database support' # not packaged
  'python-flask: for Flask integration utilities'
  'python-gevent: for green-thread based async support'
  'python-greenlet: for async database support'
  'python-psycopg2: for PostgreSQL database support'
  'python-psycopg: for PostgreSQL database support via psycopg3'
  'python-pymysql: for MySQL database support'
)
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('443adb5429b9a8ca99dddef2b198153a16e5c0fe781f037bc02929658c6832d1057d6283e823a60f740268dec3839ec6772e0839a603345b19615f40dc2fe780')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  python -m unittest discover -v
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
