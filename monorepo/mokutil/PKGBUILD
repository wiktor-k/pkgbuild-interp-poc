# Maintainer: Tobias Powalowski <tpowa@archlinux.org>
# Contributor: Mantas Mikulėnas <grawity@gmail.com>
pkgname=mokutil
pkgver=0.7.2
pkgrel=2
pkgdesc="Shim.efi Machine Owner Key management tool"
arch=(x86_64)
url="https://github.com/lcp/mokutil"
license=('GPL-3.0-or-later')
depends=(efivar openssl keyutils libxcrypt)
makedepends=(git)
source=("git+https://github.com/lcp/mokutil#tag=$pkgver")
b2sums=('8fc3481b293c5534a354322420e81e5b3f6dd26c3528d2a8792c6d66343c71e093e758f40603465dbf246fed0ba7b229ce7b0ba2cc28229162a00f2ce907e33e')

prepare() {
  cd "${pkgname}"
  autoreconf -f -i
}

build() {
  cd "${pkgname}"
  ./configure \
    --prefix=/usr \
    --sbindir=/usr/bin \
    --sysconfdir=/etc
  make
}

package() {
  cd "${pkgname}"
  make DESTDIR="$pkgdir" install
}

