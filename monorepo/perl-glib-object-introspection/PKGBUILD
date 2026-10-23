# Maintainer: Muflone http://www.muflone.com/contacts/english/
# Contributor: Brian Bidulock <bidulock@openss7.org>
# Contributor: John D Jones III AKA jnbek <jnbek1972 -_AT_- g m a i l -_Dot_- com>

pkgname=perl-glib-object-introspection
_perl_namespace=Glib
_perl_module=Object-Introspection
pkgver=0.052
pkgrel=2
pkgdesc="Dynamically create Perl language bindings"
arch=('x86_64')
url="https://metacpan.org/release/${_perl_namespace}-${_perl_module}"
license=('LGPL-2.1-or-later')
makedepends=('gobject-introspection' 'perl-extutils-depends' 'perl-extutils-pkgconfig' 'perl-cairo-gobject')
depends=('glib-perl' 'gobject-introspection-runtime')
source=("https://cpan.metacpan.org/authors/id/X/XA/XAOC/${_perl_namespace}-${_perl_module}-${pkgver}.tar.gz"
         https://gitlab.gnome.org/GNOME/perl-glib-object-introspection/-/commit/22b1d2c8f3cdf9351f37791b27b36c139a145383.patch)
sha512sums=('50cdb8e628f09440656e337974c0dd87d7b90cbadadc5ba29fa34df31d42916f691134e081296027035dc98a2ce6b8fa241aace898a9613cf3c94b91c3eec500'
            'cea5c0063e9c108b022214373d8062b23cd288424147bd786ff5509de01329b70e9fb8f5cd3737d6455087d9091bdc08dbf434c16717cd2b08dd5d18e4e5a95b')
options=('!emptydirs')

prepare() {
  cd "${_perl_namespace}-${_perl_module}-${pkgver}"
  patch -p1 -i ../22b1d2c8f3cdf9351f37791b27b36c139a145383.patch # Support glib2 2.90
}

build() {
  cd "${_perl_namespace}-${_perl_module}-${pkgver}"
  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1 PERL_AUTOINSTALL=--skipdeps
  perl Makefile.PL
  make
}

check() {
  cd "${_perl_namespace}-${_perl_module}-${pkgver}"
  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  export PERL_MM_USE_DEFAULT=1
  make test
}

package() {
  cd "${_perl_namespace}-${_perl_module}-${pkgver}"
  unset PERL5LIB PERL_MM_OPT PERL_MB_OPT PERL_LOCAL_LIB_ROOT
  make pure_install INSTALLDIRS=vendor DESTDIR="${pkgdir}"
  # Delete unuseful files
}
