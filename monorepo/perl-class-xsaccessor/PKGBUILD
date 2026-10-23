# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=perl-class-xsaccessor
_pkgname=Class-XSAccessor
pkgver=1.19
pkgrel=2
pkgdesc='Fast access generation module for Perl'
arch=(x86_64)
url='https://github.com/tsee/Class-XSAccessor'
license=('GPL-1.0-or-later OR Artistic-1.0-Perl')
depends=(glibc perl)
makedepends=(git)
options=(!emptydirs)
source=("$pkgname::git+$url#tag=${_pkgname}-${pkgver}")
sha512sums=('80e1e68d1fbf89c0c5cae2e86e13b18443b0005834277b94135a60c6d8bec3e3e41c6e3730205d2c83dd8e05b620860054bf182a8f713a8335359eba71c5d6e4')
b2sums=('624d980931232f61080aeccf5ecbc077044ba0157b80c21701366c9d06148c9e6ddbaa43dc28f9b10c343b16d5033f094adefd3e4b07bb74ce32bd1338c27a39')

build() {
  cd "$pkgname"

  unset PERL_MM_OPT PERL5LIB PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1 PERL_AUTOINSTALL=--skipdeps

  perl Makefile.PL NO_PACKLIST=1 NO_PERLLOCAL=1
  make
}

package() {
  cd "$pkgname"
  unset PERL5LIB PERL_LOCAL_LIB_ROOT

  make install INSTALLDIRS=vendor DESTDIR="$pkgdir"
}
