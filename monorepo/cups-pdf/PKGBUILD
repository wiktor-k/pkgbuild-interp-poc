# Maintainer: Andreas Radke <andyrtr at archlinux.org>
# Contributor: Thomas Baechler <thomas.baechler@rwth-aachen.de>

pkgname=cups-pdf
pkgver=3.0.3
pkgrel=1
pkgdesc="PDF printer for cups"
arch=(x86_64)
depends=('cups' 'libcups' 'ghostscript' 'glibc')
install=cups-pdf.install
url="https://www.cups-pdf.de/welcome.shtml"
license=('GPL-2.0-or-later')
source=(https://www.cups-pdf.de/src/cups-pdf_$pkgver.tar.gz)
backup=(etc/cups/cups-pdf.conf)
# http://www.cups-pdf.de/src/cups-pdf_3.0.3.tar.gz.md5 - no other checksums
md5sums=('3a9fdedbf5805e6e49cc4bb0f14620bf')
sha256sums=('4a66e16976b30f3c6df5ed113ed9d6e13f7f34617eab349f70b0382fce54e08b')

build() {
  cd "$pkgname-$pkgver"/src
  [ -z "$CC" ] && CC=gcc
  $CC $CFLAGS $LDFLAGS -s cups-pdf.c -o cups-pdf -lcups  
}

package() {
  cd "$pkgname-$pkgver"/src
  install -D -m700 cups-pdf "$pkgdir"/usr/lib/cups/backend/cups-pdf

  # Install Postscript Color printer
  cd ../extra
  install -D -m644 CUPS-PDF_noopt.ppd "$pkgdir"/usr/share/cups/model/CUPS-PDF_noopt.ppd
  install -D -m644 CUPS-PDF_opt.ppd "$pkgdir"/usr/share/cups/model/CUPS-PDF_opt.ppd

  # Install config file
  install -D -m644 cups-pdf.conf "$pkgdir"/etc/cups/cups-pdf.conf

  # Install docs
  install -D -m 644 ../README "$pkgdir"/usr/share/doc/$pkgname/README
  
  # use cups group from cups pkg FS#57311/FS#56818/FS#57313
  chgrp -R cups "${pkgdir}"/etc/cups
  sed -i "s:### Default\: lp:### Default\: cups:" "${pkgdir}"/etc/cups/cups-pdf.conf
  sed -i "s:#Grp lp:Grp cups:" "${pkgdir}"/etc/cups/cups-pdf.conf
}
