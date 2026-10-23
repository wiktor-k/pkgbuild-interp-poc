# Maintainer: Torsten Keßler <tpkessler@archlinux.org>
# Contributor: Maikel Wever <maikelwever@gmail.com>
# Contributor: Matthew Gamble <git@matthewgamble.net>

pkgname=python-coloredlogs
pkgver=15.0.1
pkgrel=8
pkgdesc="Colored terminal output for Python's logging module"
arch=('any')
url="https://github.com/xolox/python-coloredlogs"
license=('MIT')
depends=('python' 'python-humanfriendly')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/${pkgver}.tar.gz"
        "pipes-removal-python3-13.patch")
sha256sums=('e8161222671bb129f7936cd220c275a3cbc0a6c22313bd4483114b9526e5695f'
            '13064e300c7c6c05f9e20b95e729650e48648fec7962ad516f20b39cfc0843a6')

prepare() {
  cd "python-coloredlogs-${pkgver}"
  patch -Np1 -i ../pipes-removal-python3-13.patch
}

build() {
  cd "python-coloredlogs-${pkgver}"

  python -m build --wheel --no-isolation
}

package() {
  cd "python-coloredlogs-${pkgver}"

  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE.txt"
}

# vim:set ts=2 sw=2 et:
