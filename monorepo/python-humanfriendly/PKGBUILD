# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-humanfriendly
pkgver=10.0
pkgrel=10
pkgdesc="Human friendly input/output for text interfaces using Python"
arch=(any)
url="https://github.com/xolox/python-humanfriendly"
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-docutils
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz
	pipes-removal-python3-13.patch)
sha512sums=('a0f324aa3af27088b77bd68c1d2ee65090ca6ace71414ae698452bc2a81de82f82f8d53b47145ea9822b93e600da3b871e9244a9681d06b54fa2c31efc2857ef'
            'def7329a11d01854e0521845efd2f6636cd5b221346dc9f14a67e09d26ede1dd90a58ed3cc1ca0de1dea4c2d0431ad7329759dd5fc3bde7445b383f94323ff3f')
b2sums=('0e133bc4f871455f9aa20f542a9a5be83a47db07865b44cb53f2ba093425925404ab91d5019322a9a25448362b83cc040bdf1502c200022dd0a23c423799ee8f'
        '70e7b5ea895b012ac03ee06ada29d184b94d60be548c8495d352e108a2236edc6a57dc30fc8ecd5e2b1f39a72dad1a5c450b56e8dd1d20595bc2908a85f6e56e')

prepare() {
  # replace use of mock with unittest.mock
  # disable tests requiring even more dead capturer project
  sed -e 's/from mock import/from unittest.mock import/;s/test_terminal_cap/off_test_terminal_cap/; s/test_prompt_for_conf/off_test_prompt_for_conf/; s/test_generate_warning/off_test_generate_warning/' -i $pkgname-$pkgver/${pkgname#python-}/tests.py

  # Python 3.13 compatibility https://github.com/xolox/python-humanfriendly/pull/75
  cd $pkgname-$pkgver
  patch -Np1 -i ${srcdir}/pipes-removal-python3-13.patch
}

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname-$pkgver
  python -m unittest discover -v ${pkgname#python-}/
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGELOG,README}.rst -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE.txt -t "$pkgdir/usr/share/licenses/$pkgname/"
}
