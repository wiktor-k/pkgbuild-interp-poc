# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Modified from aur/python2-pytrie; original maintainership info:
# Contributor: dseg

pkgname=python-pytrie
pkgver=0.4.0
pkgrel=10
pkgdesc='A pure Python implementation of the trie data structure'
arch=(any)
url='https://github.com/gsakkis/pytrie'
license=(BSD-3-Clause)
depends=(
  python
  python-sortedcontainers
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
# Tests in this package imports tests from CPython :/
checkdepends=(python-tests)
source=(
  "$pkgname-$pkgver.tar.gz::https://files.pythonhosted.org/packages/source/P/PyTrie/PyTrie-$pkgver.tar.gz"
  python-3.14.patch
)
sha512sums=('c951568774c1e35d2b2bd5a52e5368da5ef2ffc83476d60882aa842debcd53c5eb0ed19f497a09f78c6a736ffba57702eda3aa0f3afa74dd0d1e770c17885c74'
            '5fa5870caf0eda7a398a2953c91be67b8d3afa739aa20a739cef0271863936e88a1f271eb8544f4f2c82328ad60f2ed9dbfc7d0c4eb50eeb0f5e355d04dd7136')
b2sums=('de57b4c36429c3e2d9ec7cb89e6515fd0a131e890bbfec17797a9a03b932ca2370cf283c67b4a99bb742b275c6f005789af410b875e857d30e6ba9e959403899'
        '80b5decb13eb33efa127c80d867340ca160263c2658c0d99fa9d548809b0fdd8a8f5adf3ffc43e15bb4383fcfc09ff2a424839b0bf00e5f3280e49e45cdcb5e9')

prepare() {
  cd "PyTrie-$pkgver"

  patch -p1 -i "$srcdir/python-3.14.patch"
}

build() {
  cd "PyTrie-$pkgver"

  python -m build --wheel --no-isolation
}

check() {
  cd "PyTrie-$pkgver"

  python -m unittest discover -v ./tests
}

package() {
  cd "PyTrie-$pkgver"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
