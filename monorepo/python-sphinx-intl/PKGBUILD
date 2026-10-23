# Maintainer:

_pyname=sphinx-intl
pkgname=python-$_pyname
pkgver=2.4.0
pkgrel=1
pkgdesc='A Sphinx utility that make it easy to translate and to apply translations'
arch=(any)
url='https://github.com/sphinx-doc/sphinx-intl'
license=(BSD-2-Clause)
depends=(python
         python-babel
         python-click
         python-sphinx)
makedepends=(git
             python-build
             python-installer
             python-setuptools
             python-setuptools-scm)
source=(git+https://github.com/sphinx-doc/sphinx-intl#tag=$pkgver)
sha512sums=('e437c1b8fbc5f708a31c3a1b1d86546be401e833ae99e5c3ada6ee67c0d17b3f6b4882fcd4baa9625fa4ac734b3a57d75b608f6be98209f75939aad6833a1417')

build() {
  cd $_pyname

  python -m build --wheel --no-isolation
}

package() {
  cd $_pyname

  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
