# Maintainer: David Runge <dvzrv@archlinux.org>

_name=flufl.i18n
pkgname=python-flufl.i18n
pkgver=8.0.0
pkgrel=1
pkgdesc="A high level API for internationalization"
arch=(any)
url="https://gitlab.com/warsaw/flufl.i18n"
license=(Apache-2.0)
depends=(
  python
  python-atpublic
)
makedepends=(
  python-build
  python-hatchling
  python-installer
)
checkdepends=(
  python-sybil
  python-pytest
)
provides=(python-flufl-i18n)
replaces=(python-flufl-i18n)
source=($url/-/archive/$pkgver/$_name-$pkgver.tar.gz)
sha512sums=('c5724a050091832b10a492861d728481bba51cc33884a1b59fc5887c32441b5352c7c3c67d33e64af4de5d7345f50f56112276dc51ed900f11603d9f92d4b9d7')
b2sums=('ae5f60316f4ae670fe3fc2aff9d3da9ea9c78ab3000cef568451a50e8a31f25afd61c11c46d10922d395b402fe5c373865456e95fb99d41b612387f5f6a4f500')

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
  pytest -vv -o addopts=''
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
