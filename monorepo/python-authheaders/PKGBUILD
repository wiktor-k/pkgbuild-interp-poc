# Maintainer: David Runge <dvzrv@archlinux.org>

_name=authheaders
pkgname=python-authheaders
pkgver=0.16.3
# last updated as per https://www.psddmarc.org/registry.html
_version=2023.09.02
pkgrel=4
pkgdesc="Library for the generation of email authentication headers"
arch=(any)
url="https://github.com/ValiMail/authentication-headers"
license=(
  MPL-2.0
  ZPL-2.1
  Zlib
)
depends=(
  publicsuffix-list
  python
  python-authres
  python-dkim
  python-dnspython
  python-publicsuffix2
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=(
  $url/releases/download/$pkgver/$_name-$pkgver.tar.gz{,.asc}
)
sha512sums=('94741ec6127aa20b128958ba5b386b63d00f27ae68fe6b5ede3a0308bd6a75a9b940fb40b6066b65fefb07e40e85cb412bd59fde6c76f60099d6bf6872ae9a68'
            'SKIP')
b2sums=('26e9f9d188da2483132d5a86b1b4c93181b1524c3950a521d52eff279ebad42b2672a756317cf76b4f8667b65fabd99d77fa97c34124aecdf87829b8505a1c8a'
        'SKIP')
validpgpkeys=('E7729BFFBE85400FEEEE23B178D7DEFB9AD59AF1') # Donald Scott Kitterman <scott@kitterman.com>

prepare() {
  cd $_name-$pkgver
  # add empty file, because upstream is gone: https://github.com/ValiMail/authentication-headers/issues/35
  touch $_name/psddmarc.csv
  printf 'location = "/usr/share/publicsuffix/public_suffix_list.dat"\n' > $_name/findpsl.py
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH"
  pytest -vv
}

package() {
  cd $_name-$pkgver
  local python_stdlib_basepath="$(python -c "from sysconfig import get_path; print(get_path('stdlib'))")"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGES,README.md} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 COPYING -t "$pkgdir/usr/share/licenses/$pkgname/"
}
