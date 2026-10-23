# Maintainer: Christian Hesse <mail@eworm.de>
# Maintainer: Christian Heusel <gromit@archlinux.org>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Andrea Scarpino <andrea@archlinux.org>
# Contributor: dorphell <dorphell@archlinux.org>
# Contributor: Tom Newsom <Jeepster@gmx.co.uk>

pkgname=feh
pkgver=3.13.1
pkgrel=1
pkgdesc='Fast and light imlib2-based image viewer'
url='https://feh.finalrewind.org/'
license=('MIT')
arch=('x86_64')
depends=('curl' 'libcurl.so'
         'file' 'libmagic.so'
         'glibc'
         'hicolor-icon-theme'
         'imlib2' #'libImlib2.so'
         'libexif' 'libexif.so'
         'libpng' 'libpng16.so'
         'libx11' #'libX11.so'
         'libxinerama' #'libXinerama.so'
)
optdepends=('imagemagick: support more file formats')
makedepends=('git' 'libxt')
validpgpkeys=('429AF7B8E9EC9C0709D32F7F5333FB7712E24FE8'  # Birte Kristina Friesel <birte.friesel@uni-osnabrueck.de>
              '781BB7071C6BF648EAEB08A1100D5BFB5166E005'  # Daniel Friesel <derf@finalrewind.org> 
              '64FE6EC055560F9EF13A304419E6E524EBB177BA') # Derf Null <derf@ccc.de>
source=("git+https://git.finalrewind.org/feh.git?signed#tag=${pkgver}")
sha256sums=('0afdd8c517f67981444b4b4615b752524048ce6a5180d70175f5b7d7bc53b191')

build() {
  cd "${srcdir}/${pkgname}"
  make PREFIX=/usr \
    exif=1 \
    help=1 \
    inotify=1 \
    magic=1 \
    stat64=1
}

package() {
  cd "${srcdir}/${pkgname}"
  make PREFIX=/usr DESTDIR="${pkgdir}" install
  install -D -m0644 COPYING "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
