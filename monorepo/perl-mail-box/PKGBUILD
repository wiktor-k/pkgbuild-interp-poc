# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Maintainer: Charles Mauch <cmauch@gmail.com>
# Contributor: Francois Charette <firmicus@gmx.net>

pkgname=perl-mail-box
pkgver=4.02
pkgrel=1
pkgdesc="Mail folder manager and MUA backend"
arch=(any)
url="https://search.cpan.org/dist/Mail-Box"
license=('Artistic-1.0-Perl' 'GPL-1.0-or-later')
depends=('perl-timedate' 'perl-digest-hmac' 'perl-file-remove>=0.20'
	 'perl-io-stringy' 'perl-mime-types>=1.004' 'perl-mailtools'
	 'perl-object-realize-later>=0.14' 'perl-uri>=1.23'
	 'perl-devel-globaldestruction' 'perl-user-identity>=0.90'
	 'perl-mail-message')
options=('!emptydirs')
source=(https://cpan.metacpan.org/authors/id/M/MA/MARKOV/Mail-Box-$pkgver.tar.gz)
sha256sums=('8f064df0ed089f47f1922c11ce7468784f037b351968add191692a614c9d5f25')

prepare() {
  cd  "$srcdir"/Mail-Box-$pkgver
}

build() {
  cd  "$srcdir"/Mail-Box-$pkgver
  PERL_MM_USE_DEFAULT=1 perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd  "$srcdir"/Mail-Box-$pkgver
#  ( export PERL_MM_USE_DEFAULT=1 PERL5LIB=""
#    make test
#  )
}

package() {
  cd  "$srcdir"/Mail-Box-$pkgver
  make install DESTDIR="$pkgdir"
}
