# Maintainer: Bruno Pagani <archange@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-mss
pkgver=10.2.0
pkgrel=1
pkgdesc="Ultra fast cross-platform multiple screenshots module"
arch=(any)
url="https://github.com/BoboTiG/python-mss"
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-hatchling
  python-installer
  python-wheel
)
checkdepends=(
  libxrandr
  lsof
  python-flaky
  python-pytest
  python-pytest-rerunfailures
  xorg-server-xvfb
)
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('1ce62de151d3f503c75effc23e90a7039852c612f55c1f6d7ad03b6cc8d5f7d5')

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname-$pkgver
  python -m installer --destdir=tmp_install dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  PYTHONPATH="$PWD/tmp_install/$site_packages" \
    xvfb-run -a -s "+extension RANDR +render -screen 0 1024x768x24" \
    pytest --override-ini="addopts=" \
    --deselect=src/tests/test_leaks.py:
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
