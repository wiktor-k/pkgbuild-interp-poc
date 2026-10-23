# Maintainer: Maxime Gauduin <alucryd@archlinux.org>

pkgname=python-jinxed
pkgver=2.1.0
pkgrel=1
pkgdesc='A curses replacement for Windows and Unix-like systems'
arch=(any)
url=https://github.com/Rockhopper-Technologies/jinxed
license=(MIT)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=(git+https://github.com/Rockhopper-Technologies/jinxed.git#tag=${pkgver})
b2sums=('ceda8826c6bf917be33672bccf05b899f2964636cdc4f78d91b6f8064795ee1abe015e04d5cd2c7e2a35869f7181503c3695013fa9f76f6ce0d507d5f4c79eab')

build() {
  cd jinxed
  python -m build --wheel --no-isolation
}

package() {
  cd jinxed
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/python-jinxed/"
}
