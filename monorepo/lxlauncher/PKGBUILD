# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor:  Bartłomiej Piotrowski <nospam@bpiotrowski.pl>
# Contributor: Angel 'angvp' Velasquez <angvp[at]archlinux.com.ve> 
# Contributor: Geoffroy Carrier <geoffroy.carrier@koon.fr>

pkgname=lxlauncher
pkgver=0.2.8
pkgrel=3
pkgdesc='Clone of the Asus launcher for EeePC'
arch=(x86_64)
url='https://github.com/lxde/lxlauncher'
license=(GPL-2.0-or-later)
groups=(lxde)
depends=(
  cairo
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  libx11
  lxmenu-data
  menu-cache
  startup-notification
)
makedepends=(
  git
  intltool
)
backup=(
  etc/xdg/lxlauncher/gtk.css
  etc/xdg/lxlauncher/gtkrc
  etc/xdg/lxlauncher/settings.conf
)
conflicts=(lxlauncher-gtk3)
replaces=(lxlauncher-gtk3)
source=(
  "git+https://github.com/lxde/lxlauncher.git#tag=$pkgver"
  lxlauncher-no-decoration.patch
)
b2sums=(
  260cb4ad514381bc9cb4230f98227d9f43dd4237a031604fdcf93de47d680bc2b0127833de6008b113839277fbd1a39db7239d5b42a5855bbbb8b9df6dca7847
  6bf2f542bf4ad0fdba152638bf2e474daf2ec1c98029c178799d6c45193733d600e6f0b9318391f3cf51b4a55b5672af0d93c47f00f30c4098344c6cd6c2196d
)

prepare() {
  cd $pkgname

  # Disable window decoration
  # https://github.com/lxde/lxlauncher/pull/4
  git apply -3 ../lxlauncher-no-decoration.patch

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-gtk3
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
