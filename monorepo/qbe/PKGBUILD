# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=qbe
pkgver=1.3
pkgrel=1
pkgdesc='Small embeddable C compiler backend'
arch=(x86_64)
url='https://c9x.me/compile/'
license=(MIT)
depends=(glibc)
makedepends=(git)
source=("$pkgname::git://c9x.me/qbe.git#tag=v$pkgver")
sha512sums=('ebd326173756019ae23a9826735c000c80a70d780d00b82aa4ce1be0cc8970d779e006bb2f564ef5969da92866e3c499be32553823ccbe769400fa3578a913d1')
b2sums=('a924d576880eb455edc3a1cb4d0300765eff7aa4456d05753741dc2c25dcbaf25ef002b88536390f04f8b8b3361a664f1e0561c88e5ff0b7bef9aed2d8e2da32')

build() {
  cd "$pkgname"

  make
}

check() {
  cd "$pkgname"

  make -k check
}

package() {
  cd "$pkgname"

  make DESTDIR="$pkgdir" PREFIX=/usr install

  # documentation
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" doc/*

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
