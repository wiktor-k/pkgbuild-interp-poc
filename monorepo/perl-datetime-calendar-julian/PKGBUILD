# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Rémy Oudompheng <remy@archlinux.org>

pkgname=perl-datetime-calendar-julian
_dist=DateTime-Calendar-Julian
pkgver=0.108
pkgrel=1
pkgdesc='Perl module for Dates in the Julian calendar'
arch=(any)
url="https://metacpan.org/release/${_dist}"
license=('GPL-1.0-or-later OR Artistic-1.0-Perl')
depends=(perl-datetime)
options=(!emptydirs)
source=(https://search.cpan.org/CPAN/authors/id/W/WY/WYANT/${_dist}-${pkgver}.tar.gz)
sha512sums=('a8cd9467f76f51214e75e2c7149a7aac7b562f0e59a97406ad353d7d0ec58d374628228454bdb5e7a63ac5213b703974dc398ed87938653ad4fbb028a8a0aac7')
b2sums=('f6c998a070107e6403259da8a6ebd0815fb56db1280c5e98aa3693fb5cda786a501da12dacfeb8167188b0ffeca02a5b2c0ed16f5f5f4d9a1f6c09642b9c6741')

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
