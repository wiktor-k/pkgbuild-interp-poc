# Maintainer: Jonathan Steel <jsteel at aur.archlinux.org>
# Contributor: Justin Davis <jrcd83@gmail.com>

pkgname=perl-package-stash-xs
pkgver=0.32
pkgrel=1
pkgdesc="Faster and more correct implementation of the Package::Stash API"
arch=('x86_64')
url="https://metacpan.org/release/Package-Stash-XS"
license=('Artistic-1.0-Perl' 'GPL-1.0-or-later')
depends=('perl')
checkdepends=('perl-package-stash' 'perl-test-fatal' 'perl-test-needs')
options=('!emptydirs')
source=(https://cpan.metacpan.org/authors/id/E/ET/ETHER/Package-Stash-XS-$pkgver.tar.gz)
sha512sums=('69758245572f71c6a9977705aed3be6cd264986fbdc38a5e964a1d0b7f8a8b6511f8c5ee89a24521c7ce8d4af5626697cfe69c074f52f1d786926e36159ac386')

build() {
  cd Package-Stash-XS-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd Package-Stash-XS-$pkgver
  make test
}

package() {
  cd Package-Stash-XS-$pkgver
  make DESTDIR="$pkgdir" install
}
