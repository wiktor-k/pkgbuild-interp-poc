# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: twa022 <twa022 at gmail dot com>
# Contributor: Esteban V. Carnevale <alfplayer@mailoo.com>

pkgname=python-polib
_pkgname=polib
pkgver=1.2.0
pkgrel=4
pkgdesc='A library to manipulate gettext files'
url='https://pypi.python.org/pypi/polib'
arch=('any')
license=('MIT')
depends=('python')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel')
source=("python-polib-${pkgver}.tar.gz::https://pypi.io/packages/source/p/polib/polib-${pkgver}.tar.gz")
sha256sums=('f3ef94aefed6e183e342a8a269ae1fc4742ba193186ad76f175938621dbfc26b')

build() {
  cd "${srcdir}/polib-${pkgver}"
  python -m build --wheel --no-isolation
}

package() {
  cd "${srcdir}/polib-${pkgver}"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
