# Maintainer: Bruno Pagani <archange@archlinux.org>
# Contributor: Maarten de Vries <maarten@de-vri.es>

_pkg=filter-dkimsign
pkgname=opensmtpd-$_pkg
pkgver=0.6
pkgrel=5
pkgdesc="OpenSMTPD filter for signing mail with DKIM"
license=(ISC)
url="https://src.imperialat.at/?action=summary&path=filter-dkimsign"
arch=(x86_64)
depends=(glibc libopensmtpd openssl)
makedepends=(git gzip mandoc)
source=(git+ssh://anonymous@src.imperialat.at/$_pkg#tag=v$pkgver)
sha512sums=('b92ceec6abc756a39fb3a8da4643d7122ad264f108e8bbdb6cf26624d485e9b8fffbeed33a76d0163f9d9ec37cd141bdaf89152375446cbc5f59ecdeb53aaf18')

build() {
  cd $_pkg
  make -f Makefile.gnu
}

package() {
  cd $_pkg
  make -f Makefile.gnu BINDIR=/usr/lib/smtpd/opensmtpd/ DESTDIR="$pkgdir" install
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
