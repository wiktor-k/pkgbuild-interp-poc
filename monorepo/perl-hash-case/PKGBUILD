# Maintainer: Sergej Pupykin <arch+pub@sergej.pp.ru>

pkgname='perl-hash-case'
pkgver=1.07
pkgrel=1
pkgdesc="base class for hashes with key-casing requirements"
arch=('any')
license=('PerlArtistic' 'GPL')
options=('!emptydirs')
depends=('perl')
makedepends=()
url='https://search.mcpan.org/dist/Hash-Case'
source=("https://cpan.metacpan.org/authors/id/M/MA/MARKOV/Hash-Case-$pkgver.tar.gz")
sha256sums=('f591db9f9a8355c67fba94ae27e06e6339b800ca78c5250d75c7688c0bc33969')

build() {
  ( export PERL_MM_USE_DEFAULT=1 PERL5LIB=""                 \
      PERL_AUTOINSTALL=--skipdeps                            \
      PERL_MM_OPT="INSTALLDIRS=vendor DESTDIR='$pkgdir'"     \
      PERL_MB_OPT="--installdirs vendor --destdir '$pkgdir'" \
      MODULEBUILDRC=/dev/null

    cd "$srcdir/Hash-Case-$pkgver"
    /usr/bin/perl Makefile.PL
    make
  )
}

check() {
  cd "$srcdir/Hash-Case-$pkgver"

  ( export PERL_MM_USE_DEFAULT=1 PERL5LIB=""
    make test
  )
}

package() {
  cd "$srcdir/Hash-Case-$pkgver"
  make install
  find "$pkgdir" -name .packlist -o -name perllocal.pod -delete
}
