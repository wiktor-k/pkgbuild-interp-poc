# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Bartłomiej Piotrowski <bpiotrowski@archlinux.org>
# Contributor: Angel Velasquez <angvp@archlinux.org>
# Contributor: Juergen Hoetzel <juergen@archlinux.org>

pkgname=lxsession
pkgver=0.5.7
pkgrel=1
epoch=1
pkgdesc='Lightweight X11 session manager'
arch=(x86_64)
url='https://github.com/lxde/lxsession'
license=(GPL-2.0-or-later)
depends=(
  bash
  cairo
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  libx11
  polkit
)
makedepends=(
  git
  intltool
  vala
)
conflicts=(
  lxpolkit
  lxsession-gtk3
)
replaces=(
  lxpolkit
  lxsession-gtk3
)
groups=(lxde)
source=("git+https://github.com/lxde/lxsession.git#tag=$pkgver")
b2sums=(049a4bd86469f95e6f260fac0ab51458bf23a19a202665be3cb6c7b5489ed096c602dcb807885d6d1ba9531600a5a29571d8d1e9aadc4767e0c69f94e2286604)

prepare() {
  cd $pkgname

  # Replace "which" command
  # https://github.com/lxde/lxsession/pull/43
  sed -i 's/which /command -v /' lxlock/lxlock

  mkdir m4
  autoreconf -fiv
}

build() {
  cd $pkgname
   ./configure \
     --prefix=/usr \
     --sysconfdir=/etc \
     --localstatedir=/var \
     --libexecdir=/usr/lib \
     --enable-gtk3
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
