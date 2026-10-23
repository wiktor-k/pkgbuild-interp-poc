# Maintainer: Levente Polyak <levente[at]leventepolyak[dot]net>
# Contributor: s1gma <s1gma@mindslicer.com>

pkgname=ropgadget
_pkgname=ROPgadget
pkgver=7.7
pkgrel=1
pkgdesc='Search gadgets in binaries to facilitate ROP exploitation for several file formats and architectures'
url='https://github.com/JonathanSalwan/ROPgadget'
arch=(any)
license=(BSD-3-Clause)
depends=(
  python
  python-capstone
)
makedepends=(
  python-setuptools
  python-build
  python-installer
  python-wheel
)
source=(https://github.com/JonathanSalwan/${_pkgname}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('658b25d8828b1162fcf830c82d5a44cf9968d14ba1b76f5da8eaf63aab1b50d0c452ed70b8a32ae5b8f35ff5887387775c0e5fdd215731c00c962dce6959389c')
b2sums=('83f4643d795770f4fdbe157e4f74f1fa0650bad52079c66e337995a4f270d37ce93258db273ff531477a6c97b94c780436879cc136ee7ee03a069cca739db598')

prepare() {
  cd ${_pkgname}-${pkgver}
  sed 's|python2|python|g' -i ropgadget/**/*.py ropgadget/*/*/*.py
}

build() {
  cd ${_pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pkgname}-${pkgver}
  ./ROPgadget.py --binary ./test-suite-binaries/elf-Linux-x86 --string "main"
}

package() {
  cd ${_pkgname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 LICENSE_BSD.txt -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim:set ts=2 sw=2 et:
