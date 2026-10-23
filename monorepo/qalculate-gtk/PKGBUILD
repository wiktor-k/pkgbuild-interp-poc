# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=qalculate-gtk
pkgver=5.13.0
pkgrel=1
pkgdesc='GTK frontend for libqalculate'
arch=(x86_64)
url='https://qalculate.github.io/'
license=(GPL-2.0-only)
depends=(cairo
         glib2
         glibc
         gtk3
         libgcc
         libqalculate
         libstdc++
         pango)
makedepends=(intltool
             git
             glib2-devel)
source=(git+https://github.com/Qalculate/qalculate-gtk#tag=v$pkgver)
sha256sums=('d6ded8437cb09cc88283e8d7f24c9b1cb7c5d995f717ed16515b2da06f15bba8')

prepare() {
  cd $pkgname
  autoreconf -vif -I /usr/share/gettext/m4
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --libexecdir=/usr/lib
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
