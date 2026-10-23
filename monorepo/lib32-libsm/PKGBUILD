# Maintainer: Jan de Groot <jgc@archlinux.org>

_pkgbasename=libsm
pkgname=lib32-$_pkgbasename
pkgver=1.2.6
pkgrel=1
pkgdesc="X11 Session Management library (32-bit)"
arch=('x86_64')
license=('custom')
url="https://xorg.freedesktop.org/"
depends=('lib32-libice' 'lib32-util-linux' $_pkgbasename)
makedepends=('xorg-util-macros' 'xtrans' 'gcc-multilib' 'xorgproto')
source=(${url}/releases/individual/lib/libSM-${pkgver}.tar.xz)
sha512sums=('e544a1dc49a03390f76af35837bfd01f749b806d88d3ee806f20311c47ff53d0aeea4744feb875958031b17d50b57a89dcc41d81241c09333c88b268c44653a7')

build() {
  cd libSM-${pkgver}

  export CC="gcc -m32"
  export PKG_CONFIG_PATH="/usr/lib32/pkgconfig"

  ./configure --prefix=/usr --sysconfdir=/etc --libdir=/usr/lib32
  make
}

package() {
  cd libSM-${pkgver}
  make DESTDIR="${pkgdir}" install

  rm -rf "${pkgdir}"/usr/{include,share}

  mkdir -p "$pkgdir/usr/share/licenses"
  ln -s $_pkgbasename "$pkgdir/usr/share/licenses/$pkgname" 
}
