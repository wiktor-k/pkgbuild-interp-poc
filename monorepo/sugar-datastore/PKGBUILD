# Maintainer: Balló György <ballogyor+arch at gmail dot com>

pkgname=sugar-datastore
pkgver=0.121
pkgrel=5
pkgdesc='Sugar datastore service'
arch=(x86_64)
url='https://github.com/sugarlabs/sugar-datastore'
license=(GPL-2.0-or-later)
depends=(
  glib2
  glibc
  python
  python-dbus
  python-gobject
  python-xapian
)
makedepends=(git)
source=("git+https://github.com/sugarlabs/$pkgname.git#tag=v$pkgver")
b2sums=(d740a4e0c3a1b9e0149e64cbbe60c78e5dde4ac492bd717a16cad09cf8d387e739955799b68b44f56fd0624e2fe5af18901a802103fb0e87aec1da8fb9e2676f)

prepare() {
  cd $pkgname

  # Fix compatibility with Xapian 2.0
  # https://github.com/sugarlabs/sugar-datastore/pull/29
  sed -i 's/self._database.flush()/self._database.commit()/' src/carquinyol/indexstore.py

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
