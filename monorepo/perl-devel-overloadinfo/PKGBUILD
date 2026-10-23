# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=perl-devel-overloadinfo
pkgver=0.008
pkgrel=1
pkgdesc="Introspect overloaded operators"
arch=('any')
url="https://metacpan.org/dist/Devel-OverloadInfo"
license=('GPL-1.0-or-later' 'Artistic-1.0-Perl')
depends=('perl' 'perl-mro-compat' 'perl-package-stash')
checkdepends=('perl-test-fatal')
options=('!emptydirs')
source=("https://cpan.metacpan.org/authors/id/I/IL/ILMARI/Devel-OverloadInfo-$pkgver.tar.gz")
sha512sums=('1c614e1132c8c763ccac79ea3c80c08636d617af55bdcba4cc52aa06c9cefca07fd4f1921e54f5dc76b140ed8546da917602eab81f59efd6c490f23b8b8abe41')

build() {
  cd Devel-OverloadInfo-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd Devel-OverloadInfo-$pkgver
  make test
}

package() {
  cd Devel-OverloadInfo-$pkgver
  make DESTDIR="$pkgdir" install
}
