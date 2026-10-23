# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Maintainer: Charles Mauch <cmauch@gmail.com>

pkgname=perl-object-realize-later
pkgver=4.00
pkgrel=1
pkgdesc="Perl/CPAN Module Object::Realize::Later : Delay construction of real data until used"
arch=("any")
url="https://search.cpan.org/~markov/Object-Realize-Later/"
license=("GPL" "PerlArtistic")
source=("https://www.cpan.org/authors/id/M/MA/MARKOV/Object-Realize-Later-$pkgver.tar.gz")
sha256sums=('c4753d5a35f147eede09cdbd5e6d627dde3bdaaabfe9e56f2cff72b72d19979b')
options=(!emptydirs)

build() {
  cd "$srcdir"/Object-Realize-Later-$pkgver
  PERL_MM_USE_DEFAULT=1 perl Makefile.PL INSTALLDIRS=vendor
  make
}

package() {
  cd "$srcdir"/Object-Realize-Later-$pkgver
  make install DESTDIR="$pkgdir"
}
