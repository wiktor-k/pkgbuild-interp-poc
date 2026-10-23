# CPAN Name  : String::Print
# Contributor: Anton Leontiev <scileont /at/ gmail.com>
# Generator  : CPANPLUS::Dist::Arch 1.32

pkgname=perl-string-print
pkgver=1.02
pkgrel=1
pkgdesc='Perl package providing printf() alternative'
arch=('any')
url='https://metacpan.org/release/String-Print'
license=('PerlArtistic' 'GPL')
depends=('perl>=5.10.1'
	'perl-html-parser'
	'perl-timedate>=2.30'
	'perl-unicode-linebreak')
options=(!emptydirs)
source=(http://search.cpan.org/CPAN/authors/id/M/MA/MARKOV/String-Print-$pkgver.tar.gz)
sha256sums=('3049536486459e38e1d791c07ce022326a91a302beaf01dcdb0e7b703a5da6cc')

sanitize() {
	unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
	export PERL_MM_USE_DEFAULT=1
}

build() {
	cd String-Print-$pkgver
	sanitize
	perl Makefile.PL INSTALLDIRS=vendor
	make
}

check() {
	cd String-Print-$pkgver
	sanitize
	make test
}

package() {
	cd String-Print-$pkgver
	sanitize
	make install DESTDIR="$pkgdir"
	find "$pkgdir" \( -name .packlist -o -name perllocal.pod \) -delete
}
