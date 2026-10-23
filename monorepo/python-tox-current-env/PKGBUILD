# Maintainer: Filipe Laíns (FFY00) <lains@archlinux.org>

_pkgname=tox-current-env
pkgname=python-$_pkgname
pkgver=0.0.16
pkgrel=2
pkgdesc='tox plugin to run tests in current Python environment'
arch=('any')
url='https://github.com/fedora-python/tox-current-env'
license=('MIT')
depends=('python' 'python-tox')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")
sha512sums=('272d6866fbca16efdbb3b4ae034e7dc1bb68c467590998bdf2de152a091e029821b4002e6db001a92aafba063582fe0f4d1f40b43e06a7ecc62aa75dc16e6c45')

build() {
  cd $_pkgname-$pkgver

  python -m build -nw
}

package() {
  cd $_pkgname-$pkgver

  python -m installer -d "$pkgdir" dist/*.whl

  install -Dm 644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}

# vim:set ts=2 sw=2 et:
