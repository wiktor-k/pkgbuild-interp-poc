# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-icon-theme-faenza
pkgver=1.20.0
pkgrel=5
pkgdesc="Faenza icon theme for MATE"
arch=(any)
url='https://github.com/mate-desktop-legacy-archive/mate-icon-theme-faenza'
license=(GPL-3.0-only)
makedepends=(git)
source=("git+https://github.com/mate-desktop-legacy-archive/mate-icon-theme-faenza.git#tag=v$pkgver")
b2sums=(c533667e0e4c4998784fd425e7f2adda5bf686be8c2bde1cc12b9454c4708923a292fa864b995d1b3efc328aec5f68bb3918dcdeb7fdd5f259a4b13f326f0fb8)

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
