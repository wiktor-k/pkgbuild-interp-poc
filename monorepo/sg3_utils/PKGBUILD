# Maintainer: Hyacinthe Cartiaux <hyacinthe@archlinux.org>
# Maintainer: Claudia Pellegrino <auerhuhn@archlinux.org>
# Contributor: Daniel Isenmann <daniel@archlinux.org>

pkgname=sg3_utils
pkgver=1.49
pkgrel=1
pkgdesc="Generic SCSI utilities"
arch=(x86_64)
url="http://sg.danny.cz/sg/sg3_utils.html"
license=('GPL-2.0-or-later AND BSD-2-Clause')
depends=(
  'bash'
  'glibc'
  'libgcc'
)
source=("https://sg.danny.cz/sg/p/${pkgname}-${pkgver}.tar.xz")
sha512sums=('00c76d367e076feff460b619d05a043ab6404642043c000f7a5f93ba503e0bbcc53ebe408d405d252b155197bf456fcc3a8eb8519153825a9c616b2ed8ba9b10')

prepare() {
  cd "${pkgname}-${pkgver}"
  autoreconf -fiv
}

build() {
  cd "${pkgname}-${pkgver}"
  ./configure --prefix=/usr
  make
}

package() {
  cd "${pkgname}-${pkgver}"
  make DESTDIR="${pkgdir}" install
  install -m755 -d "${pkgdir}/usr/share/licenses/${pkgname}"
  install -m644 COPYING "${pkgdir}/usr/share/licenses/${pkgname}/"
}
