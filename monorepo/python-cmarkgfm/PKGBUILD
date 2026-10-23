# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>

pkgname=python-cmarkgfm
_pyname=${pkgname#python-}
pkgver=2025.10.22
pkgrel=2
pkgdesc="Minimal bindings to GitHub's fork of cmark"
url="https://github.com/theacodes/cmarkgfm"
license=('MIT')
arch=('x86_64')
depends=('cmark-gfm' 'python-cffi')
makedepends=('cmark-gfm' 'cython' 'python-build' 'python-installer' 'python-setuptools' 'python-cffi' 'python-wheel')
checkdepends=('python-pytest')
# Upstream did not tag the version on GitHub to match PyPi, see https://github.com/theacodes/cmarkgfm/issues/51
# source=("https://github.com/theacodes/cmarkgfm/archive/$pkgver/$pkgname-$pkgver.tar.gz")
source=("https://files.pythonhosted.org/packages/source/${_pyname::1}/$_pyname/$_pyname-$pkgver.tar.gz")
sha512sums=('87e97656861e63f5cc5e5b74b487cbf0daa7e4b3da00e260acdac21a70df6bf18227311b2bd99519003c6868e93bd4078ecfb18c10a5f6ba80f2a1fc72605473')

prepare() {
  cd cmarkgfm-$pkgver
  sed -e '/include_dirs=/a libraries=["cmark-gfm", "cmark-gfm-extensions"],' \
      -i src/cmarkgfm/build_cmark.py
}

build() {
  cd cmarkgfm-$pkgver
  python -m build -wn
}

check() {
  cd cmarkgfm-$pkgver
  local _pyver=cpython-$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
  PYTHONPATH="$PWD/build/lib.linux-$CARCH-${_pyver}" pytest
}

package() {
  cd cmarkgfm-$pkgver
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm0644 -t "$pkgdir"/usr/share/licenses/$pkgname/ LICENSE.txt
}
