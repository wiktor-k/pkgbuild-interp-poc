# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>

pkgname=python-plumbum
_name=plumbum
pkgver=2.0.2
pkgrel=1
pkgdesc='A small yet feature-rich Python library for shell script-like programs, and more'
url='https://github.com/tomerfiliba/plumbum'
arch=('any')
license=('MIT')
depends=(
  'python'
)
makedepends=(
  'python-build'
  'python-hatch-vcs'
  'python-hatchling'
  'python-installer'
)
source=(https://files.pythonhosted.org/packages/source/${_name::1}/$_name/$_name-$pkgver.tar.gz)
sha512sums=('49a4234ddc6e94289dff23ee0186dffb87b10500c30ebf2cf9867ff9487ba08a19e12dce5e0740ddb90aebe9f42f7283f4cd8aabaa79d95fcb969bce0dc8de24')
b2sums=('c52ea1f20f899dbe89c1b2499510fdee3082739f9ac4f44f2d08c8fbba57a7d8e2ebfd64b1a459db853a94c612f370fc8b75f1788853f8de3989cf5af7aa6e48')

build() {
  cd "${_name}-${pkgver}"
  python -m build --wheel --no-isolation
}

check() {
  cd "${_name}-${pkgver}"
  PYTHONPATH=. python -c 'import plumbum'
}

package() {
  cd "${_name}-${pkgver}"
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -D LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim:set ts=2 sw=2 et:
