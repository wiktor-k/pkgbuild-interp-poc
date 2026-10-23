# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: mortzu <me@mortzu.de>
# Contributor: fnord0 <fnord0@riseup.net>

pkgname=acpi_call
pkgver=1.2.2
pkgrel=395
pkgdesc='A linux kernel module that enables calls to ACPI methods through /proc/acpi/call'
url=https://github.com/nix-community/acpi_call
arch=(x86_64)
license=(GPL-3.0-only)
makedepends=(
  acpi_call-dkms=${pkgver}
  linux-headers
)
conflicts=(acpi_call-dkms)

build() {
  _kernver=$(</usr/src/linux/version)

  fakeroot dkms build --dkmstree "$srcdir" -m acpi_call/$pkgver -k $_kernver
}

package() {
  depends=(linux)

  _kernver=$(</usr/src/linux/version)

  install -Dt "$pkgdir/usr/lib/modules/$_kernver/extramodules" -m0644 \
    acpi_call/${pkgver}/$_kernver/$CARCH/module/*.ko

  # compress each module individually
  find "${pkgdir}" -name '*.ko' -exec zstd --rm -19 {} +

  echo acpi_call | install -Dm644 /dev/stdin "$pkgdir/usr/lib/modules-load.d/$pkgname.conf"

  mkdir -p "$pkgdir/usr/share"
  cp -a /usr/share/acpi_call "$pkgdir/usr/share/$pkgname"
}
