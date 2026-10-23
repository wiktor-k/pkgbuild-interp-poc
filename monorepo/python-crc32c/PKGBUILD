# Maintainer: David Runge <dvzrv@archlinux.org>

_name=crc32c
pkgname=python-crc32c
pkgver=2.9.post0
pkgrel=1
pkgdesc="A python package implementing the crc32c algorithm in hardware and software"
arch=(x86_64)
url="https://github.com/ICRAR/crc32c"
license=(
  BSD-3-Clause
  LGPL-2.1-or-later
)
depends=(
  glibc
  python
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('f1731e64168a875c1abad2145b4c53a897f9ef46c00c7a0a4f1255f61c7483048561bcecec1009a248b8ffe8c2bbe065a57895f19000b7d9aac5427614330331')
b2sums=('65d6bccd289e70ebccb8a75dc3b1cc7cab18db7b752ea66ea3ba47e684abb0e4adb1e934c2b856ff180bc221aeec54599166e0458fe7a53d155a6f47385220ce')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH" pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE.* -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
