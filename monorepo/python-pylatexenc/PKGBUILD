# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Asger Hautop Drewsen <asger@tyilo.com>

pkgname=python-pylatexenc
_pkgname=${pkgname#python-}
pkgver=2.11
pkgrel=1
pkgdesc="Simple LaTeX parser providing latex-to-unicode and unicode-to-latex conversion"
arch=(any)
url="https://github.com/phfaist/pylatexenc"
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")
sha256sums=('5f622ef586dbffeb8ccd6ac210431ff0acf41d74275de179696e38006c28b862')

build() {
  cd "$_pkgname-$pkgver"
  python -m build --wheel --no-isolation
}

check() {
  cd "$_pkgname-$pkgver"
  python -m installer --destdir=tmp_install dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$PWD/tmp_install/$site_packages"
  pytest
}

package() {
  cd "$_pkgname-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
