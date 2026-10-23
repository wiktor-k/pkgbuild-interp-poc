# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Kyle Keen <keenerd@gmail.com>

pkgname=jshon
pkgver=20131105
pkgrel=6
pkgdesc='A json parser for the shell'
arch=('x86_64')
url='https://github.com/keenerd/jshon'
license=('MIT')
depends=('jansson')
makedepends=('git')
source=("$pkgname::git+$url#tag=$pkgver")
sha512sums=('6fad927d6e319e613507db9fa0a0b098612ec500e74b6c91a29b7e2752b6488a1194a5b13a8b33e9ce3a3350148dcd9e8d06b4089ee254c3f79346072985b6aa')
b2sums=('2646235147588def8a6c19f395df67f67bcfc28d6072c8cd51394ac11553d22b5db4232fdbf11295919a57f67cffc1e795f692a537904599f43bd0c7e47c953b')

build() {
  cd "$pkgname"

  make
}

package() {
  cd "$pkgname"

  # binary
  install -vDm755 -t "$pkgdir/usr/bin" jshon

  # man page
  install -vDm644 -t "$pkgdir/usr/share/man/man1" jshon.1

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
