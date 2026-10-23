# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=python-aioitertools
pkgver=0.13.0
pkgrel=2
pkgdesc='Implementation of itertools, builtins, and more for AsyncIO and mixed-type iterables'
arch=(any)
url='https://github.com/omnilib/aioitertools'
license=(MIT)
depends=(python)
makedepends=(
  git
  python-build
  python-flit-core
  python-installer
)
source=("git+$url#tag=v$pkgver")
sha256sums=('ea8628e25fa75b4d87046a2b7ca247d99ba5bf73fa5d721c07d7f5845cf0bf53')

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  python -m unittest -v aioitertools.tests
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
