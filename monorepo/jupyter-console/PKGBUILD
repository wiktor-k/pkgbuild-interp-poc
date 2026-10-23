# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Kyle Keen <keenerd@gmail.com>

_pyname=jupyter_console
pkgname=${_pyname/_/-}
pkgver=6.6.3
pkgrel=5
pkgdesc='An IPython-like terminal frontend for Jupyter kernels in any language'
arch=(any)
url='https://pypi.org/project/jupyter_console/'
license=(BSD-3-Clause)
depends=(python
         python-jupyter-client)
makedepends=(git
             python-build
             python-hatchling
             python-installer)
checkdepends=(python-flaky
              python-ipykernel
              python-pytest)
source=(git+https://github.com/jupyter/jupyter_console#tag=v$pkgver)
sha256sums=('c696f4350818e7858de62e43e25f15fec43c4b54f0c7ff4f8d7c78280a1c25ef')

build() {
  cd $_pyname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pyname
  pytest -v
}

package() {
  cd $_pyname
  python -m installer --destdir="$pkgdir" dist/*.whl
  rm -r "$pkgdir"/usr/lib/python*/site-packages/$_pyname/tests
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
