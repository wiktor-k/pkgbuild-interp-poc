# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-pywayland
_name=${pkgname#python-}
pkgver=0.4.19
pkgrel=1
pkgdesc="Python bindings for the libwayland library"
arch=(x86_64)
url="https://github.com/flacjacket/pywayland"
license=(Apache-2.0)
depends=(
  glibc
  python
  python-cffi
  wayland
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
  wayland-protocols
)
checkdepends=(python-pytest)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('697f02089128777d5e195bd0dc8dd91f02e20f89029c76bbd7e1d6ff3ddc6d8e811b9a3d5c8b9fc359e3a26088d01bdd550439b73c3340e5c2d49bf089fcfb32')
b2sums=('84d767c50de00eaabaaebd1042a8dd6ddef86c6fc006665af39c4c35f805fdff3f0ffd4d003e84e7c8d305550ab56c8b79cd3f47f09faec6198fbd68e1185d00')

build() {
  cd $_name-$pkgver
  python pywayland/ffi_build.py
  python -m pywayland.scanner
  python -m build --wheel --no-isolation
}

check() {
  local _py_ver="$(python -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')"

  cd $_name-$pkgver
  mkdir -vp temp
  export PYTHONPATH="build/lib.linux-$CARCH-$_py_ver:$PYTHONPATH"
  export XDG_RUNTIME_DIR="$PWD/temp"
  pytest -vv
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname"
}
