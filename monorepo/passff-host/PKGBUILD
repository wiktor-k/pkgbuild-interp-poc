# Maintainer: Daurnimator <daurnimator@archlinux.org>
# Contributor: Jaron Kent-Dobias <jaron@kent-dobias.com>
# Contributor: arcnmx <arcnmx@gmail.com>

pkgname=passff-host
pkgver=1.2.5
pkgrel=3
epoch=1
pkgdesc="PassFF native messaging host application for Firefox, Chromium, Chrome, Vivaldi"
url="https://codeberg.org/PassFF/passff-host"
license=('GPL-2.0-only')
depends=('pass' 'python')
optdepends=('pass-otp')
arch=('any')
source=("${pkgname}-${pkgver}.tar.gz::https://codeberg.org/PassFF/passff-host/archive/${pkgver}.tar.gz")
sha256sums=('34ac2c8365d514a10787a29ed880c460728e4f1c72f83f3371d293a993854fc6')

build() {
  cd "${pkgname}"

  sed -i -e "s/_VERSIONHOLDER_/${pkgver}/g" 'src/passff.py'
  sed -i -e 's/PLACEHOLDER/\/usr\/share\/passff\/passff.py/' 'src/passff.json'
}

package() {
  cd "${pkgname}"

  install -Dm 755 "src/passff.py" "${pkgdir}/usr/share/passff/passff.py"
  python -O -m compileall "${pkgdir}/usr/share/passff/passff.py"

  install -Dm 644 "src/passff.json" "${pkgdir}/usr/share/passff/passff.json"
  for d in \
    '/etc/chromium/native-messaging-hosts' \
    '/etc/opt/chrome/native-messaging-hosts' \
    '/usr/lib/mozilla/native-messaging-hosts' \
    '/etc/vivaldi/native-messaging-hosts' \
  ; do
    mkdir -p "${pkgdir}/${d}/"
    ln -s '/usr/share/passff/passff.json' "${pkgdir}/${d}/passff.json"
  done
}
