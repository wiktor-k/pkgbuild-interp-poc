# Maintainer: Andreas Radke <andyrtr@archlinux.org>

pkgbase=gutenprint
pkgname=('gutenprint' 'foomatic-db-gutenprint-ppds')
pkgver=5.3.5
pkgrel=1
pkgdesc="Top quality printer drivers for POSIX systems"
arch=('x86_64')
license=('GPL-2.0-or-later'
         'LGPL-2.0-or-later'
         'MIT AND GPL-3.0-or-later'
         'Bison-exception-2.2')
url="https://gimp-print.sourceforge.io/"
makedepends=('cups' 'ghostscript'
            # for the docs
            'dialog' 'doxygen' 'docbook-utils' 'texi2html' 'texlive-bin')
options=('!emptydirs')
source=(https://downloads.sourceforge.net/gimp-print/$pkgname-$pkgver.tar.xz)
sha1sums=('65e1ad99e9a3df17d505e719a9123eaa3b50fd67')
sha256sums=('f5a9f47de28530b1ae2069cfbc647a9a641baeeabe809bb0ef2b3ec5b9668d70')

prepare(){
  cd "${pkgbase}"-${pkgver}
  # sbindir is not properly taken over there
  sed -i "s:cups_sbindir=\"\${cups_exec_prefix}\/sbin\":cups_sbindir=\"\${cups_exec_prefix}\/bin\":" m4local/stp_cups.m4

  # make build reproducible taken from Debian rules
  # https://sources.debian.org/src/gutenprint/5.3.4.20220624T01008808d602-1/debian/rules/
  # Remove build directory
  find -name Makefile.in -exec sed -i 's/$(CURDIR)/BUILDPATH/g' {} +

  autoreconf -vfi
}

build() {
  cd "${pkgbase}"-${pkgver}
  ./configure --prefix=/usr \
    --sbindir=/usr/bin \
    --disable-rpath \
    --enable-samples \
    --disable-static \
    --disable-static-genppd \
    --enable-cups-ppds \
    --enable-simplified-cups-ppds=only \
    --enable-translated-cups-ppds \
    --enable-globalized-cups-ppds \
    --disable-libgutenprintui2 #--help

  # make build reproducible taken from Debian rules
  # Remove timestamp and uname output
  sed -i -e 's,Generated at.*,Generated at REDACTED,g' \
	 -e 's,uname -a output:.*,uname -a output: REDACTED,g' config.summary

  # globalized ppds -> put all translations into one ppd file
  # enable translated cups ppds grows size 9,5MB -> 156MB
  # simplified cups ppds (yes, no, only) - only offer basic paper sizes, not all options 
  # 4,3MB (only), 17MB (both) -> translated 41MB simpl., 165MB, 231MB both

  sed -i -e 's/ -shared / -Wl,-O1,--as-needed\0/g' libtool
  make
}

package_gutenprint() {
  pkgdesc="Top quality printer drivers for POSIX systems"
  install=gutenprint.install
  depends=('cups' 'cups-filters'
           'glib2' 'zlib' 'glibc' 'readline')
  optdepends=('libusb:	required for drivers that depend on gutenprint52usb backend'
            'foomatic-db-gutenprint-ppds: prebuilt ppd files'
            'perl:	to run cups-genppdupdate')
  replaces=('gimp-print')
  
  cd "${pkgname}"-${pkgver}
  make DESTDIR="${pkgdir}" install

  # split out files
  mkdir "$srcdir"/tmp_foomatic-db-gutenprint-ppds
  
  mv "${pkgdir}"/usr/share/cups/model "$srcdir"/tmp_foomatic-db-gutenprint-ppds/

  # cleanup
  rm -rf "${pkgdir}"/usr/share/foomatic 
  rm -rf "${pkgdir}"/etc

  # license
  mkdir -p "${pkgdir}"/usr/share/licenses/${pkgname}
  install -m644 COPYING "${pkgdir}"/usr/share/licenses/${pkgname}/
}

package_foomatic-db-gutenprint-ppds() {
  pkgdesc="simplified prebuilt ppd files"
  optdepends=('cups:	to use cups printer spooler(recommended)')
  
  cd "${pkgbase}"-${pkgver}
  mkdir -p "${pkgdir}"/usr/share/cups
  mv "$srcdir"/tmp_foomatic-db-gutenprint-ppds/* "${pkgdir}"/usr/share/cups/

  # license
  mkdir -p "${pkgdir}"/usr/share/licenses/${pkgname}
  install -m644 COPYING "${pkgdir}"/usr/share/licenses/${pkgname}/
}
