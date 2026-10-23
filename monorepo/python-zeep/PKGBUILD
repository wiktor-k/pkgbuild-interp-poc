# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-zeep
pkgver=4.3.3
pkgrel=1
pkgdesc="A fast and modern Python SOAP client"
arch=('any')
license=('MIT')
url="https://docs.python-zeep.org"
depends=('python-attrs' 'python-isodate' 'python-lxml' 'python-platformdirs' 'python-requests'
         'python-requests-toolbelt' 'python-requests-file')
makedepends=('python-setuptools' 'python-wheel' 'python-build' 'python-installer')
checkdepends=('python-pytest' 'python-freezegun' 'python-pretend'
              'python-requests-mock' 'python-pytest-httpx' 'python-pytest-asyncio'
              'python-httpx' 'python-packaging' 'python-xmlsec')
optdepends=('python-httpx: async transport support'
            'python-packaging: async transport support'
            'python-xmlsec: WS-Security signature support')
source=("https://github.com/mvantellingen/python-zeep/archive/$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('eeafd71ab39c1075b678db1d99818bf65750b7fb1b1bd9b8b323eed1e6a30852a0f87c6daaa1f0a06c6c0c8ff7a08e7730ff9a038636adf85ba58431077bc06a')

build() {
  cd python-zeep-$pkgver
  python -m build --wheel --no-isolation
} 

check() {
  cd python-zeep-$pkgver
  PYTHONPATH="$PWD"/build/lib pytest
}

package() {
  cd python-zeep-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
