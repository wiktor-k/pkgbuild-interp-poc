# Maintainer:  Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=mate-themes
pkgver=3.22.26
pkgrel=4
pkgdesc="Official themes for the MATE desktop"
arch=(any)
url='https://github.com/mate-desktop/mate-themes'
license=('LGPL-2.1-only AND GPL-3.0-or-later')
makedepends=(git)
optdepends=('mate-icon-theme: default icon theme')
options=(!emptydirs)
groups=(mate)
source=("git+https://github.com/mate-desktop/mate-themes.git#tag=v$pkgver")
b2sums=(c1214b7aec17a4f34e21d01454102ea5bf12d37c8428d15d435b4c172c9c7efd844e930156ebcfaa450a5c0d175b713017de459a37b68d0979b7371ccbf74f28)

prepare() {
  cd $pkgname

  # Remove GTK2 theme engine check
  git cherry-pick -n b347b074ca0c48f2d33b0b9acffd96c5e368f5a5

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
