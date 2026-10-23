# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=perl-file-mimeinfo
pkgver=0.37
pkgrel=1
pkgdesc='Determine file type, includes mimeopen and mimetype'
arch=('any')
url="https://metacpan.org/release/File-MimeInfo"
license=('GPL-1.0-or-later OR Artistic-1.0-Perl')
depends=('perl-encode-locale' 'perl-file-basedir' 'perl-file-desktopentry' 'shared-mime-info')
checkdepends=('perl-test-pod' 'perl-test-pod-coverage')
options=('!emptydirs')
source=(https://cpan.metacpan.org/authors/id/M/MI/MICHIELB/File-MimeInfo-$pkgver.tar.gz)
sha512sums=('c7a5a06fc2561e6724f406142f31a4f440f813b8c8d02d7bbf26daa74a3fb9816d2001329785cd2c9b1469292b0bbeaf130a2d809a0d26514161ab11257219b4')

build() {
  cd File-MimeInfo-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd File-MimeInfo-$pkgver
  make test
}

package() {
  cd File-MimeInfo-$pkgver
  make DESTDIR="${pkgdir}" install
}
