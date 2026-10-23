# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=perl-alien-cmake3
pkgver=0.10
pkgrel=4
pkgdesc="Find or download or build cmake 3 or better"
arch=('any')
license=('PerlArtistic')
url="https://metacpan.org/release/Alien-cmake3"
depends=('cmake' 'perl-alien-build')
makedepends=('perl-html-parser' 'perl-sort-versions' 'perl-uri' 'perl-io-socket-ssl' 'perl-mozilla-ca')
options=('!emptydirs')
source=("https://cpan.metacpan.org/authors/id/P/PL/PLICEASE/Alien-cmake3-$pkgver.tar.gz")
sha512sums=('f1db909148461bd2da51f4769d018b1ba2ec543366b25cf414a2f7cdfc275db98b4bfa82bce9b82c5b31e01c216eaedd21d8fdc61547cc9086322f23f03e0b56')

build() {
  cd Alien-cmake3-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd Alien-cmake3-$pkgver
  make test
}

package() {
  cd Alien-cmake3-$pkgver
  make DESTDIR="$pkgdir" install
}
