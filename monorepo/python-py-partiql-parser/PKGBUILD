# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=python-py-partiql-parser
pkgver=0.6.3
pkgrel=2
pkgdesc='Python Parser for PartiQL'
arch=(any)
url='https://github.com/getmoto/py-partiql-parser'
# Most files are licensed under MIT, while an Apache 2.0 file is vendored since 0.4.0.
# See: https://github.com/getmoto/py-partiql-parser/pull/12
license=('MIT AND Apache-2.0')
depends=(python)
makedepends=(
  git
  python-build
  python-hatchling
  python-installer
)
checkdepends=(python-pytest)
source=("git+$url#tag=$pkgver")
sha256sums=('724c16fd615917869332d2964bb0db238ffa4fdadb6485384fd0b79481a99781')

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  pytest
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
