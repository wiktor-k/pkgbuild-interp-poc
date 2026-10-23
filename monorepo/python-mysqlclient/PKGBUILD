# Maintainer: Bert Peters <bertptrs@archlinux.org>
# Contributor: Florian Pritz <bluewind@xinu.at>
# Contributor: Massimiliano Torromeo <massimiliano.torromeo@gmail.com>

pkgname=python-mysqlclient
_libname=${pkgname/python-/}
pkgver=2.3.0
pkgrel=1
pkgdesc="Fork of MySQL-python with support for Python 3"
arch=('x86_64')
url="https://github.com/PyMySQL/mysqlclient-python"
license=('GPL-2.0-or-later')
depends=('python' 'mariadb-libs')
makedepends=('mariadb' 'python-setuptools' 'python-installer' 'python-wheel' 'python-build')
checkdepends=('pifpaf' 'python-pytest' 'python-tblib')
source=("https://files.pythonhosted.org/packages/source/${_libname:0:1}/$_libname/$_libname-$pkgver.tar.gz")

build() {
    cd "$srcdir"/$_libname-$pkgver
    python -m build --wheel --no-isolation
}

check() {
    cd "$srcdir"/$_libname-$pkgver

    # Loosely based on upstream integration tests and haskell-hasql
    eval "$(pifpaf run mysql)"
    mariadb -S "$PIFPAF_MYSQL_SOCKET" -uroot -proot -e "CREATE DATABASE mysqldb_test"
    cat > tests/arch.cnf <<EOF
[MySQLdb-tests]
socket = $PIFPAF_MYSQL_SOCKET
user = root
database = mysqldb_test
password = root
default-character-set = utf8mb4
EOF

    python -m venv --system-site-packages test-env
    test-env/bin/python -m installer dist/*.whl

    TESTDB=arch.cnf test-env/bin/python -m pytest tests

    pifpaf_stop
}

package() {
    cd "$srcdir"/$_libname-$pkgver
    python -m installer --destdir="$pkgdir" dist/*.whl
}

sha256sums=('bea8294964266f6486f1ca514ccfcdbc54d4fe0d32882b38c1d4594df870be8b')
