# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# This PKGBUILD was autocreated.  See https://xtermin.us/archperl/ for details
# Maintainer: Charles Mauch <cmauch@gmail.com>

pkgname=perl-user-identity
pkgver=4.00
pkgrel=2
pkgdesc="Perl/CPAN Module User::Identity : Define a person"
arch=("any")
url="https://search.cpan.org/~markov/User-Identity/"
license=("GPL" "PerlArtistic")
source=("https://www.cpan.org/authors/id/M/MA/MARKOV/User-Identity-$pkgver.tar.gz")
sha256sums=('46ec55c4b2c158fb9e3bd5c63aaa10695fee8508ef4ec958774dd8eaccab3847')
options=(!emptydirs)

build() {
  cd "$srcdir"/User-Identity-$pkgver
  PERL_MM_USE_DEFAULT=1 perl Makefile.PL INSTALLDIRS=vendor
  make
}

package() {
  cd "$srcdir"/User-Identity-$pkgver
  make install DESTDIR="$pkgdir"
  find "$pkgdir" -name '.packlist' -delete
  find "$pkgdir" -name '*.pod' -delete
}
