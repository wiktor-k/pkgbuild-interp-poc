# Maintainer: Evangelos Foutras <foutrelis@archlinux.org>
# Contributor: 3ED <krzysztof1987 at gmail dot com>

pkgname=perl-parallel-forkmanager
pkgver=2.04
pkgrel=2
pkgdesc="A simple parallel processing fork manager"
arch=('any')
url="https://metacpan.org/release/Parallel-ForkManager"
license=('PerlArtistic' 'GPL')
depends=('perl' 'perl-moo')
checkdepends=('perl-test-warn')
options=('!emptydirs')
source=(https://cpan.metacpan.org/authors/id/Y/YA/YANICK/Parallel-ForkManager-$pkgver.tar.gz)
sha256sums=('606894fc2e9f7cd13d9ec099aaac103a8f0943d1d80c2c486bae14730a39b7fc')

build() {
  cd Parallel-ForkManager-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd Parallel-ForkManager-$pkgver
  make test
}

package() {
  cd Parallel-ForkManager-$pkgver
  make install DESTDIR="$pkgdir"
}

# vim:set ts=2 sw=2 et:
