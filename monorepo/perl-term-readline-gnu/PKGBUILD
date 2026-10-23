# CPAN Name  : Term-ReadLine-Gnu
# Maintainer: Daniel Wallace <danielwallace at gtmanfred dot com>
# Contributor: jason ryan <jasonwryan@gmail.com>
# Contributor: AUR Perl <aurperl@juster.info>
# Generator  : CPANPLUS::Dist::Arch 1.15

pkgname=perl-term-readline-gnu
pkgver=1.47
pkgrel=3
pkgdesc="GNU Readline XS library wrapper"
arch=('x86_64')
license=('PerlArtistic' 'GPL')
options=('!emptydirs')
depends=('perl>=5.7')
url='https://search.cpan.org/dist/Term-ReadLine-Gnu'
source=("https://search.cpan.org/CPAN/authors/id/H/HA/HAYASHI/Term-ReadLine-Gnu-${pkgver}.tar.gz")
sha512sums=('f61a3b57e719ff74504cc1b4540111ee677d6d47779f8cba944b3a950b2b5ab0c7a5ffdb7af52b31d881f99f68c9ab36fe5a400cce62396b1a72343713a626d6')

build() {
  cd "${srcdir}/Term-ReadLine-Gnu-$pkgver"

  export PERL_MM_USE_DEFAULT=1 PERL5LIB="" \
    PERL_AUTOINSTALL=--skipdeps \
    PERL_MM_OPT="INSTALLDIRS=vendor" \
    PERL_MB_OPT="--installdirs vendor" \
    MODULEBUILDRC=/dev/null

  /usr/bin/perl Makefile.PL
  make
}

check() {
  cd "${srcdir}/Term-ReadLine-Gnu-$pkgver"
  PERL_MM_USE_DEFAULT=1 PERL5LIB="" TERM=xterm make test
}

package() {
  cd "${srcdir}/Term-ReadLine-Gnu-$pkgver"
  make pure_install PERL_INSTALL_ROOT="$pkgdir"
}
