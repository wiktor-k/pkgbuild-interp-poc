# Maintainer: Bert Peters <bertptrs@archlinux.org>
# Contributor: Caleb Maclennan <caleb@alerque.com>
# Contributor: Simon Legner <Simon.Legner@gmail.com>

pkgname=svgo
pkgver=4.1.0
pkgrel=1
pkgdesc='Tool for optimizing SVG files'
arch=(any)
url="https://github.com/svg/$pkgname"
license=(MIT)
depends=(nodejs)
makedepends=(npm)
provides=(nodejs-svgo)
replaces=(nodejs-svgo)
conflicts=(nodejs-svgo)
_archive="$pkgname-$pkgver"
source=("https://registry.npmjs.org/$pkgname/-/$_archive.tgz")
noextract=("$_archive.tgz")
sha256sums=('6646ae27350d38dc2f08e708ee64b885b9746775ef7c8091f171d94052a333a6')

package() {
	npm install -g --prefix "$pkgdir/usr" "$_archive.tgz"
	find "$pkgdir/usr" -type d -exec chmod 755 {} +
	chown -R root:root "$pkgdir"
	install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname" "$pkgdir/usr/lib/node_modules/svgo/LICENSE"
}
