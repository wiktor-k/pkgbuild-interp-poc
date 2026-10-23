# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Jan de Groot <jgc@archlinux.org>

pkgname=gnome-applets
pkgver=3.58.0
pkgrel=2
pkgdesc='Small applications for the GNOME Panel'
arch=(x86_64)
url='https://wiki.gnome.org/Projects/GnomeApplets'
license=(GPL-2.0-or-later)
depends=(
  at-spi2-core
  cairo
  cpupower
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gnome-panel
  gsettings-desktop-schemas
  gtk3
  hicolor-icon-theme
  libgtop
  libgweather-4
  libnotify
  libwnck3
  libx11
  libxml2
  pango
  polkit
  tinysparql
  upower
)
makedepends=(
  autoconf-archive
  git
  glib2-devel
  yelp-tools
)
optdepends=('localsearch: Search bar')
source=("git+https://gitlab.gnome.org/GNOME/$pkgname.git?signed#tag=$pkgver")
b2sums=(feba46e84504d10e9dbfddad6545b05dfce8a477f5bf38d584c4d2bb032661adbd36f67ab0c82d278cc5f0550fdde7b5d7be75020238af5c7ebf7642d1022b6f)
validpgpkeys=(7B44FD78E49334EC10B3B288A3D013EC303E1894) # Alberts Muktupāvels <alberts.muktupavels@gmail.com>

prepare() {
  cd $pkgname

  # Destroy preferences window on dispose
  # https://gitlab.gnome.org/GNOME/gnome-applets/-/merge_requests/98
  git cherry-pick -n 4219145b6e71a47302c1b09d55424ade5f4ae022 \
                     03141f7d4b593f9a28967304fb9e836244624b99 \
                     1279a018b9870ca6f28656883a40a133dab5bbfa \
                     bbc0d580d51022b93b97b148d27a70fecb6aaaae \
                     22c5f0b627e072dc0f3207db4112dd5bdc590fb9 \
                     aa80c3a0222146c374f43f455340a4e57b5c9d46 \
                     9707bc82063ee78e1d44f5797bacb49a46662238 \
                     d832cf5dd981bb7ca8fe3d80c4b6ff2014ae32e4

  # Fix reopening note properties after closing the window
  # https://gitlab.gnome.org/GNOME/gnome-applets/-/merge_requests/107
  git cherry-pick -n 465eddf9909e2eb152c652d1e19a6561ef9ffbc5

  # Fix initial display of gweather and mini-commander applets
  # https://gitlab.gnome.org/GNOME/gnome-applets/-/merge_requests/108
  git cherry-pick -n 79709a65f97e3cdac945f8c6cfe2c684c699e318 \
                     78c86881ea9df42539f61976bbc4a7846ad34916

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --disable-debug \
    --enable-compile-warnings
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
