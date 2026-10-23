# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=jimtcl
pkgver=0.83
pkgrel=3
pkgdesc='A small footprint implementation of Tcl'
arch=(x86_64)
url='https://jim.tcl-lang.org/home/doc/www/www/index.html'
license=(BSD-2-Clause)
depends=(
  glibc
  openssl
  zlib
)
makedepends=(git)
source=("$pkgname::git+https://github.com/msteveb/jimtcl#tag=$pkgver")
sha512sums=('68a1728d8cec2fc36e8ea78245e516e081c10c139d0f344974b4c8583bd07cce2f9a9d02ee195a9b9e3757687ccfbc9ecc0338898bb8b7aa7b9b94c4170cf89c')
b2sums=('4cf96d0864bb28a485ac8c6cdecd35bf960e9277bf527a86592e006300311ef4cde82a0c4d8c480c2f58e224aa92b896b4cc15d2284b210b3c580881be78e4bf')

prepare() {
  cd "$pkgname"

  # workaround for buggy readline.h
  # https://github.com/msteveb/jimtcl/issues/308
  git cherry-pick --no-commit 35e0e1f9b1f018666e5170a35366c5fc3b97309c

  ./configure \
    --shared \
    --prefix=/usr \
    --docdir=/usr/share/doc/jimtcl
}

build() {
  cd "$pkgname"

  make
}

package() {
  cd "$pkgname"

  make DESTDIR="$pkgdir" install

  # documentation
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" README

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
# vim:set ts=2 sw=2 et:
