# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Florian Pritz <bluewind@xinu.at>

pkgname=perl-cgi-fast
pkgver=2.17
pkgrel=3
pkgdesc='CGI Interface for Fast CGI'
arch=(any)
url='https://github.com/leejo/cgi-fast'
license=('Artistic-1.0-Perl OR GPL-1.0-or-later')
depends=(perl perl-cgi perl-fcgi)
makedepends=(git)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('dcb5bc412645cf44ff2d9e9882889bd75b0f50ce79af6fe8eb656360f40ca21d7e15680e8a76b3a20fe4077be64310c53300d7936796ffa1c10d41f81134e262')
b2sums=('52c361eef4c730f787aacdafaebab71a7e4bca387c4b041f91d582a108609e5038fb38b98e4078228233d59a6b3cccf01cffd8efe8022f3f09bfdde7f37d40da')
options=(!emptydirs)

build() {
  cd "$pkgname"

  export PERL_MM_USE_DEFAULT=1 PERL_AUTOINSTALL=--skipdeps
  unset PERL5LIB PERL_MM_OPT
  /usr/bin/perl Makefile.PL
  make
}

check() {
  cd "$pkgname"

  export PERL_MM_USE_DEFAULT=1
  unset PERL5LIB
  make test
}

package() {
  cd "$pkgname"

  make install INSTALLDIRS=vendor DESTDIR="$pkgdir"
}
# vim:set ts=2 sw=2 et:
