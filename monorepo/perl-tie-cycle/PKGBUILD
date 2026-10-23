# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Rémy Oudompheng <remy@archlinux.org>

pkgname=perl-tie-cycle
pkgver=1.233
pkgrel=1
pkgdesc='Cycle through a list of values via a scalar'
arch=(any)
url='https://github.com/briandfoy/tie-cycle'
license=(Artistic-2.0)
depends=(perl)
makedepends=(git)
checkdepends=(perl-test-pod perl-test-pod-coverage)
source=("$pkgname::git+$url#tag=release-$pkgver")
sha512sums=('f7f4f7fea0c06c16e907802a42425379b7ccf91790408c93b66a4d99ca34debcfa2321a7e2298f904541ef1283583421606017f1358ce07465be9bf7531ef366')
b2sums=('7055ef9be9c087e6fd53715ae8fef3582b5f9703d63583d35a9ea6c787313808d7463d259269b33f27e579bd576b02233e32aa71541224bd9ac0e403b54211c3')
options=(!emptydirs)

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
