# Maintainer: Frederik Schwan <freswa at archlinux dot org>

pkgname=python-bracex
pkgver=3.2
pkgrel=1
pkgdesc='Bash style brace expansion for Python'
arch=('any')
url='https://github.com/facelessuser/bracex'
license=('MIT')
makedepends=('python-build' 'python-hatchling' 'python-installer' 'python-wheel')
source=("https://github.com/facelessuser/bracex/archive/${pkgver}/bracex-${pkgver}.tar.gz")
b2sums=('65f54cdf23e7571784984fa57697d723e3f1d532e7ab077fcc0fe56808d264da668eab60f60cdeafda8cb396470326d4a9b8d9364dd76b482a037282d5f19358')

build() {
  cd bracex-${pkgver}
  python -m build --wheel --skip-dependency-check --no-isolation
}

package() {
  cd bracex-${pkgver}
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm644 LICENSE.md "${pkgdir}"/usr/share/licenses/${pkgname}/LICENSE
}
