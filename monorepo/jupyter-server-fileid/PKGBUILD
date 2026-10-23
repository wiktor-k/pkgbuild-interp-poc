# Maintainer: Antonio Rojas <arojas@archlinux.org>

_pyname=jupyter_server_fileid
pkgname=jupyter-server-fileid
pkgver=0.10.0
pkgrel=1
pkgdesc='A Jupyter Server extension providing an implementation of the File ID service'
arch=(any)
url='https://jupyter.org/'
license=(BSD-3-Clause)
depends=(jupyter-server
         python
         python-click
         python-jupyter-core
         python-jupyter-events
         python-tornado
         python-traitlets)
makedepends=(git
             python-build
             python-hatchling
             python-installer)
checkdepends=(python-pytest-jupyter)
source=(git+https://github.com/jupyter-server/$_pyname#tag=v$pkgver)
sha256sums=('cd255cb3fc03441af5eb28179ddf41c432b6e0402b48c8dfae873ea3074a3302')

build() {
  cd $_pyname
  python -m build -wn
}

check() {
  cd $_pyname
# https://github.com/jupyter-server/jupyter_server_fileid/issues/58
  pytest -v -k 'not test_get_path_oob_move_nested and not test_get_path_oob_move_deeply_nested'
}

package() {
  cd $_pyname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
  mv "$pkgdir"/{usr/,}etc
}
