#Maintainer: Jan de Groot <jgc@archlinux.org>

_pkgbasename=libxt
pkgname=lib32-$_pkgbasename
pkgver=1.3.1
pkgrel=1
pkgdesc="X11 toolkit intrinsics library (32-bit)"
arch=('x86_64')
url="https://xorg.freedesktop.org/"
license=('custom')
depends=('lib32-libsm' 'lib32-libx11' $_pkgbasename)
makedepends=('xorg-util-macros' 'gcc-multilib')
source=(${url}/releases/individual/lib/libXt-${pkgver}.tar.xz{,.sig})
sha512sums=('c220292f60b0f53134cf9364831a32bbaa9fa6bbb3a7143e917920957b7a48c616e946042747089f29ea9d8a18ecd64de620bcaf56d82462e7107de906f5db38'
            'SKIP')
validpgpkeys=('4A193C06D35E7C670FA4EF0BA2FB9E081F2D130E') #Alan Coopersmith <alan.coopersmith@oracle.com>
validpgpkeys+=('C52048C0C0748FEE227D47A2702353E0F7E48EDB') # "Thomas Dickey <dickey@invisible-island.net>"
validpgpkeys+=('3BB639E56F861FA2E86505690FDD682D974CA72A') #"Matt Turner <mattst88@gmail.com>"

build() {
  cd "${srcdir}/libXt-${pkgver}"

  export CC="gcc -m32"
  export PKG_CONFIG_PATH="/usr/lib32/pkgconfig"

  ./configure --prefix=/usr --sysconfdir=/etc --disable-static --libdir=/usr/lib32
  make
}

package() {
  cd "${srcdir}/libXt-${pkgver}"

  make DESTDIR="${pkgdir}" install

  rm -rf "${pkgdir}"/usr/{include,share}

  mkdir -p "$pkgdir/usr/share/licenses"
  ln -s $_pkgbasename "$pkgdir/usr/share/licenses/$pkgname" 
}
