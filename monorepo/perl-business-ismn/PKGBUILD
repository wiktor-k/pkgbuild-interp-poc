# Maintainer: Rémy Oudompheng <remy@archlinux.org>
# Contributor: John D Jones III <jnbek1972 -_AT_- g m a i l -_Dot_- com>
# Generator  : CPANPLUS::Dist::Arch 1.29

pkgname='perl-business-ismn'
_dist=Business-ISMN
pkgver=1.205
pkgrel=3
pkgdesc="Perl package to work with International Standard Music Numbers"
arch=('any')
url="https://metacpan.org/release/${_dist}"
license=('Artistic-2.0')
depends=('perl' 'perl-tie-cycle')
checkdepends=('perl-test-pod' 'perl-test-pod-coverage')
options=('!emptydirs')
source=("https://search.cpan.org/CPAN/authors/id/B/BR/BRIANDFOY/${_dist}-${pkgver}.tar.gz")
sha512sums=('b1686a610e926cd0bb8d138ed755a6eb971fe62e3ce9beff9f872602d021363cd2424b4724ff806bb4280c53f454f8280f308600f3de283e4b5c24e444c4e80c')
b2sums=('c47143e41481f21ee8668f2098d78c50a9934fc6a937a1ff1f26b4d552cd505a03c02bbe07278aeae5b370f60999c2bf716271344fa96ba6f1bd2ecb9ce148d4')

build() {
  cd "${_dist}-${pkgver}"

  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1 PERL_AUTOINSTALL=--skipdeps
  perl Makefile.PL
  make
}

check() {
  cd "${_dist}-${pkgver}"

  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1
  make test
}

package() {
  cd "${_dist}-${pkgver}"

  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  make install INSTALLDIRS=vendor DESTDIR="$pkgdir"
}
