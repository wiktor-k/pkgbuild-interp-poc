# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=nitrokey-udev-rules
pkgver=1.1.0
pkgrel=1
pkgdesc="Udev rules for Nitrokey devices"
arch=(any)
url="https://github.com/Nitrokey/nitrokey-udev-rules"
license=(CC0-1.0)
source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz")
sha512sums=('6b4e985d00ae29fa68a5e1ada8401d6df3956e6e228ec50b27c4ca1ae70ec16d161f8533106d1fb88ffc5ac115175bc3fbdabef664510f8df20ba220c9bf1cdc')
b2sums=('7df2b8e3d51f1d2db8480807568c54065e85fb7b222e8d62e13fc1cbcb30bd9d87f31e02e786ea2e7fd170c062a20ce5da1a1ed52b43c467bf5e676c87aabe9f')

package() {
  install -vDm 644 $pkgname-$pkgver/41-nitrokey.rules -t "$pkgdir/usr/lib/udev/rules.d/"
  install -vDm 644 $pkgname-$pkgver/CHANGELOG.md -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 $pkgname-$pkgver/LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
