# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>
# Contributor: Kevin Azzam <aur at kevin.azz.am>
# Contributor: Ye Shu <shuye02 at outlook.com>

pkgname=onionshare
pkgver=2.6.5
pkgrel=1
pkgdesc='Share a file over Tor Hidden Services anonymously and securely'
url='https://github.com/onionshare/onionshare'
arch=('any')
license=('GPL-3.0-or-later')
depends=(
  'hicolor-icon-theme'
  'pyside6'
  'python'
  'python-cffi'
  'python-colorama'
  'python-eventlet'
  'python-flask'
  'python-flask-compress'
  'python-flask-socketio'
  'python-gevent'
  'python-gevent-websocket'
  'python-gnupg'
  'python-packaging'
  'python-psutil'
  'python-pynacl'
  'python-pysocks'
  'python-qrcode'
  'python-requests'
  'python-setuptools'
  'python-stem'
  'python-unidecode'
  'python-urllib3'
  'python-waitress'
  'python-werkzeug'
  'tor'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-poetry-core'
  'python-wheel'
)
optdepends=(
  'python-nautilus: enable Nautilus right-click extension'
)
checkdepends=(
  'python-pytest'
)
source=("${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('7dcee3152f5fb59a119bf6d15f26100008edcdfb5b8dc7c720f7e52c855e30addc3907683113145a73100e4d953af9f48eb83f74ce89ab2bcd5bb0526195c89c')
b2sums=('0aba5d59e4f25fa6b9cb33245bedce4e7126351a05d8a089b247d718f864f0bea47b68176c65824c4489f7b4bd03bd8a4042a37d75a8e68ca6bd71208c401da4')

build() {
  cd ${pkgname}-${pkgver}/cli
  python -m build --wheel --no-isolation
  cd ../desktop
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname}-${pkgver}
  pushd cli
  PYTHONPATH=".:${PYTHONPATH}" pytest tests/
  popd
  pushd desktop
  #PYTHONPATH="src:../cli:${PYTHONPATH}" pytest tests/
}

package() {
  cd ${pkgname}-${pkgver}
  python -m installer --destdir="${pkgdir}" cli/dist/*.whl
  python -m installer --destdir="${pkgdir}" desktop/dist/*.whl

  install -Dm 644 desktop/org.onionshare.OnionShare.desktop -t "${pkgdir}/usr/share/applications/"
  install -Dm 644 desktop/org.onionshare.OnionShare.svg -t "${pkgdir}/usr/share/icons/hicolor/scalable/apps/"
  install -Dm 644 desktop/org.onionshare.OnionShare.appdata.xml -t "${pkgdir}/usr/share/metainfo/"
}

# vim: ts=2 sw=2 et:
