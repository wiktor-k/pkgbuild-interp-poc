# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-zope-deprecation
pkgver=6.0
pkgrel=3
pkgdesc="Zope Deprecation Infrastructure"
arch=('any')
url="https://github.com/zopefoundation/zope.deprecation"
license=('ZPL-2.1')
depends=('python')
makedepends=(
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=('python-zope-testrunner')
source=("$pkgname-$pkgver.tar.gz::$url/archive/$pkgver.tar.gz")
sha512sums=('8f332c4ecb42d78a0965d7fc9984d03efffc7add1309aec9c4db59dd39da84b255f35aa39aebe159b9ae612e54e1cbe8c5c515099b59d38c980940b4971627c8')

build() {
  cd zope.deprecation-$pkgver
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd zope.deprecation-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m zope.testrunner -vc --test-path src/
}

package() {
  cd zope.deprecation-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
