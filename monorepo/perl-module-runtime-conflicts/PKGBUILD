# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=perl-module-runtime-conflicts
pkgver=0.003
pkgrel=1
pkgdesc="Provide information on conflicts for Module::Runtime"
arch=('any')
url="https://metacpan.org/dist/Module-Runtime-Conflicts"
license=('GPL-1.0-or-later' 'Artistic-1.0-Perl')
depends=('perl' 'perl-dist-checkconflicts' 'perl-module-runtime')
options=('!emptydirs')
source=("https://cpan.metacpan.org/authors/id/E/ET/ETHER/Module-Runtime-Conflicts-$pkgver.tar.gz")
sha512sums=('1b72ea7cc9516721290d5ea433169697b4134b5fc77872d6822e4c395a154bf956d469ab5da2802d587d60dfb13b9784014fe492988cf3f8b5820da2d63e8ca1')

build() {
  cd Module-Runtime-Conflicts-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd Module-Runtime-Conflicts-$pkgver
  make test
}

package() {
  cd Module-Runtime-Conflicts-$pkgver
  make DESTDIR="$pkgdir" install
}
