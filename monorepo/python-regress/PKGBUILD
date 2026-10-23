# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-regress
_name=regress
pkgver=2026.9.1
pkgrel=1
pkgdesc="Python bindings to the Rust regress crate which provides ECMA regular expressions"
arch=(x86_64)
url="https://github.com/crate-py/regress"
license=(MIT)
depends=(
  libgcc
  glibc
  python
)
makedepends=(
  python-build
  python-installer
  python-maturin
  python-wheel
)
checkdepends=(python-pytest)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('a8836710a2661e61dab02ca64e946ddb4605d8210946ff65ef9974a87460f116797a009fa977541ff473819f52613fa50319c1fb54d66d30360fd3c1bfe72c41')
b2sums=('ccd3f388d4bbd2b76640e2d7bdc7632b30f2dca135d2a5bbd8bd273db4f772ab3ca94ece3aa14f2c90777ca598aef277d7c747bfe8ee2ace4b48e01944861eb9')

prepare() {
  cd $_name-$pkgver
  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  local python_build_options=(
    --wheel
    --no-isolation
    -C="--target=$(rustc --print host-tuple)"
    -C="--locked"
    -C="--release"
    -C="--all-features"
    -C="--strip"
  )

  cd $_name-$pkgver
  python -m build "${python_build_options[@]}"
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
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
