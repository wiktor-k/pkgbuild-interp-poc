# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Morten Linderud <foxboron@archlinux.org>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=python-nbxmpp
pkgdesc="XMPP Python library"
pkgver=7.4.0
pkgrel=6
url="https://gitlab.com/gajim/python-nbxmpp"
arch=('any')
license=('GPL-3.0-only')
depends=('libsoup3' 'python' 'python-gobject' 'python-idna' 'python-precis_i18n' 'python-packaging')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
source=("${url}/-/archive/${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha256sums=('ee9329dc4f3fd6a8c26ad1f677b728d67def70e41060965f8e4340ba85114dc0')
b2sums=('bbab29003942ae683df709c65c2c1a03663bc965357ab7d1c500eaf56ab4aa6908c26fba2c23479dd24280d4ffee6a626064ad454d3908bfdbf994bb72bfae3a')

build() {
	cd "${pkgname}-${pkgver}"
	python -m build --wheel --no-isolation
}

package() {
	cd "${pkgname}-${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl
}
