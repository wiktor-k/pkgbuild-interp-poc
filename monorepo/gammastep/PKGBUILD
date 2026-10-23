# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Robin Candau <antiz@archlinux.org>

pkgname=gammastep
pkgver=2.0.11
pkgrel=2
pkgdesc="Adjust the color temperature of your screen according to your surroundings."
arch=(x86_64)
url="https://gitlab.com/chinstrap/gammastep"
license=(GPL-3.0-or-later)
depends=(
  glibc
  hicolor-icon-theme
  libdrm
  libx11
  libxcb
  libxxf86vm
  wayland
)
makedepends=(
  glib2
  intltool
  python
  systemd
)
optdepends=(
  'geoclue: for automatic location based on GeoClue2'
  'python-gobject: for gammastep-indicator'
  'python-xdg: for gammastep-indicator'
  'gtk3: for gammastep-indicator'
  'libappindicator-gtk3: for gammastep-indicator'
)
source=($url/-/archive/v$pkgver/$pkgname-v$pkgver.tar.gz)
sha512sums=('411e293ad233ee131d408c5e8b80c4b8be418570bb90497d33c501f4825bbfa0aa692abbd602db66635955a291cd1802748e52edcff6fd4d262e253d25e4e42b')
b2sums=('2a7e6c38bfee3e9a7bab3ecbb552e1b6452c6d65c95376a1c368016781721ae17fa570e4e88b2aa0380cbd745d696fdb537e1c69bdd8cda04311ecfcef903c6a')

prepare() {
  cd $pkgname-v$pkgver
  ./bootstrap
}

build() {
  cd $pkgname-v$pkgver
  ./configure --prefix=/usr
  make
}

check() {
  make -k check -C $pkgname-v$pkgver
}

package() {
  depends+=(
    glib2 libgio-2.0.so libgobject-2.0.so libglib-2.0.so
  )

  make DESTDIR="$pkgdir/" install -C $pkgname-v$pkgver
  install -vDm 644 $pkgname-v$pkgver/{CONTRIBUTING,NEWS,README}.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
