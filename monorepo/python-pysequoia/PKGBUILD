# Maintainer: David Runge <dvzrv@archlinux.org>

_name=pysequoia
pkgname=python-pysequoia
pkgver=0.1.35
pkgrel=1
pkgdesc="Provides OpenPGP facilities in Python through the Sequoia PGP library"
arch=(x86_64)
url="https://github.com/wiktor-k/pysequoia"
license=(Apache-2.0)
groups=(sequoia)
depends=(
  libgcc
  glibc
  python
)
makedepends=(
  clang
  python-build
  python-installer
  python-maturin
  python-wheel
)
source=($_name-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz)
sha512sums=('fc37f820db79ec6d3b34fef275b871f21e86280639dda38c759cd8ebe10d92c1b4204d9a7c56815b0e4f5e9d8fcd87201f666052973523a1d69c895e2580772a')
b2sums=('a7e2a1be3960e7e0afa289157280363918c6459c8cb07819f9118c6b9b68f208137bce74888de35fec62c35cab1c59f164c927a531ef0d620f052991bd6413ff')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
