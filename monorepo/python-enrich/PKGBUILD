# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: David Runge <dvzrv@archlinux.org>

pkgname=python-enrich
_pkgname=${pkgname#python-}
pkgver=1.3.0
pkgrel=3
pkgdesc="Extends rich library functionality with a set of changes that were not accepted"
arch=(any)
url="https://github.com/pycontribs/enrich"
license=(MIT)
depends=(
  python
  python-rich
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-mock
  python-pytest-xdist
)
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")
sha512sums=('9fcb30383fb79ac06e97b00cc3c205c1e173749b06ab20f5a3cb09d65429e424de363168b0c4e18f06ba82ae980e23c8e0e725252a1f8e11a0a9682910984f74')

build() {
  cd "$_pkgname-$pkgver"
  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd "$_pkgname-$pkgver"
  python -m installer --destdir=tmp_install dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  PYTHONPATH="$PWD/tmp_install/$site_packages" pytest
}

package() {
  cd "$_pkgname-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" README.md
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
