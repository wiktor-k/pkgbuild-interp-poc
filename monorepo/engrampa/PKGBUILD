# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Brad Fanella <cesura@archlinux.org>
# Contributor: Martin Wimpress <code@flexion.org>

pkgname=engrampa
pkgver=1.28.5
pkgrel=3
pkgdesc="A file archiver for MATE"
arch=(x86_64)
url='https://github.com/mate-desktop/engrampa'
license=(GPL-2.0-or-later)
depends=(
  cairo
  dconf
  gdk-pixbuf2
  glib2
  glibc
  gtk3
  gzip
  hicolor-icon-theme
  json-glib
  libice
  libsm
  pango
  tar
  unzip
  zip
)
makedepends=(
  caja
  git
  glib2-devel
  mate-common
  yelp-tools
)
optdepends=(
  '7zip: 7Z and ARJ archive support'
  'brotli: Brotli compression support'
  'caja: File archiver extension'
  'cpio: RPM support'
  'rpm-tools: RPM support'
  'unace: ACE archive support'
  'unrar: RAR archive support'
)
groups=(mate-extra)
source=(
  "git+https://github.com/mate-desktop/engrampa.git#tag=v$pkgver"
  git+https://github.com/mate-desktop/mate-submodules.git
  engrampa-fix-window-size.patch
)
b2sums=(
  a9023f79df4eff7afc5856d70bf8a9955e5a14e38f4102220b872e47d774169aec9957fd136057c659bf7ad25de5191295d0d7d08a4a98ca3124e73127c097b4
  SKIP
  4959e03b454983550d2c5fa96766ffe515ec0045bc70c832240ce1fe56b31f2fbb1bfcacae7bb63bbb460f86ec17cf3b72e97918c77a427924ce35a6feb7f74e
)

prepare() {
  cd $pkgname

  git submodule init
  git config submodule.mate-submodules.url "$srcdir/mate-submodules"
  git -c protocol.file.allow=always submodule update

  # Fix window size growing with client-side decorations
  # https://github.com/mate-desktop/engrampa/pull/559
  git apply -3 ../engrampa-fix-window-size.patch

  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --libexec="/usr/lib/$pkgname" \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --disable-packagekit
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
