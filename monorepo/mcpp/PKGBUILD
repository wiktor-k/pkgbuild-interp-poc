# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Alexander Baldeck <alexander@archlinux.org>
# Contributor: Jan de Groot <jgc@archlinux.org>

pkgname=mcpp
pkgver=2.7.2.3
pkgrel=1
pkgdesc="Matsui's CPP implementation precisely conformed to standards"
arch=(x86_64)
url="http://mcpp.sourceforge.net"
license=(BSD-2-Clause)
depends=(glibc)
source=(
  "https://github.com/museoa/mcpp/archive/$pkgver/$pkgname-$pkgver.tar.gz"
  "$pkgname-double-namlen.patch"
)
b2sums=('2c43c4da044ebd9c0fb7b093dad8a6418e71bb3cd524db3535178cefaa2f0f20cf6dfa23308068d5b0647d82d757f3d7a491de5d44f848b2eb502f66f1ee353f'
        '8c9c0f2dfa7af0050492c821873ca0e5712fed1683447608af1cbae07b76f5b07d031b35076e03430de6565572e29bbf29182709513808cae1594387590a79fb')

prepare() {
  cd $pkgname-$pkgver
  patch -Np1 -i ../$pkgname-double-namlen.patch

  autoreconf -fiv
}

build() {
  cd $pkgname-$pkgver
  export CFLAGS=" \
    -Wno-implicit-function-declaration \
    -Wno-incompatible-pointer-types \
    -Wno-discarded-qualifiers \
  "
  ./configure --prefix=/usr --enable-mcpplib --disable-static
  make
}

package() {
  cd $pkgname-$pkgver
  make DESTDIR="$pkgdir" install
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
