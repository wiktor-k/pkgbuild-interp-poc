# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=perl-file-desktopentry
pkgver=0.23
pkgrel=1
pkgdesc='Object to handle .desktop files'
arch=('any')
license=('GPL' 'PerlArtistic')
url="https://codeberg.org/michielb/File-DesktopEntry"
depends=('perl-file-basedir' 'perl-uri')
checkdepends=('perl-test-pod' 'perl-test-pod-coverage')
options=('!emptydirs')
source=(https://search.cpan.org/CPAN/authors/id/M/MI/MICHIELB/File-DesktopEntry-$pkgver.tar.gz)
sha256sums=('63f95798d80b81ea8cea4848f342f26593c60f152dbd04d8981c846f979e6835')

build() {
  cd File-DesktopEntry-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd File-DesktopEntry-$pkgver
  make test
}

package() {
  cd File-DesktopEntry-$pkgver
  make DESTDIR="${pkgdir}" install
}
