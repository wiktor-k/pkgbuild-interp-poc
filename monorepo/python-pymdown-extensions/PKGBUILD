# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-pymdown-extensions
pkgver=12.1
pkgrel=1
pkgdesc="Extensions for Python Markdown"
arch=(any)
url="https://github.com/facelessuser/pymdown-extensions"
license=(MIT)
depends=(
  python
  python-markdown
  python-yaml
)
makedepends=(
  python-build
  python-hatchling
  python-installer
  python-wheel
)
checkdepends=(
  python-pytest
  python-pygments
)
optdepends=('python-pygments: for code highlighting')
provides=(pymdown-extensions)
replaces=(pymdown-extensions)
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('204bbf96f5b2671f4d6c4120e10061a203ec9dca8ef0baf2c4a85a81885371168494d0548a7d8cc535a96643d9e1ba76cad84ee01eab56bd15a3a783da51f3b3')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  pytest
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.md
}
