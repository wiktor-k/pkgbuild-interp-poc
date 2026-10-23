# Maintainer: David Runge <dvzrv@archlinux.org>

_name=dep-logic
pkgname=python-dep-logic
pkgver=0.7.2
pkgrel=1
pkgdesc="Python dependency specifications supporting logical operations"
arch=(any)
url="https://github.com/pdm-project/dep-logic"
license=(Apache-2.0)
depends=(
  python
  python-packaging
)
makedepends=(
  python-build
  python-installer
  python-pdm-backend
  python-wheel
)
checkdepends=(
  python-pytest
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('652aef5f6a5571b03abdac34bd3c51cc8a6baeff5d2e7945cc6323950e21af28ff00a4d36da05d88cdd455b6f5ff82f7e42316c9d8365b973289c579954dca33')
b2sums=('f38f6793d9a4ff5d500fb7e75ace25a1f94e9e6d7e3cbdfbe43fe8a87bb41a9bc7560dcfcc38eb42cea9f844d1ddfb8376685971dab4e57f1c82c3d238b46b23')

build() {
  cd $_name-$pkgver
  PDM_BUILD_SCM_VERSION=$pkgver python -m build --wheel --no-isolation
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
}
