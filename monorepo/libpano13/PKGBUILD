# Maintainer: Bruno Pagani <archange@archlinux.org>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Tobias Kieslich <tobias@justdreams.de>

pkgname=libpano13
pkgver=2.9.23
pkgrel=1
pkgdesc="Basic library to calculate panoramical pictures - core functions of the panotools"
arch=(x86_64)
url="http://panotools.sourceforge.net/"
license=(GPL-2.0-or-later)
depends=(libjpeg-turbo libpng libtiff suitesparse)
makedepends=(cmake)
source=(https://downloads.sourceforge.net/sourceforge/panotools/${pkgname}-${pkgver}.tar.gz)
sha256sums=('e7c076d37a14c39434962115e47ddbe18452ca3de5ce40e2aaefa7cf5815ea28')

build() {
  cmake -B build -S ${pkgname}-${pkgver} \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DUSE_SPARSE_LEVMAR=ON
  make -C build
}

package() {
  make -C build DESTDIR="${pkgdir}" install
  mkdir "${pkgdir}"/usr/share/doc
  mv "${pkgdir}"/usr/share/{pano13/doc,doc/pano13}
  rmdir "${pkgdir}"/usr/share/pano13
}
