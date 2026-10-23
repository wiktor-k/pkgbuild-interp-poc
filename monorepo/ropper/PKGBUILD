# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

pkgname=ropper
_pkgname=Ropper
pkgver=1.13.13
pkgrel=2
pkgdesc='Show information about binary files and find gadgets to build rop chains for different architectures'
url='https://github.com/sashs/Ropper'
arch=(any)
license=(BSD-3-Clause)
depends=(
  python
  python-capstone
  python-filebytes
  python-z3-solver
  python-pyvex
  python-archinfo
)
makedepends=(
  python-build
  python-installer
  python-keystone
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
)
optdepends=(
  'python-keystone: assemble command support'
)
conflicts=('python-ropper')
replaces=('python-ropper')
provides=('python-ropper')
source=(https://github.com/sashs/${_pkgname}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('f4a868c90bbae459b85e7d7846b39f0391de68c6fa36c8cf72572f1d420f127a9d052f1cf3560b4150d86064687433aa7a6af2c29bd85b89e5008fa3cd0556ad')
b2sums=('94e1c25346c219665318cd320cd98762450811cca146dcbc4740eed550f453e10c11d1b85356c8ee06df9a2b798ec93ca831bbe2170d6a73512ecc0cf53b2d2c')

build() {
  cd ${_pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_pkgname}-${pkgver}
  py.test
}

package() {
  cd ${_pkgname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  rm -rf "${pkgdir}"/usr/lib/python*/site-packages/testcases/
  install -Dm 644 README.md sample.py -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 COPYING -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
