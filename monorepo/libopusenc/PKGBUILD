# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Kyle Keen <keenerd@gmail.com>
# Contributor: Chocobo1 <chocobo1 AT archlinux DOT net>

pkgname=libopusenc
pkgver=0.3
pkgrel=1
pkgdesc='High-level API for encoding .opus files'
arch=(x86_64)
url='https://opus-codec.org/'
license=(BSD-3-Clause)
depends=(glibc opus)
makedepends=(git)
options=(staticlibs)
source=("https://archive.mozilla.org/pub/opus/libopusenc-$pkgver.tar.gz")
source=("$pkgname::git+https://gitlab.xiph.org/xiph/libopusenc#tag=v$pkgver")
sha512sums=('74d43e0dfee674a9a58aecd6b5036335ed400cec6190d4c5c35cc090133f286eecedcbd96618ef44c47f80e11a1aea304ab71b1f97d8b467b5d31b6300de284b')
b2sums=('e252c7b0cdd6168c87992468aa9c3d3271b6317cecc29472c6b69f3c3b83b313308cecdbb8a02dab1920de3b977c6d16ad94ad3159ca81ce15ce265e5650fcf7')

prepare() {
  cd "$pkgname"

  autoreconf -vfi
}

build() {
  cd "$pkgname"

  # LTO fat objects
  # https://archlinux.org/todo/lto-fat-objects/
  export CFLAGS+=" -ffat-lto-objects"

  ./configure --prefix="/usr"

  make
}

package() {
  cd "$pkgname"

  make DESTDIR="$pkgdir" install

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" COPYING

  rm -rf "$pkgdir/usr/share/doc"
}
