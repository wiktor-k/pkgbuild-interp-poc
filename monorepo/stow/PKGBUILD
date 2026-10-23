# Maintainer: Brett Cornwall <ainola@archlinux.org>
# Contributor: Sébastien Luttringer
# Contributor: eric
# Contributor: Tom Newsom
# Contributor: Chris Brannon

pkgname=stow
pkgver=2.4.1
pkgrel=1
pkgdesc='Manage installation of multiple softwares in the same directory tree'
url='https://www.gnu.org/software/stow/'
arch=('any')
license=('GPL-3.0-only')
depends=('perl')
checkdepends=('perl-test-output' 'perl-io-stringy')
validpgpkeys=('942B9075ACCA04E9037C73FED31B5563DAC1D4FA')  # Adam Spiers
source=("https://ftp.gnu.org/gnu/$pkgname/$pkgname-$pkgver.tar.gz"{,.sig})
sha256sums=('2a671e75fc207303bfe86a9a7223169c7669df0a8108ebdf1a7fe8cd2b88780b'
            'SKIP')

build() {
    cd "$pkgname-$pkgver"
    ./configure \
        --prefix=/usr \
        --mandir=/usr/share/man \
        --infodir=/usr/share/info \
        --with-pmdir=/usr/share/perl5/vendor_perl
    make
}

check() {
    cd "$pkgname-$pkgver"
    make check
}

package() {
    cd "$pkgname-$pkgver"
    make DESTDIR="$pkgdir" install
}
