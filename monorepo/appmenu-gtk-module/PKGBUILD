# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: rilian-la-te <ria.freelander@gmail.com>
# Contributor: Kaizhao Zhang <zhangkaizhao@gmail.com>

pkgname=appmenu-gtk-module
pkgver=25.04
pkgrel=3
pkgdesc="Application Menu GTK Module"
depends=(bash
         dconf
         gdk-pixbuf2
         glib2
         glibc
         libgcc
         gtk3)
makedepends=(git
             meson)
url="https://gitlab.com/vala-panel-project/vala-panel-appmenu/"
arch=(x86_64)
license=(LGPL-3.0-only)
source=(git+https://gitlab.com/vala-panel-project/vala-panel-appmenu#tag=$pkgver
        80-appmenu-gtk-module.sh)
sha256sums=('843b24f98f02feb8cbfdda26630018ae95f8ac4959de9deb88cf1a13506f845f'
            'e8aff88851145085e24989cc99da9f631c593b506d0224225f3eb347d16c6455')

build() {
  meson build vala-panel-appmenu/subprojects/$pkgname \
    --prefix=/usr \
    -D gtk=3
  ninja -C build
}

package() {
  DESTDIR="$pkgdir" ninja install -C build
  install -Dm755 80-appmenu-gtk-module.sh -t "$pkgdir"/etc/X11/xinit/xinitrc.d/
}
