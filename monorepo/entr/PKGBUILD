# Maintainer: Filipe Laíns (FFY00) <lains@archlinux.org>
# Maintainer: Robin Candau <antiz@archlinux.org>

pkgname=entr
pkgver=5.9
pkgrel=1
pkgdesc="Run arbitrary commands when files change"
url="http://eradman.com/entrproject"
arch=('x86_64')
license=('MIT')
depends=('glibc')
checkdepends=('procps-ng' 'git' 'vim' 'tmux')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/eradman/entr/archive/refs/tags/${pkgver}.tar.gz")
sha512sums=('60b46312c3aa3bf9193ad2f594a3d6e97d8fb3a90e22921cefb30bfcf6f37e3d4230d74d22aeab2aa6b9cc106fd4eed4f91c58fdbbbc3410915c6a8c6c685172')

build() {
	cd "${pkgname}-${pkgver}"
	export PREFIX='/usr'
	./configure
	make
}

check() {
	cd "${pkgname}-${pkgver}"
	TERM='xterm' make test
}

package() {
	cd "${pkgname}-${pkgver}"
	make DESTDIR="${pkgdir}" install
	install -Dm 644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
