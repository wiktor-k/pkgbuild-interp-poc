# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-tabview
pkgver=1.5.0
pkgrel=1
pkgdesc="A curses command-line CSV and list (tabular data) viewer"
url="https://github.com/TabViewer/tabview"
license=('MIT')
arch=('any')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-hatchling')
checkdepends=('expect')
source=("git+https://github.com/TabViewer/tabview.git#tag=$pkgver")
sha512sums=('5cb6cad2e0aa9db514876cb77cd68e5cc8fc603e768f3ab473f93348bb154438ac7466dc88c613a4927689f00963a11ac9e4567a5ea3ac3b15943222cacd3acb')

build() {
  cd tabview
  python -m build --wheel --no-isolation
}

check() {
  cd tabview
  unbuffer python -m unittest discover
}

package() {
  cd tabview
  python -m installer --destdir="$pkgdir" dist/*.whl

  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
