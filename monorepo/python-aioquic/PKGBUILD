# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>

pkgname=python-aioquic
_name=aioquic
pkgver=1.6.1
pkgrel=1
pkgdesc='QUIC and HTTP/3 implementation in Python'
url='https://github.com/aiortc/aioquic'
arch=('x86_64')
license=('BSD-3-Clause')
depends=(
  'glibc'
  'openssl'
  'python'
  'python-certifi'
  'python-cryptography'
  'python-pylsqpack'
  'python-pyopenssl'
  'python-service-identity'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
source=(https://files.pythonhosted.org/packages/source/${_name::1}/$_name/$_name-$pkgver.tar.gz)
sha256sums=('9a10d20e69b5914d6fa034becb30ab75ef6a777a3f1b16623ca599e9291c0a5c')
b2sums=('ea5b965bf35e27a0bed0b71d48c0c175b67487d47026ecba2d499db37b4126fb017be1780ac796eb9cea8a4c64ccd97c063dbb812dab126c02262b3d1887aa4e')

build() {
  cd "${_name}-${pkgver}"
  python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd ${_name}-${pkgver}
  python -m installer --destdir=test_dir dist/*.whl

  PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH" python -c '
import aioquic
import aioquic.h3.connection
import aioquic.quic.packet
'
}

package() {
  cd "${_name}-${pkgver}"
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -D LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim:set ts=2 sw=2 et:
