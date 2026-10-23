# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=perl-app-docmake
_pkgname=App-XML-DocBook-Builder
pkgver=0.1101
pkgrel=1
pkgdesc='Automated builder from docbook/xml to various output formats'
arch=(any)
url='https://github.com/shlomif/docmake'
license=(MIT)
depends=(
  perl
  perl-file-shouldupdate
  perl-class-xsaccessor
  perl-test-trap
)
makedepends=(perl-module-build)
options=(!emptydirs)
source=("$pkgname-$pkgver.tar.gz::https://cpan.metacpan.org/authors/id/S/SH/SHLOMIF/App-XML-DocBook-Builder-${pkgver}.tar.gz")
sha512sums=('99b420c1f4ef4b61e722e91dae9387814e54ea5a577b9da4db420bf4cdd99f5a9d99f76b4b64c14cb86ed349dfa97d67bbbd14a144a0f5e7b9e2ae57e031ecd5')
b2sums=('59e74f6a77615e77c174d3518fde19145d19d8acf0f35167cf7c71ca4651fc300a4d310debe7a5948bee275bc5a4c71c8320d1d296a50f884346f036f0e48b64')

build() {
  cd "$_pkgname-$pkgver"

  unset PERL_MB_OPT PERL5LIB PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1 MODULEBUILDRC=/dev/null

  perl Build.PL --create_packlist=0
  ./Build
}

package() {
  cd "$_pkgname-$pkgver"

  unset PERL5LIB PERL_LOCAL_LIB_ROOT

  ./Build install --installdirs=vendor --destdir="$pkgdir"

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
