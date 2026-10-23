# Maintainer: Torsten Keßler <tpkessler@archlinux.org>
pkgname=intel-oneapi-umf
pkgver=1.1.0_340
_pkgver=$(echo $pkgver | sed 's/_/-/')
_majmin=$(echo $pkgver | cut -d. -f1,2)
pkgrel=2
pkgdesc="Unified Memory Framework"
arch=('x86_64')
url="https://www.intel.com/content/www/us/en/developer/tools/oneapi/overview.html"
license=('LicenseRef-Intel-Simplified')
depends=('intel-oneapi-common' 'glibc' 'gcc-libs' 'hwloc')
source=("https://apt.repos.intel.com/oneapi/pool/main/${pkgname}-${_majmin}-${_pkgver}_amd64.deb")
b2sums=('9866450bc041a9a6956c756f82eb2cecb153bc2a3abb0870cc642c316a4caee90bfbfdfd4cbe9cd4c9998efb4dd8950544577a978ea999ab4e23656eec312911')
noextract=("${pkgname}-${majmin}-${_pkgver}_amd64.deb")
conflicts=('intel-oneapi-basekit')

package() {
	ar x "${pkgname}-${_majmin}-${_pkgver}_amd64.deb"
	tar xvf data.tar.xz -C "${pkgdir}"
	rm data.tar.xz

	local _prefix=/opt/intel/oneapi/umf
	ln -vs "${_prefix}/${_majmin}" "${pkgdir}/${_prefix}/latest"

	install -d "${pkgdir}"/usr/share/licenses/"${pkgname}"
	ln -vs /usr/share/licenses/intel-oneapi "${pkgdir}"/usr/share/licenses/"${pkgname}"/oneapi
}
