# Maintainer: David Runge <dvzrv@archlinux.org>
# Contributor: Daniel M. Capella <polyzen@archlinux.org>

_name=tzdata
pkgname=python-tzdata
pkgver=2026.5
pkgrel=1
pkgdesc='Provider of IANA time zone data'
arch=(any)
url=https://github.com/python/tzdata
license=(Apache-2.0)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-subtests
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('587427fac3512bd5f9976606373fcefb2eb7ad6d8fa1a31eb92b3872b761dae99ced778f5bb8418c20c9b8461d9c76ee0a9f8e8e84cead2a4e3eca6115f1fbe0')
b2sums=('c2d66489b27e73f91c9da2aec933478f6b8d35b5c720919777d1d85cea9aa5d6634a6501914c35ef8cf00273102444c2b0fcb5022375d4688ab4ba0988d76c4e')

build() {
  cd $_name-$pkgver
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  local pytest_options=(
    -vv
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH"
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl

  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  ln -s /etc/localtime "$pkgdir/$site_packages/$_name/zoneinfo/localtime"
}
