# Maintainer: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Camille Moncelier <pix@devlife.org>

pkgname=m17n-lib
pkgver=1.8.7
pkgrel=1
pkgdesc='Multilingual text processing library (runtimes)'
url='https://www.nongnu.org/m17n/'
arch=('x86_64')
license=('LGPL-2.1-or-later')
depends=('libxft' 'm17n-db' 'fribidi' 'libthai' 'libxml2' 'gd' 'libotf')
makedepends=('git' 'anthy' 'ispell')
optdepends=('anthy: to use the anthy input method'
            'ispell: to use the ispell input method')
options=('!makeflags' '!emptydirs')
validpgpkeys=('D68584F9B40A6502FD7C53015F592A4029890630')
source=("git+https://git.savannah.nongnu.org/git/m17n/m17n-lib.git#tag=REL-${pkgver//./-}")
sha256sums=('f982ffffd252b08b9be567c9a4431461c674ede76be72a1cafe9d398b6233f42')

prepare() {
  cd ${pkgname}
  ./bootstrap.sh
}

build() {
  cd ${pkgname}
  ./configure --prefix=/usr
  make
}

package() {
  cd ${pkgname}
  make DESTDIR="${pkgdir}" install
}
