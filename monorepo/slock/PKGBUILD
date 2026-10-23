# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Maintainer: T.J. Townsend <blakkheim@archlinux.org>
# Contributor: Sebastian A. Liem <sebastian at liem dot se>

pkgname=slock
pkgver=1.7
pkgrel=1
pkgdesc="A simple screen locker for X"
arch=('x86_64')
url="https://tools.suckless.org/slock"
license=('MIT')
makedepends=('git')
depends=('libxext' 'libxrandr')
source=("git+https://git.suckless.org/slock#tag=${pkgver}")
sha256sums=('8f65a5d2ef546e2823349ae3ed6a632ad876b060c6ba2c27eb0a6ff0034429fd')

prepare() {
  cd $pkgname
  sed -i 's|static const char \*group = "nogroup";|static const char *group = "nobody";|' config.def.h
  sed -ri 's/((CPP|C|LD)FLAGS) =/\1 +=/g' config.mk
}

build() {
  cd $pkgname
  make X11INC=/usr/include/X11 X11LIB=/usr/lib/X11
}

package() {
  cd $pkgname
  make PREFIX=/usr DESTDIR="$pkgdir" install
  install -m644 -D LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
