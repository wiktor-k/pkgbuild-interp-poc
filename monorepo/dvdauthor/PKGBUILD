# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Giovanni Scafora <giovanni@archlinux.org>
# Contributor: Travis Willard <travisw@wmpub.ca>
# Contributor: Jaroslaw Swierczynski <swiergot@juvepoland.com>
# Contributor: Nicolai Lissner <nlissne@linux01.gwdg.de>

pkgname=dvdauthor
pkgver=0.7.2
pkgrel=16
pkgdesc='DVD authoring tools'
arch=(x86_64)
url='http://dvdauthor.sourceforge.net/'
license=(GPL-2.0-or-later)
depends=(
  bash
  fontconfig
  freetype2
  fribidi
  glibc
  imagemagick
  libdvdread
  libpng
  libxml2
)
makedepends=(
  docbook-sgml
  docbook-utils
  git
  perl-sgmls
)
source=(
  "git+https://github.com/ldo/dvdauthor.git#tag=$pkgver"
  dvdauthor-0.7.2-imagemagick7.patch
)
b2sums=(
  bf79688ded7fa17ae87d0e6dcd60d09e36df57fba2ac6203a5264b64b4bbcd1f81138b583ba52509372ac096efe4551e1e5d3c3d5ff5e6dcdac012e87adabf6e
  e7949e4d5d26f2b54157a13957199b509a4686c70c50be691890d4312b5b7aebf448729ee39328abd7599ed7c7a461942f22f0aa8ff0eb89734c583816203b20
)

prepare() {
  cd $pkgname

  # Don't search for obsolete freetype-config
  git cherry-pick -n d5bb0bdd542c33214855a7062fcc485f8977934e

  # Port to imagemagick 7
  # https://bugs.gentoo.org/610574#c2
  patch -Np1 -i ../dvdauthor-0.7.2-imagemagick7.patch

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
