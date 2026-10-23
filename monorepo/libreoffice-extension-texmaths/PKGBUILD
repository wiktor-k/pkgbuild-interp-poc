# Maintainer: Antonio Rojas < arojas@archlinux.org >

pkgname=libreoffice-extension-texmaths
_extname=texmaths
pkgver=0.52.6
pkgrel=1
pkgdesc='A LaTeX equation editor for LibreOffice'
arch=(any)
url='http://roland65.free.fr/texmaths/'
license=(GPL-2.0-or-later)
depends=(libreoffice
         texlive-basic)
source=(https://downloads.sourceforge.net/project/$_extname/$pkgver/TexMaths-$pkgver.oxt)
sha256sums=('c4545539b17caa042110a356bf6f02631061a990e234b328670cb2ec8eecb9be')

package() {
  install -d -m755 "$pkgdir"/usr/lib/libreoffice/share/extensions/$_extname

  cp -r * "$pkgdir"/usr/lib/libreoffice/share/extensions/$_extname
  rm "$pkgdir"/usr/lib/libreoffice/share/extensions/$_extname/*.oxt
}
