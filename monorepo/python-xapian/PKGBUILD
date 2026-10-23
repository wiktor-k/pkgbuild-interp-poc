# Maintainer: David Runge <dvzrv@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: DonVla <donvla/users.sourceforge/net>
# Contributor: Hervé Cauwelier <herve/oursours/net>

pkgname=python-xapian
_name=xapian-bindings
pkgver=2.1.0
pkgrel=1
epoch=1
pkgdesc="Python bindings for Xapian"
arch=(x86_64)
url="https://xapian.org/docs/bindings/python/"
license=(GPL-2.0-or-later)
depends=(
  libgcc
  libstdc++
  glibc
  python
)
makedepends=(
  python-sphinx
  "xapian-core=$epoch:$pkgver"
)
source=(https://oligarchy.co.uk/xapian/$pkgver/$_name-$pkgver.tar.xz{,.asc})
sha512sums=('05ae8a65dcfbf97b5e220e8b603f1771bd39a19db4d746a4e48670861898147036a7b17af6236c477d503737ac7f6f1e68c7f38e678e7888fa931e559c7de1ef'
            'SKIP')
b2sums=('62f2fd1d78422c00fb134f42dd6e694c2d6b4861504e2a710e0624d1bdd6e15f74b1660fbb503fa118e0d0596eb62a0cbf9313437da6742f4899668f500a1b2a'
        'SKIP')
validpgpkeys=('08E2400FF7FE8FEDE3ACB52818147B073BAD2B07') # Olly Betts <olly@debian.org>

prepare() {
  cd $_name-$pkgver
  autoreconf -fiv
}

build() {
  cd $_name-$pkgver
  ./configure --prefix=/usr --with-python3
  # prevent excessive overlinking due to libtool
  sed -i -e 's/ -shared / -Wl,-O1,--as-needed\0/g' libtool
  make
}

check() {
  cd $_name-$pkgver
  make check
}

package() {
  depends+=(
    xapian-core libxapian.so
  )

  cd $_name-$pkgver
  make DESTDIR="$pkgdir" install
  install -vDm 644 {AUTHORS,ChangeLog,HACKING,NEWS,README,TODO} -t "$pkgdir/usr/share/doc/$_name"
}
# vim:set ts=2 sw=2 et:
