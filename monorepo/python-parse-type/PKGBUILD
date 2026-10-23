# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-parse-type
pkgver=0.6.6
pkgrel=2
pkgdesc="Simplifies to build parse types based on the parse module"
url="https://github.com/jenisys/parse_type"
license=('MIT')
arch=('any')
depends=('python-parse' 'python-six')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-setuptools-scm' 'python-wheel')
checkdepends=('python-pytest')
source=("https://github.com/jenisys/parse_type/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('abef14f38045913287fd41a19c6fab90d3771cea223bf26f543dfcf9f67b404a659a0f24d5ac71c3bc7c1490b1773ed63a85363734594e7cce4d4b4dc646ee75')

build() {
  cd parse_type-$pkgver
  python -m build -nw
}

check() {
  cd parse_type-$pkgver
  pytest -o addopts=''
}

package() {
  cd parse_type-$pkgver
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
