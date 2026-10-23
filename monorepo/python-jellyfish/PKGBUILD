# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Jelle van der Waa <jelle@archlinux.org>

pkgname=python-jellyfish
pkgver=1.3.0
pkgrel=1
pkgdesc='A python library for doing approximate and phonetic matching of strings'
arch=(x86_64)
url='https://github.com/jamesturk/jellyfish'
license=(BSD)
depends=(python)
makedepends=(
  git
  python-maturin
  python-build
  python-installer
)
source=("git+https://github.com/jamesturk/jellyfish.git#tag=v${pkgver}")
b2sums=('8cd7cde2e7152fce52a22b5048d8736c923a383a650fe3fc13cb65248b7c339a00ea1f08a69871b91213b06ed767b96b7b9f74d719e9b0264d46b1ce83c87312')

pkgver() {
  cd jellyfish
  git describe --tags | sed 's/^v//'
}

build() {
  cd jellyfish
  python -m build --wheel --no-isolation
}

package() {
  python -m installer --destdir="${pkgdir}" jellyfish/dist/*.whl
  install -Dm 644 jellyfish/LICENSE -t "${pkgdir}"/usr/share/licenses/python-jellyfish/
}

# vim: ts=2 sw=2 et:
