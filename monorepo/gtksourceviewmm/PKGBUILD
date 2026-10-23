# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Ionut Biru <ibiru@archlinux.org>

pkgbase=gtksourceviewmm
pkgname=(
  gtksourceviewmm
  gtksourceviewmm-docs
)
pkgver=3.91.1
pkgrel=2
epoch=1
pkgdesc='C++ bindings for the gtksourceview library'
arch=(x86_64)
url='https://gitlab.gnome.org/GNOME/gtksourceviewmm'
license=(LGPL-2.1-or-later)
depends=(
  atkmm
  cairomm
  glib2
  glibc
  glibmm
  gtk3
  gtkmm3
  gtksourceview4
  libgcc
  libsigc++
  libstdc++
  pangomm
)
makedepends=(
  git
  gtkmm3-docs
  mm-common
)
source=("git+https://gitlab.gnome.org/GNOME/$pkgname.git#tag=$pkgver")
b2sums=(25f6fc2f05b947153fd69f6d20c45261c15c453d74b4609d04d97dad6731fd6d754f4244879137b6399206c79e7db33a328764faa669413b974ba92703206e76)

prepare() {
  cd $pkgbase
  mm-common-prepare -cf
  autoreconf -fi
}

build() {
  cd $pkgbase
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-maintainer-mode
  sed -i -e 's/ -shared / -Wl,-O1,--as-needed\0/g' libtool
  make M4PATH='/usr/lib/gtkmm-3.0/proc/m4/:/usr/lib/pangomm-1.4/proc/m4/:/usr/lib/atkmm-1.6/proc/m4/'
}

package_gtksourceviewmm() {
  cd $pkgbase
  make DESTDIR="$pkgdir" install doc_subdirs=
}

package_gtksourceviewmm-docs() {
  pkgdesc='Developer documentation for gtksourceviewmm'
  depends=(gtkmm3-docs)

  cd $pkgbase
  make -C docs DESTDIR="$pkgdir" install
}
