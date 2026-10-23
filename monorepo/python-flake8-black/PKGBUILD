# Maintainer: George Rawlinson <george@rawlinson.net.nz>
# Contributor: Daniel M. Capella <polyzen@archlinux.org>

pkgname=python-flake8-black
pkgver=4.0.0
pkgrel=2
pkgdesc='Flake8 plugin to call black as a code style validator'
arch=(any)
url=https://github.com/peterjc/flake8-black
license=(MIT)
depends=(
  python
  flake8
  python-black
)
makedepends=(
  git
  python-build
  python-installer
  python-wheel
  python-setuptools
)
source=(
  "$pkgname::git+$url.git#tag=v$pkgver"
  pyproject-no-pip.patch
)
sha512sums=('6d39ab427c0d01c96682091b5c7d7f4f4b27867311f060372de14522b860550c6e971e733bd6651b06f63091e11cf2d84dc654b37920c0c096b35cf303a7b0b5'
            '0dc458646781a621b571ac6a2ef892da242fa353b0be3e63a189ff0da8d0d10495521ff559072e50ada62e3565e43103cf605766b4cdf9f1b97a039a1ebd03fd')
b2sums=('83de3590352308fe1ac4742c1576fa04d04e38bbacec01d662ed1c7fd7e3a0bc4f7876d72678617861503e9012998ed47550553076fce2819f127a6641cd33ec'
        'abc8840cf1cbf6b2386c38e690dcb479671cf2df64c584787f7ebb06c31dc3cec1d1d407e3bf3173460f31f8a61023bd1609e20fcac9317f67a3c18c610b2b19')

prepare() {
  cd "$pkgname"

  patch -p1 -i "$srcdir/pyproject-no-pip.patch"
}

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

# Requires black==23.9.1 for tests
# https://github.com/peterjc/flake8-black/commit/8abe1ae59ee2b37b5c790152a7c7ba44118ebb8e
#check() {
#  cd "$pkgname"
#
#  python -m venv --system-site-packages test-env
#  source test-env/bin/activate
#  python -m installer dist/*.whl
#  cd tests
#  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
#  PYTHONPATH="../build:$site_packages" ./run_tests.sh
#  deactivate
#}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.rst
}
