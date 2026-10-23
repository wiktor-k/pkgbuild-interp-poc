# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=pysolfc-cardsets
pkgver=3.1
pkgrel=1
pkgdesc="A collection of cardsets adapted for use with PySolFC"
arch=('any')
url="http://pysolfc.sourceforge.net/"
license=('GPL-3.0-or-later')
source=("https://downloads.sourceforge.net/pysolfc/PySolFC-Cardsets-$pkgver.tar.bz2")
sha256sums=('ed6dc447477a65fb60f4abd9ceca47fe142e801133ad90ee02e7d711a3637b35')

package() {
  cd PySolFC-Cardsets-$pkgver
  install -vdm755 "$pkgdir/usr/share/PySolFC"
  cp -va -t "$pkgdir/usr/share/PySolFC/" ./*
}
