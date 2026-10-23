# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-stone
pkgver=3.5.5
pkgrel=1
pkgdesc='The Official API Spec Language for Dropbox API V2'
arch=(any)
url='https://github.com/dropbox/stone'
license=(MIT)
depends=(
  python
  python-packaging
  python-ply
  python-jinja
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(python-pytest)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('34ab840900755426ca823c637fcca9a3005c43c814b04b84ae22746dafb887634632d4227ae97274ed9500513e8135164dc60e649305e7d67c25a5791ad43381')
b2sums=('e87622fefafad51bc2796003fda9aea00b52b2b787f4b5906c449957ed847d6bb10dd4c67f35f91e88edef0b72ddcca4b430357fdf203ee68cf26e1c4af9f3c7')

build() {
  cd "$pkgname"

  export SETUPTOOLS_SCM_PRETEND_VERSION="$pkgver"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  pytest -vv
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
