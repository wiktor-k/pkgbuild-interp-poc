# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-jaconv
pkgver=0.5.0
pkgrel=1
pkgdesc='Japanese character interconverter for Hiragana, Katakana, Hankaku, and Zenkaku'
url='https://ikegami-yukino.github.io/jaconv/jaconv.html'
license=(MIT)
arch=(any)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
)
source=("$pkgname::git+https://github.com/ikegami-yukino/jaconv#tag=v$pkgver")
sha512sums=('998292cb4e008e4c3ca986c5c5d03db4ba63267b1c0c3905ed1969e1223bfdb65f6c6eb04d90d542610e3668f40218ba8bea115fdb0f91f028cf7517dee21013')
b2sums=('ceb3025f9a5be78fb5314ad634c1184a354c9e6410579010fdaf74f23ac06f1090a54459a31399909f379856066cbab36b842be666567252d16516bb48ae6d7a')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  pytest -v
}
  
package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
