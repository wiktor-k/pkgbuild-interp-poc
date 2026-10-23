# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=debootstrap
pkgver=1.0.145
pkgrel=1
pkgdesc="Bootstrap a basic Debian system"
arch=('any')
license=('MIT')
url="https://tracker.debian.org/pkg/debootstrap"
depends=('distro-info' 'perl' 'wget')
makedepends=('git')
source=(git+https://salsa.debian.org/installer-team/debootstrap.git#tag=$pkgver
        arch-detect.patch)
optdepends=('gnupg: check release signatures'
            'debian-archive-keyring: check release signatures for Debian'
            'debian-ports-archive-keyring: check release signatures for Debian Ports'
            'ubuntu-keyring: check release signatures for Ubuntu')
sha512sums=('c926ee4396d00a7a5d4363259ffcbda4997bc835bb574c9408ba1aa4ddea608285d5cc51e7547f51f6fa1e1612e9c3bedd37061efa689db16d8dce560d251c25'
            'b1fdff6142326c2a5075e457181a85cb2a37848c777f128f4c0eb3740980784d25d4f143a76b97c90ef252463cc5d8bd0797b2544c27b6dd8e05a6258a6b0c02')

prepare() {
  cd debootstrap

  sed -i 's/sbin/bin/g' Makefile

  # Debian defaults
  sed -i 's|export PATH|export PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"|' debootstrap

  # Detect Architecture
  patch -p0 -i ../arch-detect.patch
}

package() {
  cd debootstrap
  make DESTDIR="$pkgdir" install

  install -Dm644 debootstrap.8 "$pkgdir"/usr/share/man/man8/debootstrap.8
}
