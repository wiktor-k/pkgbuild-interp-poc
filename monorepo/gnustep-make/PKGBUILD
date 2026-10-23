# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Vesa Kaihlavirta <vegai@iki.fi>
# Contributor: Sebastian Sareyko <public@nooms.de>

pkgname=gnustep-make
pkgver=2.9.3
pkgrel=1
pkgdesc='The GNUstep make package'
arch=('x86_64') # See FS#29546
url='https://www.gnustep.org/'
license=('GPL-3.0-or-later')
makedepends=('git')
groups=('gnustep-core')
backup=('etc/GNUstep/GNUstep.conf')
source=("$pkgname::git+https://github.com/gnustep/tools-make#tag=make-${pkgver//./_}")
sha512sums=('102193e1eef30905be5e1b1de969a244b549a4a0545661155b094e5bb25ec614360fc2cff00bff791421db9814a65f747469d21c717cf936461cf57ae4ea3bcf')
b2sums=('a2d45b87b7a79c80dc4e0fc9aaccd024b1ac81b49c4116e965b0fc3bbbd1bb967cffe1e87b0d5cfb98a5448172932def27076c067b444bc191790bef75a8abba')

prepare() {
  cd "$pkgname"

  for i in GNUSTEP_SYSTEM_ADMIN_TOOLS GNUSTEP_NETWORK_ADMIN_TOOLS GNUSTEP_LOCAL_ADMIN_TOOLS; do
    sed -i "s|@$i@|/usr/bin|g" GNUstep{,-strict-v2}.conf.in filesystem.{sh,csh,make}.in
  done
}

build() {
  cd "$pkgname"

  ./configure \
    --prefix=/usr \
    --sbindir=/usr/bin \
    --sysconfdir=/etc/GNUstep \
    --enable-native-objc-exceptions
  make
}

package() {
  cd "$pkgname"

  make DESTDIR="$pkgdir" install
}

