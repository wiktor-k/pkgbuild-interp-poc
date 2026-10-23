# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-filetype
pkgver=1.2.0
pkgrel=7
pkgdesc="Infer file type and MIME type of any file/buffer"
url="https://github.com/h2non/filetype.py"
license=('MIT')
arch=('any')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-setuptools')
checkdepends=('python-pytest' 'python-pytest-benchmark')
source=("https://github.com/h2non/filetype.py/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('e6116dd5505c01803c29016a8e85ce3f126c998f882eb1cfa8ff67990f560fee486d06b27dc64dbd74490f5d0f0327a9af0074e9cf1e16d3d55e71eeb2575c75')

build() {
  cd filetype.py-$pkgver

  python -m build --wheel --no-isolation
}

check() {
  cd filetype.py-$pkgver

  python -m pytest
}

package() {
  cd filetype.py-$pkgver

  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
