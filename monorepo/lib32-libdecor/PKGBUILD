# Maintainer:

pkgdesc='Client-side decorations library for Wayland clients'
pkgname=lib32-libdecor
pkgver=0.2.5
pkgrel=1
url="https://gitlab.gnome.org/jadahl/libdecor"
license=(MIT)
arch=(x86_64)
depends=(libdecor lib32-{wayland,pango,dbus,gtk3,glibc,cairo,glib2})
makedepends=(cmake ninja meson wayland-protocols)
source=("https://gitlab.freedesktop.org/libdecor/libdecor/-/archive/${pkgver}/libdecor-${pkgver}.tar.gz")
sha512sums=('c696ed38a80256e1db68fcc0fd8a8c18439c4a1c4bf4d7e4068a741b0e2642d99dc07ae92e209c15f2387ea9330f9b3281ab9114b5d8989e085165f3eae6abed')

build() {
  export CC='gcc -m32'
  export CXX='g++ -m32'
  export PKG_CONFIG=i686-pc-linux-gnu-pkg-config

  meson -Ddemo=false build libdecor-$pkgver \
    --prefix=/usr --libdir=/usr/lib32
  meson compile -C build
}

package() {
  meson install -C build --destdir "$pkgdir"
  rm -fr "$pkgdir"/usr/include
  install -Dm644 libdecor-$pkgver/LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
