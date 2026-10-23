# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
pkgname=box2d
pkgver=3.2.0
pkgrel=1
pkgdesc="2D rigid body simulation library for games"
url="http://www.box2d.org/"
license=('MIT')
arch=('x86_64')
depends=('glibc')
makedepends=('cmake' 'ninja' 'git' 'wayland' 'gtk3'
             'libxinerama' 'libxkbcommon' 'libxkbcommon-x11' 'libxt'
             'xorg-xrandr' 'libxcursor' 'xorg-xinput' 'glfw')
source=("$pkgname-$pkgver.tar.gz::https://github.com/erincatto/Box2D/archive/v${pkgver}.tar.gz")
sha512sums=('874bde300e7e33e950152f1d86cb4e2989cd7e30b2d89ac29d9596d9251b514972285eaba19b2d2941870f6d0e446265bf22e4d2a6df287f2c9807e0a82c8ed1')

build() {
  cd $pkgname-$pkgver

  cmake \
    -Bbuild \
    -GNinja \
    -DBUILD_SHARED_LIBS=ON \
    -DCMAKE_INSTALL_PREFIX=/usr
  ninja -C build
}

package() {
  cd $pkgname-$pkgver

  DESTDIR="$pkgdir" ninja -C build install

  cp -r include "$pkgdir"/usr/
  install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/box2d/LICENSE
}

# vim: sw=2 ts=2 et:
