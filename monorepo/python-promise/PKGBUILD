# Maintainer: Maxime Gauduin <alucryd@archlinux.org>

pkgname=python-promise
pkgver=2.3.0
pkgrel=10
pkgdesc='Ultra-performant Promise implementation in Python'
url=https://github.com/syrusakbary/promise
arch=(any)
license=(MIT)
depends=(
  python
  python-gevent
  python-six
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=(git+https://github.com/syrusakbary/promise.git#tag=v${pkgver})
b2sums=('9497a3787d2c12bb6882378d9f7f5b72926b365ae20a2d96b728973728e47fdc48ede437f5cfeea23faa335f74dc4045aa07794a81c2ec442a73e205233c34a8')

build() {
  cd promise
  python -m build --wheel --no-isolation
}

package() {
  cd promise
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}"/usr/share/licenses/python-promise/
}
