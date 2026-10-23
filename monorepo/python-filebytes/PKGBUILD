# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

pkgname=python-filebytes
_pyname=${pkgname/python-/}
pkgver=0.10.2
pkgrel=10
pkgdesc='Library to read and edit executable binary files'
url='https://github.com/sashs/filebytes'
arch=(any)
license=(BSD-3-Clause)
depends=(
  python
)
makedepends=(
  python-build
  python-installer
  python-wheel
  python-setuptools
)
source=(https://github.com/sashs/filebytes/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz fix-ast-str-access.patch)
sha512sums=('12317d6b5063e710b4410e4dcc61b45c0ac1cbc80655212dacb0929a19e100849df39ba84a05cb96699abe7ec16758e0b9656b3d3ce044f3d122b8636106e564'
            '72c4b588bad569c685ed3c33aa322b32e4d5ce0a5f3a4df90da83d709b40ea08e3643dab101391dbec243f8ee695236443ed151a4f2de8e62e7aa12130bfe6cd')
b2sums=('6b7fb865d01cc9ea0f14b32f6e0326da21fa2fa2a36549b99ca88e9bf015162d3b0f82e350be1cca805294fd56ef47e2db4394db5ffb3061773f146eda1ba666'
        '743e2b83e7e7cd2ff341f694eb95f3f703eb08a9ba3dfb7f2efc026af3843e97429518065687dba69aabf32faea3ce29e02aa36f2a1003e637d48d382fa1830f')

prepare() {
  cd ${_pyname}-${pkgver}
  patch -Np1 -i ${srcdir}/fix-ast-str-access.patch
}

build() {
  cd ${_pyname}-${pkgver}
  python -m build --wheel --no-isolation
}

package() {
  cd ${_pyname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 COPYING -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 samples/* -t "${pkgdir}/usr/share/doc/${pkgname}/samples"
}

# vim: ts=2 sw=2 et:
