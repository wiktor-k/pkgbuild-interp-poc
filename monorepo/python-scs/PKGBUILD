# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=python-scs
pkgver=3.3.1
pkgrel=1
pkgdesc='Splitting Conic Solver'
url='https://github.com/bodono/scs-python'
license=(MIT)
arch=(x86_64)
depends=(blas
         glibc
         lapack
         python
         python-scipy)
makedepends=(git
             meson-python
             python-build
             python-installer)
checkdepends=(python-pytest)
source=(git+https://github.com/bodono/scs-python#tag=$pkgver
        git+https://github.com/cvxgrp/scs)
sha256sums=('d31ab1755dfd058e14f5b3b8a31fe4a02c5e77a294af478746fcb976d012c694'
            'SKIP')

prepare() {
  cd scs-python
  git submodule init
  git submodule set-url scs_source "$srcdir"/scs
  git -c protocol.file.allow=always submodule update
}

build() {
  cd scs-python
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd scs-python
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest -v
}

package() {
  cd scs-python
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
