# Maintainer: Laurent Carlier <lordheavym@gmail.com>
# Contributor: Lone_Wolf <lonewolf at xs4all dot nl>
# Contributor: Moritz Lipp <mlq@pwmt.org>

pkgname=lib32-libunwind
pkgver=1.8.3
pkgrel=1
pkgdesc="Determine and manipulate the call-chain of a program (32-bit)"
url="https://www.nongnu.org/libunwind/"
arch=(x86_64)
license=(MIT)
depends=(
  lib32-glibc
  lib32-xz
  lib32-zlib
  libunwind
)
makedepends=()
provides=(
  libunwind-{coredump,ptrace,setjmp,x86}.so
  libunwind.so
)
source=(
  https://github.com/libunwind/libunwind/releases/download/v$pkgver/libunwind-$pkgver.tar.gz{,.asc}
)
b2sums=('13ac888e72c080b9e2588024578740971f721497864c5f2eb930227cbc5bb040551db6280048b0b6f2dde193d5d1719f546088618ebd2e19f96524c2b767fda1'
        'SKIP')
validpgpkeys=(
  42FA3D4C00D0AA116C3F45DAA4CCF616E0FF69D2 # Stephen M. Webb <stephenw@xandros.com>
)

prepare() {
  cd libunwind-$pkgver
}

build() {
  local configure_options=(
    --build=$CHOST
    --host=i686-pc-linux-gnu
    --prefix=/usr
    --libdir=/usr/lib32
    --sysconfdir=/etc
    --localstatedir=/var
    --disable-documentation
    --disable-tests
  )

  export CC="gcc -m32"
  export CXX="g++ -m32"
  export PKG_CONFIG="i686-pc-linux-gnu-pkg-config"

  cd libunwind-$pkgver
  ./configure "${configure_options[@]}"
  sed -i -e 's/ -shared / -Wl,-O1,--as-needed\0/g' libtool
  make
}

check() {
  cd libunwind-$pkgver
  # Tests fail on i686: https://github.com/libunwind/libunwind/issues/393
  make check || :
}

package() {
  cd libunwind-$pkgver
  make DESTDIR="$pkgdir" install

  # we need one specific include file for the i686 platform
  find "$pkgdir/usr/include" -type f -not -name "*x86*" -exec rm {} \;
}

# vim:set sw=2 sts=-1 et:
