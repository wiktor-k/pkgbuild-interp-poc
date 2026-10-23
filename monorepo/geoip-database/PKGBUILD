# Maintainer: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Evangelos Foutras <foutrelis@archlinux.org>

pkgname=(geoip-database geoip-database-extra)
pkgver=20260204
pkgrel=1
arch=(any)
url='https://mailfud.org/geoip-legacy/'
license=('LicenseRef-GeoLite2-End-User-License-Agreement')
makedepends=('util-linux')
checkdepends=('geoip')
_source_base="https://mailfud.org/geoip-legacy"
source=(GeoIP-$pkgver.dat.gz::$_source_base/GeoIP.dat.gz
        GeoIPv6-$pkgver.dat.gz::$_source_base/GeoIPv6.dat.gz
        GeoIPCity-$pkgver.dat.gz::$_source_base/GeoIPCity.dat.gz
        GeoIPCityv6-$pkgver.dat.gz::$_source_base/GeoIPCityv6.dat.gz
        GeoIPASNum-$pkgver.dat.gz::$_source_base/GeoIPASNum.dat.gz
        GeoIPASNumv6-$pkgver.dat.gz::$_source_base/GeoIPASNumv6.dat.gz
        GeoLite2-End-User-License-Agreement.pdf)
sha256sums=('035a91a494ac26d584c6a583c1dbdf46db55596c635d804fded7cfc04e008e97'
            '9a465d0799f7940820f305f67df36d493fe1e540f18f97ed44c4f156573f1dad'
            '877ec77d5a128242438c91b6a3709d9d91f14f3f4d04b61ab0160a73d164b351'
            '20ce76951cd0d231527526fed23b4a6c9727abfb3e18797d578193e22b127104'
            'e2305cbc0258255619da9fac7d0d8094185dac61108c0caff896195bb6fde9ff'
            'b152f3b8593de6c9c65ec072fac67b4976bd37a1f94cf7fd8c6c00700494ff69'
            'b9c7fbfcdd5b682b5635592281d72006df89d73f163308b7ef35a59f293b036e')

prepare() {
  cd "$srcdir"
  rename -v -- "-$pkgver" '' *.dat
}

check() {
  cd "$srcdir"

  if [[ $(geoiplookup -f GeoIP.dat archlinux.org) != *'US, United States' ]]; then
    echo >&2 'Unable to resolve IPv4 address to country.'
    return 1
  fi

  if [[ $(geoiplookup6 -f GeoIPv6.dat archlinux.org) != *'US, United States' ]]; then
    echo >&2 'Unable to resolve IPv6 address to country.'
    return 1
  fi

  if [[ $(geoiplookup -f GeoIPCity.dat archlinux.org) != *'US, '* ]]; then
    echo >&2 'Unable to resolve IPv4 address to city.'
    return 1
  fi

  if [[ $(geoiplookup6 -f GeoIPCityv6.dat archlinux.org) != *'US, '* ]]; then
    echo >&2 'Unable to resolve IPv6 address to city.'
    return 1
  fi

  if [[ $(geoiplookup -f GeoIPASNum.dat archlinux.org) != *'AS11019 HAProxy Technologies'* ]]; then
    echo >&2 'Unable to resolve IPv4 address to ASN.'
    return 1
  fi

  if [[ $(geoiplookup6 -f GeoIPASNumv6.dat archlinux.org) != *'AS11019 HAProxy Technologies'* ]]; then
    echo >&2 'Unable to resolve IPv6 address to ASN.'
    return 1
  fi
}

package_geoip-database() {
  pkgdesc="GeoIP legacy country database (based on GeoLite2 data created by MaxMind)"
  install=$pkgname.install

  cd "$srcdir"
  install -Dm644 -t "$pkgdir/usr/share/GeoIP" GeoIP{,v6}.dat
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname" GeoLite2-End-User-License-Agreement.pdf
}

package_geoip-database-extra() {
  pkgdesc="GeoIP legacy city/ASN databases (based on GeoLite2 data created by MaxMind)"
  install=$pkgname.install

  cd "$srcdir"
  install -Dm644 -t "$pkgdir/usr/share/GeoIP" GeoIPCity{,v6}.dat GeoIPASNum{,v6}.dat
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname" GeoLite2-End-User-License-Agreement.pdf
}
