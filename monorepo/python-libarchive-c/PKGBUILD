# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

pkgname=python-libarchive-c
pkgver=5.3
pkgrel=1
pkgdesc='Python interface to libarchive'
url='https://github.com/Changaco/python-libarchive-c'
arch=(any)
license=(CC0-1.0)
depends=(
  python
  libarchive
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
)
source=(https://github.com/Changaco/python-libarchive-c/archive/${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('0cd1f95bf0045ce882012e4b217d3deacb921fb34fd3395afddacba6b1a291cc5a5383cadd1c2c0edae86b52beff535b801cb4ade714f510762835fc314c3d72')
b2sums=('771a9f4373c1ac94a283e20f50ee0437934ab3e6adf961802e20841176cc36e8a4151fcc6951559db78ad90ea8f0a005b03596b8fc55df4db522b05bdc1dfe14')

prepare() {
  sed '/os.umask(2)/d' -i ${pkgbase}-${pkgver}/setup.py
  cd ${pkgbase}-${pkgver}
}

check() {
  cd ${pkgbase}-${pkgver}
  LC_CTYPE=en_US.UTF-8 py.test
}

build() {
  cd ${pkgbase}-${pkgver}
  python -m build --wheel --no-isolation
}

package() {
  cd ${pkgbase}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 README.rst -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 LICENSE.md -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
