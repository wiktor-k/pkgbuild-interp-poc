# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-betamax-serializers
pkgver=0.2.1
pkgrel=12
pkgdesc="A set of third-party serializers for Betamax"
arch=('any')
license=('Apache-2.0')
url='https://gitlab.com/betamax/serializers'
depends=('python-betamax')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=("git+https://gitlab.com/betamax/serializers.git#tag=$pkgver")
sha256sums=('c33ef6f2a7753d8d5746c3d4607562f3da678408b9e671f9743cc8925926ae37')

build() {
  cd serializers
  python -m build --wheel --no-isolation
}

package() {
  cd serializers
  python -m installer --destdir="$pkgdir" dist/*.whl
}
