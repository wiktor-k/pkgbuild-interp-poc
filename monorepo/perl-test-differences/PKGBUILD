# This PKGBUILD was autocreated.  See https://xtermin.us/archperl/ for details
# Maintainer: Charles Mauch <cmauch@gmail.com>

pkgname=perl-test-differences
pkgver=0.73
pkgrel=1
pkgdesc="Perl/CPAN Module Test::Differences"
arch=(any)
depends=(perl-text-diff)
url="https://search.cpan.org/dist/Test-Differences"
license=("GPL" "PerlArtistic")
source=("https://search.cpan.org/CPAN/authors/id/D/DC/DCANTRELL/Test-Differences-$pkgver.tar.gz")
sha256sums=('3c113f4e25c38323a3fe5992d30c337264d00ef872aecfd2d9431ca8ea0083e2')
options=(!emptydirs)

build() {
  cd "$srcdir"/Test-Differences-$pkgver
  PERL_MM_USE_DEFAULT=1 perl Makefile.PL INSTALLDIRS=vendor  
  make 
}
package() {
  cd "$srcdir"/Test-Differences-$pkgver
  make install DESTDIR="$pkgdir" 
}
