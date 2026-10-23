# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>

pkgname=perl-specio
pkgver=0.53
pkgrel=1
pkgdesc="Type constraints and coercions for Perl"
arch=('any')
url="https://search.cpan.org/dist/Specio"
license=('GPL' 'PerlArtistic')
depends=('perl-role-tiny' 'perl-test-fatal' 'perl-mro-compat' 'perl-devel-stacktrace'
	 'perl-clone' 'perl-eval-closure' 'perl-clone-choose')
options=('!emptydirs')
source=(https://search.cpan.org/CPAN/authors/id/D/DR/DROLSKY/Specio-$pkgver.tar.gz)
sha256sums=('0d0eecfb9e89bd0f5f710fac42e1200a882d513a862f98497eaef5927ac6c183')

build() {
  cd  "$srcdir"/Specio-$pkgver
  PERL_MM_USE_DEFAULT=1 perl Makefile.PL INSTALLDIRS=vendor 
  make
}
package() {
  cd  "$srcdir"/Specio-$pkgver
  make install DESTDIR="$pkgdir"
  find "$pkgdir" -name '.packlist' -delete
  find "$pkgdir" -name '*.pod' -delete
}
