# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Tom Hacohen <tom@stosb.com>
pkgname=acme-tiny
pkgver=5.0.3
pkgrel=2
pkgdesc="A tiny script to issue and renew TLS certs from Let's Encrypt"
url="https://github.com/diafygi/acme-tiny"
depends=('python' 'openssl')
license=('MIT')
arch=('any')
source=($pkgname-$pkgver.tar.gz::https://github.com/diafygi/acme-tiny/archive/${pkgver}.tar.gz)
sha256sums=('c036e6e62037f5736c5c6810b6e3fd951f33dc4cda41cc169d7800d06e6ba3d0')

package() {
    cd "$srcdir"/acme-tiny-${pkgver}

    install -Dm 755 acme_tiny.py "$pkgdir/usr/bin/acme-tiny"
    install -Dm 644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}
