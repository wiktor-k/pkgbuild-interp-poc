# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-praw
pkgver=8.0.3
pkgrel=1
pkgdesc="Python Reddit API Wrapper that allows for simple access to Reddit's API"
arch=('any')
license=('BSD-2-Clause')
url="https://praw.readthedocs.org"
depends=('python-defusedxml' 'python-prawcore' 'python-update-checker' 'python-websocket-client')
makedepends=('python-build' 'python-installer' 'python-hatchling')
checkdepends=('python-pytest' 'python-vcrpy')
source=("https://github.com/praw-dev/praw/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('6a0f0c78e7f74e6256609cc8c53ae8aaef0d0c0e75d44ea20170171679a1f496b84533c3472bfa64ec869c331ffefb61426e6c3c73539070bb2187c303bdbba8')

build() {
  cd praw-$pkgver
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd praw-$pkgver
  pytest
}

package() {
  cd praw-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
