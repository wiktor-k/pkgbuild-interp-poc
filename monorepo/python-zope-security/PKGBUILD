# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-zope-security
_pkgname=zope.security
pkgver=8.4
pkgrel=1
pkgdesc="Zope Security Framework"
arch=('x86_64')
url="https://github.com/zopefoundation/zope.security"
license=('ZPL-2.1')
depends=(
  'glibc'
  'python'
  'python-zope-component'
  'python-zope-i18nmessageid'
  'python-zope-interface'
  'python-zope-location'
  'python-zope-proxy'
  'python-zope-schema'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-btrees'
  'python-zope-configuration'
  'python-zope-testing'
  'python-zope-testrunner'
)
optdepends=('python-zope-configuration: for ZCML support')
source=("git+$url.git#tag=$pkgver")
sha512sums=('60356662ee188b515b9c605a2295e924c95d26f26532596a7aeec0df27eee8a9164bcb974813488b3c72a5c73564dee239e80f033bf9e32d1eccad8f43277dd1')

build() {
  cd $_pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pkgname
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m zope.testrunner -vc --test-path src/
}

package() {
  cd $_pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
}
