# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-digitalocean
pkgver=1.17.0
pkgrel=9
pkgdesc='digitalocean.com API to manage Droplets and Images'
arch=('any')
license=('LGPL-3.0-only')
url='https://github.com/koalalorenzo/python-digitalocean'
depends=('python-requests' 'python-jsonpickle')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-pytest' 'python-responses')
source=("git+https://github.com/koalalorenzo/python-digitalocean.git#tag=v$pkgver")
sha512sums=('39be3bf6df79a1895d1a90cded4bf41ad723541696274ab57e55e71a8d5f52054ec6abfba564e581c42b28b1c5435879fa7d8bf401dbd3594ce62cb18d728b27')

build() {
  cd python-digitalocean
  python -m build --wheel --no-isolation
}

check() {
  cd python-digitalocean
  pytest --deselect digitalocean/tests/test_firewall.py
}

package() {
  cd python-digitalocean
  python -m installer --destdir="$pkgdir" dist/*.whl
}
