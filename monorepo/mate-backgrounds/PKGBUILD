# Maintainer: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-backgrounds
pkgver=1.28.0
pkgrel=3
pkgdesc="Background images and data for MATE"
arch=(any)
url='https://github.com/mate-desktop/mate-backgrounds'
license=('GPL-2.0-or-later AND CC-BY-4.0')
groups=(mate)
options=(!emptydirs)
makedepends=(git)
source=("git+https://github.com/mate-desktop/mate-backgrounds.git#tag=v$pkgver")
b2sums=(f489c2ec57a3653bfaebb862650349b6c2d62fc32619798c4b866755e96e6e613b2981416a2a87a5210f9a2f2ed3bb46429b69f62d023d41ec26b220b1116884)

prepare() {
  cd $pkgname
  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
