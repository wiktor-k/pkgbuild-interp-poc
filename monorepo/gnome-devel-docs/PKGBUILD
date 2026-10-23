# Maintainer: Jan Alexander Steffens (heftig) <heftig@archlinux.org>
# Contributor: Jan de Groot <jgc@archlinux.org>

pkgname=gnome-devel-docs
pkgver=40.3
pkgrel=3
pkgdesc="Developer documentation for GNOME"
url="https://gitlab.gnome.org/GNOME/gnome-devel-docs"
arch=(any)
license=(FDL)
depends=(yelp)
makedepends=(yelp-tools git)
source=("git+https://gitlab.gnome.org/GNOME/gnome-devel-docs.git#tag=$pkgver")
b2sums=('4a81ecba337050dc8f8baa29f0ef8f4bc626d09231b50587bd914cf31b12ca782317d3d0fa647bf3e7ac18765f1fe2a46d7c1bfd29b1ffc834a4500a1ebbe658')

prepare() {
  cd $pkgname
  autoreconf -fvi
}

build() {
  cd $pkgname
  ./configure --prefix=/usr --sysconfdir=/etc --localstatedir=/var
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
