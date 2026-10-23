# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: John D Jones III <j[nospace]n[nospace]b[nospace]e[nospace]k[nospace]1972 -_AT_- the domain name google offers a mail service at ending in dot com>
# Generator  : CPANPLUS::Dist::Arch 1.28

pkgname=perl-devel-checkcompiler
pkgver=0.07
pkgrel=14
pkgdesc="Check the compiler's availability"
arch=('any')
license=('PerlArtistic' 'GPL')
options=('!emptydirs')
depends=('perl')
makedepends=('perl-module-build-tiny')
checkdepends=('perl-test-requires')
url='https://search.cpan.org/dist/Devel-CheckCompiler'
source=("https://search.cpan.org/CPAN/authors/id/S/SY/SYOHEX/Devel-CheckCompiler-$pkgver.tar.gz")
sha512sums=('147fbdab382ff40bf1402d2d6be1143fe77403cd2444e854fdf11eb02e8acdcab8cb6bbdd7f3aef01cc64fac2862efaa044c5148c7770b84b464f7360c399d7a')
_distdir="Devel-CheckCompiler-$pkgver"

build() {
  cd "$srcdir/$_distdir"
  perl Build.PL
  perl Build
}

check() {
  cd "$srcdir/$_distdir"
  perl Build test
}

package() {
  cd "$srcdir/$_distdir"
  perl Build install --installdirs=vendor --destdir="$pkgdir"
}

# vim:set ts=2 sw=2 et:
