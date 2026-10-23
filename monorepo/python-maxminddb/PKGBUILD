# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-maxminddb
pkgver=3.2.0
pkgrel=1
pkgdesc="Reader for the MaxMind DB format"
url="https://github.com/maxmind/MaxMind-DB-Reader-python"
license=('Apache-2.0')
arch=('x86_64')
depends=('python' 'libmaxminddb')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools'
             'python-setuptools-scm')
checkdepends=('python-pytest')
source=("git+https://github.com/maxmind/MaxMind-DB-Reader-python.git#tag=v$pkgver"
        "git+https://github.com/maxmind/MaxMind-DB.git"
        "git+https://github.com/maxmind/libmaxminddb.git")
sha512sums=('e2059cf7d2b51cddb5632eedce00a34ccb5e56be1e685b0cc1fa85c34c339ed1cd1535135bb25fff7bd63ee645d35ee2392c0fc0f0ffcc65038852e1c33ae6ed'
            'SKIP'
            'SKIP')

prepare() {
  cd MaxMind-DB-Reader-python
  git config submodule.tests/data.url "$srcdir"/MaxMind-DB
  git config submodule.extension/libmaxminddb.url "$srcdir"/libmaxminddb
  git -c protocol.file.allow=always submodule update --init
}

build() {
  cd MaxMind-DB-Reader-python
  python -m build -nw
}

check() {
  cd MaxMind-DB-Reader-python
  python -m installer -d tmp_install dist/*.whl
  local python_version=$(python -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')
  PYTHONPATH=tmp_install/usr/lib/python${python_version}/site-packages pytest
}

package() {
  cd MaxMind-DB-Reader-python
  python -m installer -d "$pkgdir" dist/*.whl
}
