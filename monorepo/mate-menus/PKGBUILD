# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-menus
pkgver=1.28.1
pkgrel=3
pkgdesc="MATE menu specifications"
url='https://github.com/mate-desktop/mate-menus'
arch=(x86_64)
license=('GPL-2.0-or-later AND LGPL-2.0-or-later')
depends=(
  glib2
  glibc
)
makedepends=(
  git
  gobject-introspection
  mate-common
)
groups=(mate)
source=("git+https://github.com/mate-desktop/mate-menus.git#tag=v$pkgver")
b2sums=(54e8e4e837c1b06a850b8300e85b69ff0bd4d497d8ce62d29860f7f2d585d1d4233125c0ba3f7f9010cd9441cd75f682462ff63737940566a3e8c1366e6312cd)

prepare() {
  cd $pkgname

  # Don't strip safety checks
  # https://github.com/mate-desktop/mate-menus/pull/129
  git cherry-pick -n abb1f769cb74772bce95681cb24202c390ae8755

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
