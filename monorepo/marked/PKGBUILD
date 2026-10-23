# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=marked
pkgver=17.0.0
pkgrel=1
pkgdesc="A markdown parser built for speed"
arch=('any')
url='https://marked.js.org'
license=('MIT')
depends=('nodejs')
makedepends=('npm')
source=(https://registry.npmjs.org/$pkgname/-/$pkgname-$pkgver.tgz)
noextract=($pkgname-$pkgver.tgz)
sha512sums=('2a40d8116104898270fca0be0d59b5cf396931048c22ee98465b64702bf085e0a9f0758f5c293d2703a61c92819467f3729cdc22deb1dec3274ec466ff9d680e')

package() {
  npm install -g --prefix "$pkgdir"/usr "$srcdir"/$pkgname-$pkgver.tgz

  # Non-deterministic race in npm gives 777 permissions to random directories.
  # See https://github.com/npm/npm/issues/9359 for details.
  chmod -R u=rwX,go=rX "$pkgdir"

  # npm installs package.json owned by build user
  # https://bugs.archlinux.org/task/63396
  chown -R root:root "$pkgdir"

  install -d "$pkgdir"/usr/share/licenses/$pkgname
  ln -s ../../../lib/node_modules/marked/LICENSE.md "$pkgdir"/usr/share/licenses/$pkgname/LICENSE.md
}
