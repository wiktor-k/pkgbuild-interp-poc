# Maintainer: Christian Hesse <mail@eworm.de>
# Contributor: Aaron Fischer <mail@aaron-fischer.net>
# Contributor: Steven Allen <steven@stebalien.com>
# Contributor: trile7 at gmail dot com
# Contributor: Ernia <monghitri@aruba.it>

pkgname=yad
pkgver=15.0
pkgrel=2
pkgdesc='A fork of zenity - display graphical dialogs from shell scripts or command line'
url='https://github.com/v1cont/yad'
arch=('x86_64')
license=('GPL-3.0-or-later')
depends=('gtk3' 'webkit2gtk-4.1' 'gtksourceview4' 'gspell')
source=("https://github.com/v1cont/yad/releases/download/v${pkgver}/${pkgname}-${pkgver}.tar.xz")
sha256sums=('7cf4dc136a643bf484ceab27d9b28a513624dffe38a08731132c0e4e24073772')

prepare() {
  cd "${srcdir}/${pkgname}-${pkgver}"

  sed -i 's/gtksourceview-3.0/gtksourceview-4/' configure.ac

  autoreconf -ivf
}

build() {
  cd "${srcdir}/${pkgname}-${pkgver}"

  ./configure \
    --prefix=/usr \
    --enable-icon-browser \
    --enable-html \
    --enable-spell \
    --enable-sourceview

  make
}

package() {
  cd "${srcdir}/${pkgname}-${pkgver}"

  make DESTDIR="${pkgdir}" install
}
