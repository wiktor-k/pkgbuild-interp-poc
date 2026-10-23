# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>

pkgname=python-rpyc
_name=rpyc
pkgver=6.0.1
pkgrel=3
pkgdesc='Remote Python Call - A transparent and symmetric RPC library for python'
url='https://github.com/tomerfiliba-org/rpyc'
arch=('any')
license=('MIT')
depends=(
  'python'
  'python-plumbum'
)
makedepends=(
  'python-build'
  'python-hatchling'
  'python-installer'
  'python-wheel'
)
source=(https://github.com/tomerfiliba-org/${_name}/archive/${pkgver}/${_name}-${pkgver}.tar.gz)
sha256sums=('8b73d5ae5bd055dc89635182f9625d4399e7f2dacebdbfce8bf2ad5ee590a00c')
b2sums=('b90b8e3f16fc6c5ac481437f7b729f6bedf3f6504e3ace724b5fddccb5b42dc3354532925f3d4eaa4e1fd4fa632e649f10e562828f88585ac445a97747e53e49')

build() {
  cd "${_name}-${pkgver}"
  python -m build --wheel --no-isolation
}

package() {
  cd "${_name}-${pkgver}"
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -D LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim:set ts=2 sw=2 et:
