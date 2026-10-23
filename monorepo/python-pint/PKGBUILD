# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Matthew Sloan matthew@sloan.cc

pkgname=python-pint
pkgver=0.26.1
pkgrel=1
pkgdesc="A unit library for Python"
arch=('any')
license=('BSD-3-Clause')
url="https://pint.readthedocs.org"
depends=(
  'python'
  'python-flexcache'
  'python-flexparser'
  'python-platformdirs'
  'python-typing_extensions'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-hatchling'
  'python-hatch-vcs'
)
checkdepends=(
  'python-numpy'
  'python-pytest'
  'python-pytest-benchmark'
  'python-pytest-subtests'
  'python-uncertainties'
)
source=("git+https://github.com/hgrecco/pint.git#tag=$pkgver")
sha512sums=('86aeec9758e9805f89b892e2f441cc5f262b2d88ca7075b006423b018918d3efac68a596ecb5cb042bc2caa5890b9110ca886dcdc0558d19c77b36d00ef75622')

build() {
  cd pint
  python -m build --wheel --no-isolation
}

check() {
  cd pint
  pytest
}

package() {
  cd pint
  python -m installer --destdir="$pkgdir" dist/*.whl

  # Symlink license file
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  install -d "$pkgdir"/usr/share/licenses/$pkgname
  ln -s "$site_packages"/pint-$pkgver.dist-info/licenses/LICENSE \
    "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}
