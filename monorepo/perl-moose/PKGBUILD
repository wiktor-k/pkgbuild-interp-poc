# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=perl-moose
pkgver=2.4001
pkgrel=2
pkgdesc="A postmodern object system for Perl 5"
arch=('x86_64')
url="https://metacpan.org/dist/Moose"
license=('GPL-1.0-or-later' 'Artistic-1.0-Perl')
depends=(
  'perl'
  'perl-class-load'
  'perl-class-load-xs'
  'perl-data-optlist'
  'perl-devel-globaldestruction'
  'perl-devel-overloadinfo'
  'perl-devel-stacktrace'
  'perl-dist-checkconflicts'
  'perl-eval-closure'
  'perl-module-runtime'
  'perl-module-runtime-conflicts'
  'perl-mro-compat'
  'perl-package-deprecationmanager'
  'perl-package-stash'
  'perl-package-stash-xs'
  'perl-params-util'
  'perl-sub-exporter'
  'perl-try-tiny'
)
checkdepends=(
  'perl-cpan-meta-check'
  'perl-moo'
  'perl-specio'
  'perl-super'
  'perl-test-deep'
  'perl-test-fatal'
  'perl-test-leaktrace'
  'perl-test-memory-cycle'
  'perl-test-needs'
  'perl-test-output'
  'perl-test-warnings'
  'perl-type-tiny'
)
options=('!emptydirs')
source=("https://cpan.metacpan.org/authors/id/E/ET/ETHER/Moose-$pkgver.tar.gz")
sha512sums=('22a1f48ecbaa8fbe9033bd7297b60417bcdb2c12a2e4118ebaaf009b0a4a75e6e0cdd681caa31c1d0a87f2b72408971bf7927b5fbab8c61cb813c1b17727ee93')

build() {
  cd Moose-$pkgver
  PERL_MM_USE_DEFAULT=1 perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd Moose-$pkgver
  make test
}

package() {
  cd Moose-$pkgver
  make DESTDIR="$pkgdir" install
}
