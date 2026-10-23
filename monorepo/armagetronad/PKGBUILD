# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Kyle Keen <keenerd@gmail.com>
# Contributor: Mateusz Herych <heniekk@gmail.com>
# Contributor: Filippo 'JoeyrS' Civiletti <joeyrs@gmail.com>

pkgname=armagetronad
pkgver=0.2.9.3.1
pkgrel=1
pkgdesc='A Tron Clone in 3D'
arch=('x86_64')
url='http://armagetronad.net/'
license=('GPL-2.0-only')
# Disable build flags to fix FS#80060
options=('!buildflags')
depends=('sdl_image' 'libxml2' 'sdl_mixer' 'ftgl' 'hicolor-icon-theme')
source=("https://launchpad.net/armagetronad/0.2.9/${pkgver}/+download/armagetronad-${pkgver}.tbz")
sha256sums=('0ff6273cbe552dadbfb682c015c2fef0c999e0bb264e2b3d17a85bcddb6caf3d')

prepare() {
  cd "$srcdir/armagetronad-$pkgver"
  # Fix AppStream metadata
  # https://gitlab.com/armagetronad/armagetronad/-/merge_requests/197
  sed -i 's/x-flatpak-//g' desktop/armagetronad.appdata.xml.in
}

build() {
  cd "$srcdir/armagetronad-$pkgver"
  ./configure --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --disable-games \
    --enable-automakedefaults \
    --disable-uninstall
  make
}

package() {
  cd "$srcdir/armagetronad-$pkgver"
  DESTDIR="$pkgdir" make -j1 install
  mv "$pkgdir"/usr/share/{appdata,metainfo}
}
