# Maintainer: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Miguel Revilla <yo@miguelrevilla.com>
# Contributor: Massimiliano Torromeo <massimiliano.torromeo@gmail.com>
# Contributor: goodmen <goodmenzy@gmail.com>

pkgname=openipmi
_pkgname=OpenIPMI
pkgver=2.0.37
pkgrel=4
pkgdesc='Full-function IPMI (Intelligent Platform Management Interface) system'
url='http://openipmi.sourceforge.net/'
arch=('x86_64')
license=('LGPL2.1')
makedepends=('python' 'swig')
depends=('popt' 'ncurses' 'net-snmp' 'glib2' 'gdbm' 'libedit')
optdepends=('python: bindings')
source=("https://downloads.sourceforge.net/project/${pkgname}/${_pkgname}%202.0%20Library/${_pkgname}-${pkgver}.tar.gz"
         swig-4.5.patch)
sha512sums=('ef959ece52f073c108a84911f0aeff421c9408f70b7dacf286ec5fdb023e281d3e97b425ac2c3c70bdfebfafda61ba923c8c052da52448d26d5fcf5e4d1c764a'
            'a2df73f76a26edbe2c5ef6a414154de3dd74b0623e84190c23b4cfae6ef0e104af47491774b4276e22cf673a2d225dfda4834eefa582a97b92bff9b68b99c9df')

prepare() {
	cd "${srcdir}/${_pkgname}-${pkgver}"
	sed \
		-e '/Requires:/s/pthread//' \
		-e '/Libs:/s/$/ -lpthread/' \
		-i OpenIPMIpthread.pc.in
	patch -p1 -i ../swig-4.5.patch
	autoreconf -fiv
}

build() {
	cd "${srcdir}/${_pkgname}-${pkgver}"
	PYTHONWARNINGS=ignore ./configure \
		--prefix=/usr \
		--sysconfdir=/etc \

	make
}

package() {
	cd "${srcdir}/${_pkgname}-${pkgver}"
	make DESTDIR="${pkgdir}" install-exec
	make DESTDIR="${pkgdir}" install
	install -Dm644 doc/IPMI.pdf "${pkgdir}/usr/share/doc/${pkgname}/IPMI.pdf"
}
