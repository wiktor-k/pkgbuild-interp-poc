# Maintainer:
# Contributor: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Maato <maato softwarebakery com>

pkgname=volumeicon
pkgver=0.5.1
pkgrel=8
pkgdesc='Volume control for the system tray'
arch=(x86_64)
url='https://github.com/Maato/volumeicon'
license=(GPL-3.0-only)
depends=(alsa-lib gtk3 libnotify)
makedepends=(git intltool)
source=("git+$url#tag=$pkgver"
        "notify.patch::$url/commit/78a9c23775ebe12d3f3be93c0798144a6e732d90.patch"
        volumeicon.desktop)
b2sums=('adb2f3f0419fbccc8df11949fb16e55987fb2aac46870540672b0631c16ddd27b31b70022946ddbaea867d4879752a57790cf5474f16f068e628004df8ec265b'
        '899eba04b6959e50807e6728aaffe3c144e484514d9cf3c79e4f886916fe4b6a63847074395a02eb158c6a27f23e8277541438a869f8e883b5323ca29e51a547'
        '56333c46320108873d2c229c6904f02dc4bfc93c549a2a906cb0efd7e63f569b2310a7178b91cfc46d3246aa366dfc12d6b42d5f960c9f7d0b026d7afdeacdac')

prepare() {
  cd $pkgname
  patch -Np1 -i ../notify.patch
  ./autogen.sh
}

build() {
  cd $pkgname
  ./configure --prefix=/usr --enable-notify
  make
}

package() {
  make -C $pkgname DESTDIR="$pkgdir" install
  install -Dm644 volumeicon.desktop "$pkgdir/usr/share/applications/volumeicon.desktop"
}
