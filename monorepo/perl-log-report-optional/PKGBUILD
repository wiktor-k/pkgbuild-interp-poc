# CPAN Name  : Log::Report::Optional
# Contributor: Anton Leontiev <scileont /at/ gmail.com>
# Generator  : CPANPLUS::Dist::Arch 1.32

pkgname=perl-log-report-optional
pkgver=1.08
pkgrel=1
pkgdesc='Perl package providing Log::Report or ::Minimal'
arch=('any')
url='https://metacpan.org/release/Log-Report-Optional'
license=('PerlArtistic' 'GPL')
depends=('perl' 'perl-string-print')
options=(!emptydirs)
source=(http://search.cpan.org/CPAN/authors/id/M/MA/MARKOV/Log-Report-Optional-$pkgver.tar.gz)
sha256sums=('77b248d4cf7fecaa7e865930e72df0b9d5b333358d00c5bd45e2c71d5df113ad')

sanitize() {
	unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
	export PERL_MM_USE_DEFAULT=1
}

build() {
	cd Log-Report-Optional-$pkgver
	sanitize
	perl Makefile.PL INSTALLDIRS=vendor
	make
}

check() {
	cd Log-Report-Optional-$pkgver
	sanitize
	make test
}

package() {
	cd Log-Report-Optional-$pkgver
	sanitize
	make install DESTDIR="$pkgdir"
	find "$pkgdir" \( -name .packlist -o -name perllocal.pod \) -delete
}
