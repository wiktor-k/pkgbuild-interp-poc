# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Florian Pritz <flo@xinu.at>
# Contributor: Jan Fader <jan.fader@web.de>

pkgbase=highlight
pkgname=('highlight' 'highlight-perl' 'highlight-gui')
pkgver=4.21
pkgrel=2
pkgdesc="Fast and flexible source code highlighter"
url="http://www.andre-simon.de/doku/highlight/highlight.html"
license=('GPL-3.0-or-later')
arch=('x86_64')
makedepends=('qt6-base' 'lua' 'boost' 'swig')
source=("http://www.andre-simon.de/zip/${pkgbase}-${pkgver}.tar.bz2"{,.asc}
        'use_gcc.patch')
sha256sums=('310785a78bd054c4592ff012924c0542826ab53b7afd68d207462321c409a999'
            'SKIP'
            '5aaacca96f1f4307bbe9cfb5b0f67a98edb0b5d653270990176e20840196ab50')
validpgpkeys=('B8C55574187F49180EDC763750FE0279D805A7C7')

prepare() {
	cd "${pkgbase}-${pkgver}"

	# Use gcc instead of clang
	patch -Np1 -i "${srcdir}/use_gcc.patch"
}

build() {
	cd "${pkgbase}-${pkgver}"

	make QMAKE=qmake6
	make QMAKE=qmake6 gui
	make -C extras/swig perl
}

package_highlight() {
	pkgdesc="${pkgdesc} - CLI version"
	depends=('lua')
	backup=("etc/${pkgbase}/filetypes.conf")

	cd "${pkgbase}-${pkgver}"

	make DESTDIR="${pkgdir}" QMAKE=qmake6 install

	# Remove extras modules files from doc folder
	rm -rf "${pkgdir}/usr/share/doc/${pkgbase}/extras"
}

package_highlight-perl() {
	pkgdesc="${pkgdesc} - perl module"
	depends=('highlight' 'perl')
	_perl_path=$(perl -V:vendorarch | awk -F"'" '{print $2}')

	cd "${pkgbase}-${pkgver}"

	install -Dm 644 "extras/swig/${pkgbase}.pm" "${pkgdir}${_perl_path}/${pkgbase}.pm"
	install -Dm 755 "extras/swig/${pkgbase}.so" "${pkgdir}${_perl_path}/auto/${pkgbase}/${pkgbase}.so"
}
package_highlight-gui() {
	pkgdesc="${pkgdesc} - QT GUI version"
	depends=('qt6-base' 'highlight' 'hicolor-icon-theme')

	cd "${pkgbase}-${pkgver}"

	install -dm 755 "${pkgdir}/usr/bin"
	make DESTDIR="${pkgdir}" QMAKE=qmake6 install-gui
}
