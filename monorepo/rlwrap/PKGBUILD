# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Kyle Keen <keenerd@gmail.com>
# Contributor: wizzomafizzo <wizzomafizzo@gmail.com>
# Contributor: Mateusz Herych <heniekk@gmail.com>
# Contributor: Benjamin Andresen <benny AT klapmuetz DOT org>
# Contributor: Douglas Thrift <douglas@douglasthrift.net>

pkgname=rlwrap
pkgver=0.48
pkgrel=1
pkgdesc='Adds readline-style editing and history to programs'
arch=(x86_64)
url='https://github.com/hanslub42/rlwrap'
license=(GPL-2.0-only)
depends=(
  glibc
  readline
  ncurses
  perl
  python
  python-pexpect
  libptytty
)
makedepends=(git)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('b5761ad4bff124ece68c41560c1cabb9e2c66588b2037f7d204614bd9bef34a3317f6c0a992e4b3093539cacc4d5091b1dbdec9ed63a99788c1225a95870d35e')
b2sums=('8f27bb182a9f17ab64bbdb0e36d265618916f90073843c007098b133276936dff6442839087078dd244274fd19ce0e7c18ad8d15bb0bc5e710998795db5e2ec4')

prepare() {
  cd "$pkgname"

  autoreconf -vi
}

build() {
  cd "$pkgname"

  ./configure --prefix=/usr

  make
}

check() {
  cd "$pkgname"

  make check
}

package() {
  cd "$pkgname"

  make DESTDIR="$pkgdir" install
}
