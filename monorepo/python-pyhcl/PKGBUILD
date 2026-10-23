# Maintainer: Justin Kromlinger <hashworks@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: brent s. <bts[at]square-r00t[dot]net>

_pkgname=pyhcl
pkgname=python-pyhcl
pkgver=0.4.5
pkgrel=5
pkgdesc='HCL configuration parser for python'
arch=(any)
url='https://pypi.org/project/pyhcl/'
license=('MPL-2.0')
depends=(python)
makedepends=(git python-build python-installer python-setuptools python-wheel)
checkdepends=(python-pytest)
# setup.py needs a git repo to generate necessary version file
source=("git+https://github.com/virtuald/pyhcl.git#tag=${pkgver}")
sha256sums=('efdc12b78e3015a4980163ab7cb40bd7e542cc4ac8fc6265b5e5780e91f3fa21')

build() {
  cd $_pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pkgname
  PYTHONPATH="$PWD/src" pytest tests
}

package() {
  cd $_pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
}
