# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-rcssmin
_name="${pkgname#python-}"
pkgver=1.3.0
pkgrel=1
pkgdesc="Fast CSS minifier for Python"
arch=(x86_64)
url="https://github.com/ndparker/rcssmin"
license=(Apache-2.0)
depends=(
  glibc
  python
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('55126afafb8a73a26312983153ef15a2f3db51cb553efc62943ce3e96dffb2514c6a172ded888bed8148e06bba41ac419723cf71a1034a576f9ab7010b80d8d2')
b2sums=('517ffaf08acc90266ef2d89ac68db8f7ab806c52b473f4f86ebaad26aedc8a265b88874b40fa85f2bed3ae21074e4367c596109afd993696aa3c15885c02e40e')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH"
  pytest -vv
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
