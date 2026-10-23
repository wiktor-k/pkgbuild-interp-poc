# Maintainer: Maxime Gauduin <alucryd@archlinux.org>

pkgname=python-django-picklefield
pkgver=3.4.0
pkgrel=3
pkgdesc='A pickled object field for Django'
arch=(any)
url=https://github.com/gintas/django-picklefield
license=(MIT)
depends=(
  python
  python-django
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=(git+https://github.com/gintas/django-picklefield.git#tag=v${pkgver})
b2sums=('bc0bf6be5cfc8730f0d689d023276f74b9fe923597cf1ed95e19aa8ba8eb58d73475a9a1fbeb0fc2651d4ef86e7d8f940abe3002e852e89d0de9080aa942800d')

build() {
  cd django-picklefield
  python -m build --wheel --no-isolation
}

package() {
  cd django-picklefield
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}"/usr/share/licenses/python-django-picklefield/
}
