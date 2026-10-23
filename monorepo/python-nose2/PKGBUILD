# Maintainer: David Runge <dvzrv@archlinux.org>

_name=nose2
pkgname=python-nose2
pkgver=0.16.0
pkgrel=1
pkgdesc="The next generation of nicer testing for python"
arch=(any)
url="https://github.com/nose-devs/nose2"
license=(
  BSD-2-Clause
  BSD-3-Clause
)
depends=(
  python
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-coverage
)
optdepends=(
  'python-coverage: for coverage_plugin'
  'python-setuptools: for eggdiscovery plugin'
  'python-sphinx: for doc'
  'python-sphinx_rtd_theme: for doc'
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('0991e3ad18b04c1ecbb909270bfaa37b1bba38e31fd33753cff2d81f485fc8e11e13a0305a84bd76052a90b084b4f2d4b8ca2497015398d7d0232d68a1a26235')
b2sums=('7538a914ece48bfb0deb08d8e39f337e2739939d17231f21285896c4cbf837b1bd9432f48e068b43c9215456dedee2c43cfbd6b30ad183e4311ec55671551254')

prepare() {
  sed -n '1,27p' $_name-$pkgver/LICENSE > BSD-2-Clause.txt
  sed -n '34,65p' $_name-$pkgver/LICENSE > BSD-3-Clause.txt
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location to be able to run tests with nose2 itself
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$site_packages:$PYTHONPATH"
  export PATH="$PWD/test_dir/usr/bin:$PATH"
  nose2 -v --start-dir test_dir/$site_packages/$_name/tests/unit
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 ../BSD-{2,3}-Clause.txt -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 {AUTHORS,README.rst} -t "$pkgdir/usr/share/doc/$pkgname/"
}
