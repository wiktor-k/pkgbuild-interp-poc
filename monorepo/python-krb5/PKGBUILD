# Maintainer: Bruno Pagani <archange@archlinux.org>

_pkg=krb5
pkgname=python-${_pkg}
pkgver=0.9.0
pkgrel=1
pkgdesc="krb5 API interface"
arch=(x86_64)
url="https://github.com/jborean93/pykrb5"
license=(MIT)
depends=(python krb5)
makedepends=(cython python-build python-installer python-wheel python-setuptools)
checkdepends=(python-pytest python-k5test)
# No tests in pypi tarball
#source=(https://files.pythonhosted.org/packages/source/${_pkg::1}/${_pkg}/${_pkg}-${pkgver}.tar.gz)
source=(${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha256sums=('b60226a9346eb8c42f6164b81fdd59cd1b0fb8a29d3e008266c52938f5344a8f')

prepare() {
  cd py${_pkg}-${pkgver}
  sed -e '/Cython ==/d' -i pyproject.toml # Drop pinned cython dependency
}

build() {
  cd py${_pkg}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd py${_pkg}-${pkgver}
  local python_version=$(python -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')
  PYTHONPATH="${PWD}"/build/lib.linux-$CARCH-cpython-${python_version/./} pytest -vv --color=yes
}

package() {
  cd py${_pkg}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "${pkgdir}"/usr/share/licenses/${pkgname}
}
