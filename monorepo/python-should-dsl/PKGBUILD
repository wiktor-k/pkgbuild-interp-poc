# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-should-dsl
pkgver=2.1.2
pkgrel=8
pkgdesc="Should assertions in Python as clear and readable as possible"
url="https://github.com/nsi-iff/should-dsl"
license=('MIT')
arch=('any')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=("git+https://github.com/nsi-iff/should-dsl.git#commit=a8be403439d30fddb3907642b987403640869ae5")
sha512sums=('e18f5499f32593bc4c779ecc8bea96cf7172a5222bdf52e80507192f501b68c709de34cec7c3bb2cd2cc13b347834131de32a4d0199225ecde09ba96b3362de8')

build() {
  cd should-dsl
  python -m build --wheel --no-isolation
}

check() {
  cd should-dsl
  python run_all_examples.py
}

package() {
  cd should-dsl
  python -m installer --destdir="$pkgdir" dist/*.whl

  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
