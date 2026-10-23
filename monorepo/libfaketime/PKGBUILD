# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Contributor: Dan Serban
# Contributor: Ross Melin <rdmelin@gmail.com>

pkgname=libfaketime
pkgver=0.9.13
pkgrel=1
pkgdesc='Report fake dates and times to programs without having to change the system-wide time.'
arch=('x86_64')
url='https://github.com/wolfcw/libfaketime'
license=('GPL-2.0-only')
depends=('glibc'
         'libgcc')
source=("https://github.com/wolfcw/$pkgname/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('8e56deeb805682b025107e095f1d94a6ea677472f05824cd475f5c5e6e1a5ddf')
options=(!lto)

build() {
  cd "${pkgname}-${pkgver}"
  make PREFIX=/usr
}

package() {
  cd "${pkgname}-${pkgver}"
  make PREFIX=/usr DESTDIR="${pkgdir}" install
}
