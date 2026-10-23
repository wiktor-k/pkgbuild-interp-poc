# Maintainer: Jiachen Yang <farseerfc@gmail.com>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Vlad M. <vlad@archlinux.net>

pkgname=toxic
pkgver=0.16.3
pkgrel=1
pkgdesc='CLI Tox client'
arch=('x86_64')
url='https://github.com/JFreegman/toxic'
license=('GPL-3.0-or-later')
depends=(
  'curl'
  'freealut'
  'glibc'
  'libconfig'
  'libnotify'
  'libpng'
  'libvpx'
  'libx11'
  'ncurses'
  'openal'
  'qrencode'
  'toxcore'
)
makedepends=('git')
source=("git+https://github.com/JFreegman/$pkgname?signed#tag=v$pkgver")
sha512sums=('2e3415127af7bf8a2d5614ccd1213b5078dd9c9a5ef59f1500cb587d4a284daab48a2e20d226b4ea12eab2bc833f0a11edd3ea0fdc271a66a5cf322a63d1b25d')
validpgpkeys=('BABD00573A065BFA90D53D563627F3144076AE63') # Jfreegman <jfreegman@gmail.com>

build() {
  cd "$pkgname"
  make PREFIX=/usr DISABLE_GAMES=1
}

package() {
  cd "$pkgname"
  make PREFIX=/usr DESTDIR="$pkgdir" install
}
