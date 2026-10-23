# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=python-sarif-om
_pkgname=sarif-python-om
pkgver=1.0.4
pkgrel=9
pkgdesc='Python classes for the SARIF object model'
arch=(any)
url='https://github.com/microsoft/sarif-python-om'
license=(MIT)
depends=(python-attrs)
makedepends=(
  python-build
  python-installer
  python-pbr
  python-wheel
)
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('dfc652990998ffe5c9935bb2b30e981a2ff0a7763fce70b2f728b8ab3e25426b8beb4255deacd20e9406c4becca9975bf6923ba13060f6b9e94778151300aa06')

build() {
  cd $_pkgname-$pkgver
  export PBR_VERSION=$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd $_pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
