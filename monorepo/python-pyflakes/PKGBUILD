# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Karol 'Kenji Takahashi' Woźniak <kenji.sx>
# Contributor: Tianjiao Yin <ytj000+AUR@gmail.com>
# Contributor: Thomas Dziedzic < gostrc at gmail >
# Contributor: TDY <tdy@gmx.com>
# Contributor: Tiago Pierezan Camargo <tcamargo@gmail.com>

pkgname=python-pyflakes
pkgver=4.0.3
pkgrel=1
pkgdesc='A lint-like tool for Python to identify common errors quickly without executing code'
arch=('any')
url='https://github.com/PyCQA/pyflakes'
license=('MIT')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz")
sha512sums=('286ef5c011eae75911f8460dc3666c5148aac5898d702c5d0428ec3c6ad4f944778a4c82b1927dd49b4652a46b16e5486513e9d819a02e6c986d471201574505')

build() {
  cd pyflakes-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd pyflakes-$pkgver
  # Disable failing test with Python 3.13 https://github.com/PyCQA/pyflakes/issues/812
  sed -i 's/test_errors_syntax/xtest_errors_syntax/' pyflakes/test/test_api.py
  python -m unittest discover pyflakes
}

package() {
  cd pyflakes-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl

  # We have python 3 as default python, and want to keep compatibility with the old pyflakes3k naming
  ln -s pyflakes "$pkgdir/usr/bin/pyflakes3k"

  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
