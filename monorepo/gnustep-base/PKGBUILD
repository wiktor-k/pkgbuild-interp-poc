# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Maintainer: Vesa Kaihlavirta <vegai@iki.fi>
# Contributor: Sebastian Sareyko <public@nooms.de>

pkgname=gnustep-base
pkgver=1.31.1
pkgrel=3
pkgdesc='The GNUstep base package'
arch=('x86_64')
url='http://www.gnustep.org/'
license=('GPL-2.0-only' 'GPL-3.0-only' 'LGPL-2.1-only')
depends=(avahi glibc gmp gnutls icu libffi libgcc libobjc libxml2 libxslt openssl zlib)
makedepends=(gcc-objc gnustep-make)
provides=('libgnustep-base.so')
groups=('gnustep-core')
options=('!emptydirs' '!makeflags')
source=(
  "https://github.com/gnustep/libs-base/releases/download/base-${pkgver//./_}/gnustep-base-${pkgver}.tar.gz"{,.sig}
  'ftbfs-gcc15.patch'
)
sha512sums=('6de3a935766d6b3075e60da71525511de940521dd50d604af1eb67681bd6aa4693b889e6655f39c96b53178a86623a4c4510be004487a28fddd175dec928c933'
            'SKIP'
            '29bff6ebeccbbad799f3b52e3b8a14a1878ac5cf70f3d8126ea2834790db02afce714ea327b3f73ba66e1ea1ab2542d087d860bfcae9949a12364f3091d1a398')
b2sums=('7bbab70fd5e78f75c0342d8d5a3e2175958a850b5ecc04fa6449e8a71aa998659755793d192c795865810d3b38ec10cfb284e7263a5ce301979930a25aa8a6ac'
        'SKIP'
        '280f796c91ce0d44f7a9bcff74a790d60e6800a08a5e410cea48f5f8d39ab3af7384f020cab986bf226d63561dd10c2e9ec6f5cf735cd484e775499a809f1ed1')
validpgpkeys=('83AAE47CE829A4146EF83420CA868D4C99149679') # GNUstep Maintainer <gnustep-maintainer@gnu.org>

prepare() {
  cd "$pkgname-$pkgver"

  # Fix FTBFS with gcc15 (-Wincompatible-pointer-types)
  patch -p1 -i "$srcdir/ftbfs-gcc15.patch"
}

build() {
  cd "$pkgname-$pkgver"

  # this uses malloc_usable_size, which is incompatible with fortification level 3
  export CFLAGS="${CFLAGS/_FORTIFY_SOURCE=3/_FORTIFY_SOURCE=2}"
  export CXXFLAGS="${CXXFLAGS/_FORTIFY_SOURCE=3/_FORTIFY_SOURCE=2}"

  source /usr/share/GNUstep/Makefiles/GNUstep.sh

  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc/GNUstep \
	  --with-ffi-include=/usr/lib/libffi-`pacman -Q libffi | cut -f2 -d\ |cut -f1 -d-`/include/

  # fix file ownership
  sed -i 's/tar -xf $(TIMEZONE_ARCHIVE);/tar -xf $(TIMEZONE_ARCHIVE);chown -R root:root * ;/' NSTimeZones/Makefile.postamble
  sed -i 's|.*gnutls_transport_set_lowat.*||' Source/GSSocketStream.m

  make
}

package() {
  cd "$pkgname-$pkgver"

  make DESTDIR="$pkgdir" install

  # fix file ownership
  chown -R root:root "$pkgdir"
}
