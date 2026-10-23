# Maintainer: David Runge <dvzrv@archlinux.org>

_name=calmjs.parse
pkgname=python-calmjs.parse
pkgver=1.3.4
pkgrel=2
pkgdesc="Various parsers for ECMA standards"
arch=(any)
url="https://github.com/calmjs/calmjs.parse"
license=(MIT)
depends=(
  python
  python-ply
  python-setuptools
)
makedepends=(
  python-build
  python-installer
  python-wheel
)
checkdepends=(python-pytest)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('de97ca594f5ecb71ec981251d8944f02dd08e3e0a0b01a7c75f14569415f01ecbfac885c7fa83930477cf2cac9445eebb088aa45ccac875da1dacd357275169e')
b2sums=('56a8e15ae398475e86e73b0b2adf57470a925028c7ad528d2c20bdd778734c5a18656ef16b977f331e2ecc7765d9e422ec3322bd9d1652b56caec25bcc66e72d')

prepare() {
  cd $_name-$pkgver
  # generate symbol tables required for lex and yacc
  # NOTE: this may require a specific python-ply version!
  PYTHONPATH="src:$PYTHONPATH" python -m calmjs.parse.parsers.optimize
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  export PYTHONPATH="$PWD/src:$PYTHONPATH"
  PYTHONDONTWRITEBYTECODE=1 pytest -vv
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGES,README}.rst -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
