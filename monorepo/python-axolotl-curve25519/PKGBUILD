# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Alexandre Bouvier <contact AT amb DOT tf>
# Contributor: Ivan Shapovalov <intelfx@intelfx.name>
# Contributor: Tommaso Sardelli <lacapannadelloziotom AT gmail DOT com>
# Contributor: Philipp Joram <phijor AT t-online DOT de>

pkgname=python-axolotl-curve25519
_pkgver=0.4.1-2
pkgver=${_pkgver//-/.}
pkgrel=14
pkgdesc='Python wrapper for curve25519 library with ed25519 signatures'
url='https://github.com/tgalal/python-axolotl-curve25519'
arch=(x86_64)
license=(GPL-3.0-or-later)
depends=(
  glibc
  python
  zlib
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=(${pkgname}-${pkgver}.tar.gz::https://github.com/tgalal/python-axolotl-curve25519/archive/${_pkgver}.tar.gz
        pymoduledef-gcc-compile-error.patch)
sha256sums=('cc8de0c140831519e45598b082fdef0e939fd30930be442e38bdec13dc16aab5'
            'b47d0eb47d4dd0534dc4fdfd6254f4e083c4cd6e94b393e078a7955a176c7ee0')
sha512sums=('cbc7c6caa47a9a811640c247a1be727d7b1b68bcdb4c5336e02b4d1eaf9fd2c57b7438b0da466a379a1c0f3f146756b9b7eea3c9b7945ce88478d4bf0b8a1e0d'
            '3d6e540451ee9bfee53d9781e661669661c0d5263a884bb2e30ce1c1e127d3ba6c5a3229795be5676be5250965dc32d88e1c7673a7ae291c964f10bc7d2c0a8d')

prepare() {
  cd ${pkgname}-${_pkgver}
  patch -Np1 -i ${srcdir}/pymoduledef-gcc-compile-error.patch
}

build() {
  echo "Building python..."
  cd ${pkgname}-${_pkgver}
  python -m build --wheel --no-isolation
}

package() {
  cd ${pkgname}-${_pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
