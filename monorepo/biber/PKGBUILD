# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Rémy Oudompheng <remy@archlinux.org>

pkgname=biber
epoch=1
pkgver=2.22
pkgrel=1
pkgdesc='A Unicode-capable BibTeX replacement for biblatex users'
arch=(any)
license=(Artistic-2.0)
url='https://github.com/plk/biber'
depends=(perl-autovivification
         perl-business-isbn
         perl-business-ismn
         perl-business-issn
         perl-class-accessor
         perl-data-compare
         perl-data-dump
         perl-data-uniqid
         perl-datetime-calendar-julian
         perl-datetime-format-builder
         perl-file-slurper
         perl-io-string # required by Biber::Output::bblxml
         perl-ipc-run3
         perl-lingua-translit
         perl-list-allutils
         perl-log-log4perl
         perl-lwp-protocol-https
         perl-mozilla-ca
         perl-parse-recdescent
         perl-perlio-utf8-strict
         perl-regexp-common
         perl-sort-key
         perl-text-balanced
         perl-text-bibtex
         perl-text-csv
         perl-text-roman
         perl-unicode-linebreak
         perl-xml-libxml-simple
         perl-xml-libxslt
         perl-xml-writer)
# ExtUtils::LibBuilder and Config::AutoConf are not actual makedepends
makedepends=(git
             perl-module-build
             perl-test-differences)
checkdepends=(perl-capture-tiny
              perl-file-which
              texlive-bibtexextra
              texlive-latexextra)
options=(!emptydirs)
# The sourceforge tarball is missing bblxml.pm
source=(git+https://github.com/plk/biber#tag=v$pkgver)
sha256sums=('07d21b457d986f19c93570816bf44b749eb65fa5af58c4bcfb1c7c177892c982')

build() {
  cd $pkgname
  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  perl Build.PL
  ./Build
}

check() {
  cd $pkgname
  # Run test suite
  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  # The website's server certificate chain was misconfigured, causing connection and download failures.
  rm t/remote-files.t
  LC_CTYPE=en_US.UTF-8 BIBER_DEV_TESTS=1 ./Build test
}

package() {
  cd $pkgname
  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  ./Build install installdirs=vendor destdir="$pkgdir"
}
