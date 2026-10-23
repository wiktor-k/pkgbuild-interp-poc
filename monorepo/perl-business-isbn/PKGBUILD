# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Rémy Oudompheng <remy@archlinux.org>

pkgname=perl-business-isbn
pkgver=3.014
pkgrel=1
pkgdesc='Perl package to deal with International Standard Book Numbers'
arch=(any)
url='https://github.com/briandfoy/business-isbn'
license=(Artistic-2.0)
depends=(perl perl-business-isbn-data)
makedepends=(git)
checkdepends=(perl-test-pod)
optdepends=('perl-libwww: for access to the xISBN service')
options=(!emptydirs)
source=("$pkgname::git+$url#tag=release-${pkgver}")
sha512sums=('47f6d917c6508b0f43ec90bbe9dde269724537218b26b9c905d98eb04e08d9ada4d7469a76df4a1f1d049d418ffef4049d4112392f54cbccd12dfdf411b5500b')
b2sums=('437bfab2239dcd40ededda5d4d383e94771fc832a139bd1538356cc14f5f88dfbdd4806619e31ca3f88bbe5ac3c7e4baf47e4385fe3c67e64dd13d4ee49951e9')

build() {
  cd "$pkgname"

  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1 PERL_AUTOINSTALL=--skipdeps
  perl Makefile.PL
  make
}

check() {
  cd "$pkgname"

  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1
  make test
}

package() {
  cd "$pkgname"

  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  make install INSTALLDIRS=vendor DESTDIR="$pkgdir"
}
