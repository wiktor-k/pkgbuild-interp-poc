# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgbase=caja-extensions
pkgname=(
  caja-extensions-common
  caja-audio-video-properties
  caja-image-converter
  caja-open-terminal
  caja-sendto
  caja-share
  caja-wallpaper
  caja-xattr-tags
)
pkgver=1.28.0
pkgrel=4
pkgdesc="Set of extensions for Caja, the MATE file manager"
arch=(x86_64)
url='https://github.com/mate-desktop/caja-extensions'
license=(GPL-2.0-or-later)
depends=(
  caja
  dbus-glib
  gst-plugins-base-libs
  gupnp
)
makedepends=(
  git
  glib2-devel
  gtk-doc
  mate-common
)
groups=(mate-extra)
source=("git+https://github.com/mate-desktop/caja-extensions.git#tag=v$pkgver")
b2sums=(d9ec75038f763c4c5f9522b5cc92297d7fb92aca8da68c532c6281416d045ecbab046d205520968a0625e70ff1fe3441567b24b2bf8daea6a9e15baa547519df)

prepare() {
  cd $pkgbase
  autoreconf -fiv
}

build() {
  cd $pkgbase
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --enable-gtk-doc \
    --disable-gksu
  make
}

package_caja-extensions-common() {
  pkgdesc="Caja extensions (common files)"
  depends=()
  groups=()

  cd $pkgbase
  make -C po DESTDIR="$pkgdir" install
}

package_caja-audio-video-properties() {
  pkgdesc="A Caja extension to see information about audio or video files"
  depends=(
    caja
    caja-extensions-common
    glib2
    glibc
    gst-plugins-base
    gst-plugins-base-libs
    gstreamer
    gtk3
  )

  cd $pkgbase
  make -C audio-video-properties DESTDIR="$pkgdir" install
}

package_caja-image-converter() {
  pkgdesc="A Caja extension for simple image conversions"
  depends=(
    caja
    caja-extensions-common
    glib2
    glibc
    gtk3
    imagemagick
  )

  cd $pkgbase
  make -C image-converter DESTDIR="$pkgdir" install
}

package_caja-open-terminal() {
  pkgdesc="A Caja extension for opening terminals in arbitrary local paths"
  depends=(
    caja
    caja-extensions-common
    dconf
    glib2
    glibc
    gtk3
    mate-desktop
  )

  cd $pkgbase
  make -C open-terminal DESTDIR="$pkgdir" install
}

package_caja-sendto() {
  pkgdesc="A Caja extension for sending files"
  depends=(
    caja
    caja-extensions-common
    dbus-glib
    dconf
    gdk-pixbuf2
    glib2
    glibc
    gssdp
    gtk3
    gupnp
    libgcc
  )

  cd $pkgbase
  make -C sendto DESTDIR="$pkgdir" install
}

package_caja-share() {
  pkgdesc="A Caja extension to quickly share a folder"
  depends=(
    caja
    caja-extensions-common
    glib2
    glibc
    gtk3
    samba
  )

  cd $pkgbase
  make -C share DESTDIR="$pkgdir" install
}

package_caja-wallpaper() {
  pkgdesc="A Caja extension to quickly set wallpaper"
  depends=(
    caja
    caja-extensions-common
    glib2
    glibc
  )

  cd $pkgbase
  make -C wallpaper DESTDIR="$pkgdir" install
}

package_caja-xattr-tags() {
  pkgdesc="A Caja extension to see tags stored on xattrs"
  depends=(
    caja
    caja-extensions-common
    glib2
    glibc
  )

  cd $pkgbase
  make -C xattr-tags DESTDIR="$pkgdir" install
}
