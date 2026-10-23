# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: John D Jones III AKA jnbek <jnbek1972 -_AT_- g m a i l -_Dot_- com>
# Generator  : CPANPLUS::Dist::Arch 1.32

pkgname=perl-namespace-autoclean
pkgver=0.31
pkgrel=3
pkgdesc="Keep imports out of your namespace"
arch=('any')
license=('PerlArtistic' 'GPL')
url='https://metacpan.org/release/namespace-autoclean'
depends=('perl-b-hooks-endofscope' 'perl-namespace-clean')
checkdepends=('perl-test-requires' 'perl-moo' 'perl-sub-name' 'perl-mouse' 'perl-test-needs')
options=('!emptydirs')
source=("https://search.cpan.org/CPAN/authors/id/E/ET/ETHER/namespace-autoclean-$pkgver.tar.gz")
sha512sums=('16f1d473ed87d196812867306996284d1f6817603399e1d67273c2e1c58ba2ba811c84c5e1762675526bfcb808aeece7e3180de189ca9ba951c6b9ff2023d802')

build() {
  cd namespace-autoclean-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd namespace-autoclean-$pkgver
  make test
}

package() {
  cd namespace-autoclean-$pkgver
  make DESTDIR="$pkgdir"/ install
}

# vim:set ts=2 sw=2 et:
