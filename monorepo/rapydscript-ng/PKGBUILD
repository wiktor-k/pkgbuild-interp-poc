# Maintainer: Eli Schwartz <eschwartz@archlinux.org>

pkgname=rapydscript-ng
pkgver=0.8.6
pkgrel=1
pkgdesc="Pythonic JavaScript that doesn't suck"
arch=('any')
url="https://github.com/kovidgoyal/${pkgname}"
license=('BSD')
depends=('nodejs')
makedepends=('git' 'npm')
source=("git+${url}.git?signed#tag=v${pkgver}")
sha256sums=('8131da33bf3bd344d26693cc870c99bb7f88befdb8f9a9030994347bea73c4db')
validpgpkeys=('3CE1780F78DD88DF45194FD706BC317B515ACE7C') # Kovid Goyal (New longer key) <kovid@kovidgoyal.net>

build() {
    cd ${pkgname}

    npm install --no-optional
    bin/rapydscript self --complete
    rm -rf release/ && mv dev/ release/
}

check() {
    cd ${pkgname}

    bin/rapydscript test
}

package() {
    cd ${pkgname}

    mkdir -p "${pkgdir}"/usr/{bin,lib/node_modules/rapydscript-ng/}
    cp -r * "${pkgdir}"/usr/lib/node_modules/rapydscript-ng/
    # Non-deterministic race in npm gives 777 permissions to random directories.
    # See https://github.com/npm/npm/issues/9359 for details.
    find "${pkgdir}"/usr -type d -exec chmod 755 {} +

    ln -s ../lib/node_modules/rapydscript-ng/bin/rapydscript "${pkgdir}"/usr/bin/rapydscript

    install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
