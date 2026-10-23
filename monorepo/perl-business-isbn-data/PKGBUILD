# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Rémy Oudompheng <remy@archlinux.org>
# Contributor: Anton Leontiev <bunder /at/ t-25.ru>

pkgname=perl-business-isbn-data
pkgver=20261016.001
pkgrel=1
pkgdesc='Data pack for Business::ISBN'
arch=(any)
url='https://github.com/briandfoy/business-isbn-data'
license=(Artistic-2.0)
depends=(perl perl-business-isbn-data)
makedepends=(git)
checkdepends=(perl-test-pod)
optdepends=('perl-libwww: for access to the xISBN service')
options=(purge !emptydirs)
source=("$pkgname::git+$url#tag=release-${pkgver}")
sha512sums=('d034a168f53d74514d6132e5f9a21535a75b60c8a0ee74cf504b7a9ae9d0fcfaf0ed38ba6bee7bec1dcd16664681411f2199478cd3ff378d3e3b54fc44b716d1')
b2sums=('44e6b1518f1bebbadf42384f5c1e991c832606972a1aa56743f5ff7519d6714a2b556092f2a99f21d50eebb24b51c46ddd61d8b0503f4bc00439ec7d59e87b9c')

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
