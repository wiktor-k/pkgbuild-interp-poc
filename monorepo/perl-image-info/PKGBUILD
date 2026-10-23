# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: Caleb Cushing <xenoterracide@gmail.com>

pkgname=perl-image-info
pkgver=1.45
pkgrel=1
pkgdesc="Extract meta information from image files"
arch=('any')
license=('PerlArtistic' 'GPL')
options=('!emptydirs')
depends=('perl')
url='https://search.cpan.org/dist/Image-Info'
source=(https://search.cpan.org/CPAN/authors/id/S/SR/SREZIC/Image-Info-$pkgver.tar.gz)
sha256sums=('9d6c1d28c29b137b1a5152e628083d8437577336d8bea191605dda20d55b3539')

build() {
  cd "$srcdir"/Image-Info-$pkgver
  export PERL_MM_USE_DEFAULT=1
  perl Makefile.PL INSTALLDIRS=vendor
  make
  make test
}

package() {
  cd "$srcdir"/Image-Info-$pkgver
  make DESTDIR="${pkgdir}/" install
}
