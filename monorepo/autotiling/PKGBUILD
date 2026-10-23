# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Piotr Miller <nwg.piotr@gmail.com>

pkgname=autotiling
pkgver=1.9.4
pkgrel=1
pkgdesc="Script for sway and i3 to automatically switch the horizontal / vertical window split orientation"
url="https://github.com/nwg-piotr/autotiling"
arch=('any')
license=('GPL-3.0-or-later')
depends=('python' 'python-i3ipc')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/v${pkgver}.tar.gz")
sha256sums=('fc3ffea5bf074482236fc39b38943ef123fd27f5c2e7136a628a0608e6917109')

build() {
        cd "${pkgname}-${pkgver}"
        python -m build --wheel --no-isolation
}

package() {
	cd "${pkgname}-${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl
	
	install -Dm 644 README.md "${pkgdir}/usr/share/doc/${pkgname}/README.md"
}
