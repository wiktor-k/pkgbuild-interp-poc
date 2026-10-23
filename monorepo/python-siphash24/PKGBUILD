# Maintainer: Jakub Klinkovský <lahwaacz at archlinux dot org>

_name=siphash24
pkgname=python-$_name
pkgver=1.9
pkgrel=1
pkgdesc="Streaming-capable SipHash implementation"
arch=(x86_64)
url="https://github.com/dnicolodi/python-siphash24"
license=("Apache-2.0 OR LGPL-2.1-or-later")
depends=(
  glibc
  python
)
makedepends=(
  cmake
  cython
  git
  meson
  meson-python
  python-build
  python-installer
  python-wheel
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
b2sums=('a0c6d0f6ea82bb8558d2a1687699b47ac75a295762bb1008adcdc3d721c77aa6ae810d7e5054bdf37e08a1b27cfa7eb83ef1fceec3ebd5fc5e6abcad8447371a')

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation -Cbuild-dir=build
}

check() {
  cd $pkgname-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m unittest discover -v
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
