# Maintainer: David Runge <dvzrv@archlinux.org>

_name=flufl.lock
pkgname=python-flufl-lock
pkgver=10.0.0
pkgrel=1
pkgdesc="NFS-safe file locking with timeouts for POSIX systems for Python"
arch=(any)
url="https://gitlab.com/warsaw/flufl.lock"
license=(Apache-2.0)
depends=(
  python
  python-atpublic
  python-psutil
)
makedepends=(
  python-build
  python-hatchling
  python-installer
)
checkdepends=(
  python-pytest
  python-sybil
)
source=($url/-/archive/$pkgver/$_name-$pkgver.tar.gz)
sha512sums=('d03c4d286c14766ff93715246c553db37f1fc140a7d6cab98bf44a31bbf964701e6ffd33b08e2b9a8b9ce8b800e664295e604bfb09187c35057ba44a8420df58')
b2sums=('3d89f9cb52b804929ee134d1d4002aef09096d4db685eed6ef8a1bd381f31a3e3e38ada86acccf641dd2c9ba3867c91110adddd4e4168fd0ecc66642f548fdf4')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  export PYTHONPATH="src:$PYTHONPATH"
  pytest -vv -c /dev/null
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
