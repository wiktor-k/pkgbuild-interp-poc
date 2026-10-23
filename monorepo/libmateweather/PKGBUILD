# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=libmateweather
pkgver=1.28.2
pkgrel=3
pkgdesc='Provides access to weather information from the Internet'
arch=(x86_64)
url='https://github.com/mate-desktop/libmateweather'
license=(GPL-2.0-or-later)
depends=(
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libsoup3
  libxml2
)
makedepends=(
  git
  glib2-devel
  gtk-doc
  mate-common
  tar
)
source=(
  "git+https://github.com/mate-desktop/libmateweather#tag=v$pkgver"
  libmateweather-libsoup3.patch
)
b2sums=(
  af967ea1debb57f97746811c2edf74d8a8f72b8879f603439a7b188634296c868f608269b66e2dff6d06d203679a8bdcf6221161b5e43ee162daef888fb70239
  59552942ed4f0887e76b271cfe67a519f6bdc48754311b498e16174a8126109097beae02728ea492ecf092ed9286a9178b2ecb6d50841a76ee73c87b3fbc4954
)

prepare() {
  cd $pkgname

  # Port to libsoup-3.0
  git apply -3 ../libmateweather-libsoup3.patch

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-locations-compression \
    --enable-gtk-doc
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
