# Maintainer: Nebulosa <nebulosa2007-at-yandex-dot-ru>

pkgname=chawan
pkgver=0.5.0
pkgrel=1
pkgdesc='A text-mode web browser and pager for Unix-like systems'
arch=(x86_64)
url='https://chawan.net/'
license=(Unlicense)
depends=(
  glibc
  brotli
  libssh2
  openssl
  sh
)
makedepends=(git 'nim>=2.0.0')
source=("$pkgname::git+https://git.sr.ht/~bptato/chawan#tag=v$pkgver")
sha512sums=('4ba601e9bb5025a6f58ad21104b67e465af3f0cf7dbc7f92f354efe1606850d179cc09967009ff3efa21c3ca121f0e6b04af1cfcd09e91e0b73f6dd345a264ea')
b2sums=('5bb7d2130e40841d4854d82b805e4a861c04382e720f774d98d9f3ed6a76a6959c1ed6946030fc1c54edcb352a189e067a796c6e491846625825d73df1e2ea3e')

build() {
  cd "$pkgname"

  export CFLAGS+=" -ffile-prefix-map=$srcdir/="

  make LIBEXECDIR=/usr/lib/chawan
}

package() {
  cd "$pkgname"

  make DESTDIR="$pkgdir" PREFIX='/usr' LIBEXECDIR="$pkgdir/usr/lib/chawan" install
}
