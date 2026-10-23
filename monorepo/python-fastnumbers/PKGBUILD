# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>

_name=fastnumbers
pkgname=python-fastnumbers
pkgver=5.2.0
pkgrel=1
pkgdesc="Super-fast and clean conversions to numbers"
arch=(x86_64)
url="https://github.com/SethMMorton/fastnumbers"
license=(MIT)
depends=(
  glibc
  libgcc
  libstdc++
  python
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-hypothesis
  python-numpy
  python-pytest
  python-typing_extensions
)
optdepends=(
  'python-numpy: for conversion to numpy arrays'
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('f2a0c60c421d73b45208526ac2794093e8c97942815d9c551aed4b692e7b4f5501bb92ce67670684fb7c14f89fe602704c73491a9b8bed2fab928e26f36cb392')
b2sums=('94a53d1e57d3b9b21060866f9712878361296c5a40295b94a5914ae68a7d7b21c17f9360839696b5e99d6933c73eb16da366549cc5f11580350b2f027693b0af')

build() {
  cd $_name-$pkgver
  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$site_packages:$PYTHONPATH"
  pytest -vv
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
