# Maintainer: David Runge <dvzrv@archlinux.org>

_name=cytoolz
pkgname=python-cytoolz
pkgver=1.1.0
pkgrel=2
pkgdesc="Cython implementation of Toolz: High performance functional utilities"
arch=(x86_64)
url="https://github.com/pytoolz/cytoolz/"
license=(BSD-3-Clause)
depends=(
  glibc
  python
  python-toolz
)
makedepends=(
  cython
  python-build
  python-installer
  python-setuptools
  python-setuptools-git-versioning
  python-wheel
)
checkdepends=(python-pytest)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
b2sums=('21628828012aae90d3a73a298b3415c461072be85cb510ad121ffcbfda3bc4c9dcbb7a32a0c04806e5700d6afdb264677d714bc61183f1278838ab415e02df62')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH"
  pytest -vv test_dir/$_site_packages
}

package() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -vDm 644 LICENSE.txt -t "${pkgdir}/usr/share/licenses/${pkgname}"
  rm -frv "$pkgdir/$_site_packages/$_name/tests"
}
