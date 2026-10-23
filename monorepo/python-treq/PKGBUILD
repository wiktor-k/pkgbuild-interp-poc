# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=python-treq
pkgver=26.7.0
pkgrel=1
pkgdesc="A requests-like API built on top of twisted.web's Agent"
arch=('any')
license=('MIT')
url='https://github.com/twisted/treq'
depends=('python' 'python-incremental' 'python-hyperlink'
         'python-pyopenssl' 'python-service-identity' 'python-idna' 'python-twisted'
         'python-attrs' 'python-typing_extensions' 'python-multipart'
         'python-zope-interface')
makedepends=('git' 'python-build' 'python-installer' 'python-hatchling' 'python-wheel' 'httpbin')
source=("git+https://github.com/twisted/treq.git#tag=treq-$pkgver")
sha512sums=('bfe3bf0aef5077ae338fd86586530a409538f3d1613a27a13e9b4339e301d12ca4f47aca320a3c6d914cc4d8da797748afaed66655be2dc3e44c84e21128b6af')

build() {
  cd treq
  python -m build --wheel --no-isolation
}

check() {
  cd treq
  PYTHONPATH="$PWD/src" trial treq
}

package() {
  cd treq
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
