# Maintainer: Frederik Schwan <freswa at archlinux dot org>
# Contributor:Francois Menning <f.menning@pm.me>
# Contributor: gilbus <aur(AT)tinkershell.eu>
# Contributor: Bruno Pagani <archange@archlinux.org>

pkgname=thermald
_pkgname=thermal_daemon
epoch=2
pkgver=2.5.13
pkgrel=1
pkgdesc='The Linux Thermal Daemon program from 01.org'
arch=('x86_64')
url='https://github.com/intel/thermal_daemon'
license=('GPL2')
depends=(
  dbus-glib
  libevdev
  libxml2
  upower
)
makedepends=(
  autoconf-archive
  git
  gtk-doc
  python
)
source=(
  "git+https://github.com/intel/thermal_daemon.git#tag=v${pkgver}"
)
b2sums=('86362962bf3487ad885ce97e477c014f751f2519df4c1a2488dca1678b7aef92a8e1b6537d70dee4d255268a14212f546a04ba64375a1c3b7437cd0b99682190')

build() {
  cd ${_pkgname}
  ./autogen.sh
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --with-dbus-sys-dir=/usr/share/dbus-1/system.d \
    --localstatedir=/var \
    --sbindir=/usr/bin
  make
}

package() {
  cd ${_pkgname}
  DESTDIR="${pkgdir}" make install
}
