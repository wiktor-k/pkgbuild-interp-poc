# Maintainer: Giancarlo Razzolini <grazzolini@archlinux.org>
# Contributor: Martin Rys <rys.rs/contact>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: Geoffroy Carrier <geoffroy.carrier@aur.archlinux.org>
# Contributor: neodreams <yanbrodeur@videotron.ca>
# Contributor: Daniel J Griffiths <ghost1227@archlinux.us>

pkgname=icoutils
pkgver=0.32.3
pkgrel=4
pkgdesc='Extracts and converts images in MS Windows(R) icon and cursor files.'
arch=('x86_64')
license=('GPL-3.0-or-later')
url='https://www.nongnu.org/icoutils/'
depends=('libpng>=1.0.0' 'perl-libwww>=5.64')
validpgpkeys=('3776FFA2692D84ECA895F98E3424F3704FF4EDEC')
source=(
  "https://download.savannah.gnu.org/releases/${pkgname}/${pkgname}-${pkgver}.tar.bz2"{,.sig}
  "iterators.patch"
)
sha256sums=('17abe02d043a253b68b47e3af69c9fc755b895db68fdc8811786125df564c6e0'
            'SKIP'
            '94537e3d65db0d86c7c66a96457657a6984a73dd329d80bffb1fcba7a78030a3')

prepare() {
  cd ${pkgname}-${pkgver}
  patch -Np1 -i "${srcdir}/iterators.patch"
}

build() {
  cd ${pkgname}-${pkgver}
  ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man
  make
}

package() {
  cd ${pkgname}-${pkgver}
  make DESTDIR="$pkgdir" install
}
