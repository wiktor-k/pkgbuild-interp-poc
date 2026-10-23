# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

_pyname=pysmt
pkgname=python-${_pyname}
pkgver=0.9.6
pkgrel=2
pkgdesc='Solver-agnostic library for SMT Formulae manipulation and solving'
url='https://github.com/pysmt/pysmt'
license=('Apache-2.0')
arch=('any')
depends=('python' 'python-z3-solver')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
checkdepends=('python-pytest')
source=(
  "${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz"
)
sha512sums=('499e1df2c3d9c7a550a02aa0f05fed90b0f8bb3f74f155f9466e11779d81d2e024d26a0dcfa4e0cc6b6b8d57ef7f99a66ae7dd7b7449b9d226d082c672416379')
b2sums=('3ccfc4264acb1a5958c2cccbf77fcdc31b1b4eb22a0d3bafa9666a3edc2d22d6ae0f6ed74aaa79d5a2816c40cfea7a54f170202aa02980cf4f23770085d9cb3a')

build() {
  cd ${_pyname}-${pkgver}
  export EXTRA_CFLAGS="${CFLAGS} ${CPPFLAGS}"
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pyname}-${pkgver}
  PYTHONPATH=build/lib pytest -k 'not test_z3_back_formulae'
}

package() {
  cd ${_pyname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 README.rst -t "${pkgdir}/usr/share/doc/${pkgname}"
  cp -r examples "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
