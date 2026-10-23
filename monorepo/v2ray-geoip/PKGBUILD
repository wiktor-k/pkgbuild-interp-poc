# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: DeepChirp <deepchirp@archlinux.org>

pkgname=v2ray-geoip
pkgver=202610050438
pkgrel=1
pkgdesc="GeoIP List for V2Ray"
arch=('any')
url="https://github.com/v2fly/geoip"
license=('CC-BY-SA-4.0')
source=("geoip-$pkgver.dat::https://github.com/v2fly/geoip/releases/download/$pkgver/geoip.dat")
sha512sums=('3dcf61d6abcd6d49b7020435e1a38129a13cacce6b5be2b5ae9ccb029c2165f301cdc2bc7673bec50cb32cb361edcf642015f122f5a8547fab4454d91bbea186')

package() {
  install -Dm644 geoip-$pkgver.dat "$pkgdir"/usr/share/v2ray/geoip.dat
}
