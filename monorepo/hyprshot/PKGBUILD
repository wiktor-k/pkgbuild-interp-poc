# Maintainer: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Gustavo Parreira <gustavotcparreira at gmail dot com>
# Contributor: xiota

pkgname=hyprshot
pkgver=1.3.0
pkgrel=4
url='https://github.com/Gustash/Hyprshot'
pkgdesc='Hyprland screenshot utility'
license=(GPL-3.0-only)
arch=(any)
depends=(bash grim jq libnotify slurp wl-clipboard)
optdepends=('hyprpicker: to freeze the screen while taking the screenshot')
source=("$pkgname-$pkgver.tar.gz::$url/archive/$pkgver.tar.gz")
b2sums=('90711199532c8ae5618ce0228602fb7993556b1a7f2e6fc5baa1e49da81ec1e8e365550c2b4fcbc676e90cdc9ae9244ed1a356641160bd09f42affccb1c08378')

package() {
  install -Dm755 Hyprshot-$pkgver/$pkgname "$pkgdir/usr/bin/$pkgname"
}
