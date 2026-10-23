# Maintainer: David Runge <dave@sleepmap.de>
# Contributor: Ray Rashif <schiv@archlinux.org>
# Contributor: James Rayner <james@archlinux.org>

pkgname=schedtool
pkgver=1.3.0
pkgrel=8
pkgdesc="Query or alter a process' scheduling policy"
arch=(x86_64)
url="https://github.com/freequaos/schedtool"
license=(GPL-2.0-only)
groups=(realtime)
depends=(glibc)
source=(
  $url/archive/$pkgname-$pkgver.tar.gz
  $pkgname-1.3.0-fix-type-mismatch.patch
)
sha512sums=('71b7bbc91ee86a83639b179c8cce70e0fd3acca3af3cbc51880cf1f177fdd901e1fb6d0ecea2aa40761608589e7b903431ac97d5046f9ec50535bad691af2964'
            '1eff715f3d414b6f5434272adc72f0ef0670bff2d14e918ef712acc60f6c615dd45ee91ed290eccdc8bb51a040618a2f587574f7fa95fd4ca0326433bc3f1ba3')
b2sums=('45e9f0c84b99be83be4e250c71a3ef43e323efff565a89a8e69960b8cb703faa7df94fbfbd3a11863e4ecb54cc07acca8cd7803e21aaafe59cc0e923d1117025'
        '1e6a0152feb742498d906660e438f10ebdd562c5fcae7b6adb74e86475512e664da0353f4b66429d10e620dac0910c31cdb33d0722c1da55481e657de7558885')

prepare() {
  # Fix type mismatch: https://github.com/freequaos/schedtool/pull/2
  patch -Np1 -d $pkgname-$pkgname-$pkgver -i ../$pkgname-1.3.0-fix-type-mismatch.patch

  cd $pkgname-$pkgname-$pkgver
  # fix install location of docs
  sed -e 's/(RELEASE)/(TARGET)/g' -i Makefile
}

build() {
  make -C $pkgname-$pkgname-$pkgver
}

package() {
  make DESTPREFIX="$pkgdir/usr" install -C  $pkgname-$pkgname-$pkgver
}

# vim:set ts=2 sw=2 et:
