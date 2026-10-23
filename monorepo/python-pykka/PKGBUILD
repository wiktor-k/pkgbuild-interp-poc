# Maintainer: Morten Linderud <foxboron@archlinux.org>
# Contributor: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Antoine Pierlot-Garcin <antoine at bokbox dot com>

pkgname=python-pykka
pkgver=4.4.2
pkgrel=1
pkgdesc="Easy to use concurrency abstractions for Python using the actor model"
arch=('any')
url="http://pykka.readthedocs.org/"
license=('Apache-2.0')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-hatchling')
optdepends=('python-gevent: Use gevent based actors from pykka.gevent')
source=("https://pypi.python.org/packages/source/p/pykka/pykka-$pkgver.tar.gz")
sha512sums=('daf0146ecb3de6719a61a9e4e75b79f2b7d4024b0e9e462feb774bfc1ee8258c78bb2e28190ecd65e37a28a6525dc5f5a3dcf95a5e95cb02950f8e265748dee2')

build() {
  cd "pykka-$pkgver"
  python -m build --wheel --no-isolation
}

package() {
  cd "pykka-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
