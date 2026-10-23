# Maintainer: Andrzej Giniewicz <gginiu@gmail.com>
# Contributor: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-xlsxwriter
pkgver=3.3.0
pkgrel=1
pkgdesc='A Python module for creating Excel XLSX files'
arch=(any)
url='https://github.com/jmcnamara/XlsxWriter'
license=(BSD-2-Clause)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-wheel
  python-setuptools
)
source=("$pkgname::git+$url#tag=RELEASE_$pkgver")
sha512sums=('e39afaa3f9099a3542efe7eb82b4c8b6526da0072975019e789c2d7a7b6ea60c3e1d4a0a2c5e09452cc0a380270f40ff163baa562e47eed5981ac29a97664b7e')
b2sums=('d06277cb2f1cf3fcd3522e007b0f81a9bd1376efd57426616a61151e15d65234b624183e419f71ff6905be94e7b0afd62e4465166987c0f3e5f7ee521b2a5a1b')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  python -m unittest discover
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
