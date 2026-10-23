# Maintainer: Rémy Oudompheng <remy@archlinux.org>
# Contributor: John D Jones III <j[nospace]n[nospace]b[nospace]e[nospace]k[nospace]1972 -_AT_- the domain name google offers a mail service at ending in dot com>
# Generator  : CPANPLUS::Dist::Arch 1.25

pkgname='perl-business-issn'
_dist=Business-ISSN
pkgver=1.008
pkgrel=3
pkgdesc="Perl package to work with International Standard Serial Numbers"
arch=('any')
url="https://metacpan.org/release/${_dist}"
license=('Artistic2.0')
depends=('perl')
checkdepends=('perl-test-pod' 'perl-test-pod-coverage')
options=('!emptydirs')
source=(https://search.cpan.org/CPAN/authors/id/B/BR/BRIANDFOY/${_dist}-${pkgver}.tar.gz)
sha256sums=('b16b3a1b0e53cd45ed3328906d33ad4d59a13b57abf341424553aecf3e443aac')

build() {
    cd "${srcdir}/${_dist}-${pkgver}"
    unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
    export PERL_MM_USE_DEFAULT=1 PERL_AUTOINSTALL=--skipdeps
    perl Makefile.PL
    make
}

check() {
    cd "${srcdir}/${_dist}-${pkgver}"
    unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
    export PERL_MM_USE_DEFAULT=1
    make test
}

package() {
    cd "${srcdir}/${_dist}-${pkgver}"
    unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
    make install INSTALLDIRS=vendor DESTDIR="$pkgdir"
}
