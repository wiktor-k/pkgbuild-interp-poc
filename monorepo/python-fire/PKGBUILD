# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-fire
_name=${pkgname#python-}
pkgver=0.7.1
pkgrel=2
pkgdesc="A library for automatically generating command line interfaces"
arch=(any)
url="https://github.com/google/python-fire"
license=(Apache-2.0)
depends=(
  python
  python-termcolor
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-hypothesis
  python-levenshtein
  python-pytest
)
source=(
  $pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz
  fix-missing-eventloop.patch
)
sha512sums=('5a54cb86b4a75560132b36507047530aab95666102095410f0e26d4341799d2809c1959f3edb963cf7f78d774c8be097cc2d4761e845fe101e9a369677c1e84c'
            'd79bf73c1f50342e20e28a15e48268ad5106db1b89288fe8b7561f494b177b78595ea4a9811a9503183eeaaf7b271d0bdcf1d083af5c7d5e3acee5ddd7e07b22')
b2sums=('207afbb87a838d9def4f88f8e030d39b7f457018f8ff13ff661763d6bc8b70d62a95ba8ca2cf149e95e9b1134adfd7367b356bd63fc3b4012682961184df0f2e'
        '824f6f835e3dd9cca5dda6eb87b6bc77ea7095f56655c1375cb6080df3f3ed9d5782420782ce9653720c2c6e349ac268738ece673337b92539337da5a3c88495')

prepare() {
  cd $pkgname-$pkgver
  patch -p1 -i ../fix-missing-eventloop.patch
}

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname-$pkgver
  export PYTHONPATH="build:$PYTHONPATH"
  pytest -v
}

package() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
  # remove test files: https://github.com/google/python-fire/issues/422
  rm -frv "$pkgdir/$_site_packages/$_name/"{testutils,*_test}.py
}
