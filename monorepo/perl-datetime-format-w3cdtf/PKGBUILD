# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: François Charette <firmicus ατ gmx δοτ net>

pkgname=perl-datetime-format-w3cdtf
pkgver=0.09
pkgrel=1
pkgdesc="Parse and format W3CDTF datetime strings"
arch=('any')
url="https://search.cpan.org/dist/DateTime-Format-W3CDTF"
license=('GPL' 'PerlArtistic')
depends=('perl>=5.10.0' 'perl-datetime')
options=('!emptydirs')
source=(https://search.cpan.org/CPAN/authors/id/G/GW/GWILLIAMS/DateTime-Format-W3CDTF-$pkgver.tar.gz)
sha256sums=('a63a913413784fdc15d39c7621ce86307ab5e7b3c1cf46a26fe5efc5c87972f7')

build() {
  cd  "$srcdir"/DateTime-Format-W3CDTF-$pkgver
  PERL_MM_USE_DEFAULT=1 perl Makefile.PL INSTALLDIRS=vendor
  make
}

package() {
  cd  "$srcdir"/DateTime-Format-W3CDTF-$pkgver
  make install DESTDIR="$pkgdir"
}
