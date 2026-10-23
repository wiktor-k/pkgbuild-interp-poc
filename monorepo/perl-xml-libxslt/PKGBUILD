# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Jonathan Steel <jsteel@aur.archlinux.org>
# Contributor: Justin "juster" Davis <jrcd83@gmail.com>
# Contributor: François Charette <firmicus ατ gmx δοτ net>

pkgname=perl-xml-libxslt
pkgver=2.003000
pkgrel=4
pkgdesc="Interface to the gnome libxslt library"
arch=('x86_64')
url="https://search.cpan.org/dist/XML-LibXSLT"
license=('GPL')
depends=('perl-xml-libxml' 'libxslt')
checkdepends=('perl-test-pod')
options=('!emptydirs')
source=(https://www.cpan.org/authors/id/S/SH/SHLOMIF/XML-LibXSLT-$pkgver.tar.gz)
sha512sums=('726af80f3f791afe8c0db1bdb75524c60d6560178759f9f37f8ad99da3f20c4712257e24fdc24635a40ac132e48cb3a86191a09e1245dc778806e85aecbd00dd')

build() {
  cd XML-LibXSLT-$pkgver
  perl Makefile.PL INSTALLDIRS=vendor
  make
}

check() {
  cd XML-LibXSLT-$pkgver
  make test
}

package() {
  cd XML-LibXSLT-$pkgver
  make DESTDIR="$pkgdir" install
}
