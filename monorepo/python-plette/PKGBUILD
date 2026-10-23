# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-plette
pkgver=2.4.0
pkgrel=1
pkgdesc="Structured Pipfile and Pipfile.lock models"
url="https://github.com/sarugaku/plette"
license=('ISC')
arch=('any')
depends=('python-tomlkit')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-setuptools-scm' 'python-wheel')
checkdepends=('python-pytest')
source=("git+https://github.com/sarugaku/plette.git#tag=v$pkgver")
sha512sums=('adfe5e48eb9467a960fdadb5fa3bb4f320d4b97c97e83e48c7da230f541405a49ad380ad2b24ac1e265ab09d00a43028afcfcc643a0b45973e77fa1fbe84242d')

build() {
  cd plette
  python -m build -nw
}

check() {
  cd plette
  python -m venv .venv --system-site-packages
  .venv/bin/python -m installer dist/*.whl
  .venv/bin/python -m pytest
}

package() {
  cd plette
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
