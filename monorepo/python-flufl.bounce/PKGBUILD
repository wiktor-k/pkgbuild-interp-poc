# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-flufl.bounce
_name="${pkgname#python-}"
pkgver=6.0.0
pkgrel=1
pkgdesc="Email bounce detectors"
arch=(any)
url="https://gitlab.com/warsaw/flufl.bounce"
license=(Apache-2.0)
depends=(
  python
  python-atpublic
)
makedepends=(
  python-build
  python-installer
  python-hatchling
)
checkdepends=(
  python-pytest
  python-sybil
)
provides=(python-flufl-bounce)
replaces=(python-flufl-bounce)
source=($url/-/archive/$pkgver/$_name-$pkgver.tar.gz)
sha512sums=('33d4898aa456549af6f1c1d52475128abd7faaff7c3632cf19b43e882fd9bccf9f9a5db87b7a7a9f95367083fc39c29a9995efd9ae1809e00aac8923560a1e04')
b2sums=('c05a331a742bc3b91031d35f0e18172be0fe5aedcb9e8893446467c473063a6451bec29cc2ad8b20b2146ece5c75f19b0525373afb3accba741f49af037cafc9')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
  )

  cd $_name-$pkgver
  PYTHONPATH="$PWD/src:$PYTHONPATH" pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
