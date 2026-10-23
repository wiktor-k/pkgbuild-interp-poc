# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-user-guide
pkgver=1.28.0
pkgrel=3
pkgdesc="MATE User Guide"
arch=(any)
url='https://github.com/mate-desktop/mate-user-guide'
license=(GFDL-1.1-or-later)
groups=(mate)
depends=(yelp)
makedepends=(
  git
  yelp-tools
)
source=("git+https://github.com/mate-desktop/mate-user-guide.git#tag=v$pkgver")
b2sums=(f933b1cedcee42f6d09e91325d6f3b17c6f63d0b410e0824b8b70ca3f1bd2a0bad516e0f1b9e9e142e99c3f6d7b4675dc8621508f22d90d0bd8c7b94c48ca29b)

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
