# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: T.J. Townsend <blakkheim@archlinux.org>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: Geoffroy Carrier <geoffroy.carrier@koon.fr>
# Contributor: Baptiste Daroussin <baptiste.daroussin@gmail.com>

pkgname=tnftp
pkgver=20260211
pkgrel=1
pkgdesc='NetBSD FTP client with several advanced features'
url='https://ftp.netbsd.org/pub/NetBSD/misc/tnftp/'
arch=('x86_64')
license=('BSD')
depends=('glibc' 'ncurses' 'libncursesw.so' 'openssl')
source=(https://ftp.netbsd.org/pub/NetBSD/misc/tnftp/$pkgname-$pkgver.tar.gz{,.asc})
sha256sums=('101cda6927e5de4338ad9d4b264304d7d15d6a78b435968a7b95093e0a2efe03'
            'SKIP')
b2sums=('3d849d01073d6b93bd60be76e3509c54fd867723c10c8352604691047f66a97e10cde1b81e5c2b98a427ed1cb689a1f484cc9795904e3f3dcf6314e503941db0'
        'SKIP')
validpgpkeys=(
  '2A8E22EDB07B5414548D8507A4186D9A7F332472'  # Luke Mewburn <lukem@NetBSD.org>
)

build() {
  cd "${pkgname}-${pkgver}"
  ./configure \
    --prefix=/usr \
    --with-socks=no \
    --enable-ssl
  make
}

package() {
  cd "${pkgname}-${pkgver}"
  make install DESTDIR="${pkgdir}"
  install -Dm 755 COPYING -t "${pkgdir}/usr/share/licenses/${pkgname}"
  install -Dm 644 ChangeLog INSTALL NEWS README THANKS todo -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
