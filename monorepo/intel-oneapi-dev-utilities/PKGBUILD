# Maintainer: Torsten Keßler <tpkessler at archlinux dot org>
# Contributor: Jingbei Li <i@jingbei.li>

pkgname=intel-oneapi-dev-utilities
pkgver=2026.0.1_16
_pkgver=$(echo $pkgver | sed 's/_/-/')
_majmin=$(echo $pkgver | cut -d. -f1,2)
pkgrel=1
pkgdesc="Intel oneAPI Dev Utilities"
arch=('any')
url='https://software.intel.com/content/www/us/en/develop/tools/oneapi.html'
license=("LicenseRef-Intel-EULA-Developer-Tools")
source=("https://apt.repos.intel.com/oneapi/pool/main/${pkgname}-${_pkgver}_amd64.deb"
        "https://apt.repos.intel.com/oneapi/pool/main/${pkgname}-eclipse-cfg-${_majmin}-${_pkgver}_all.deb")
b2sums=('df29769e4128086ab0892582b66a62b99d04b31c2810ce0b46e86f239f9bc150f41310d2e0c3892c9c2ccf28e8b8c2b281b24479fbee3624f6b3b9287ae6acdb'
        '12edb7f532fda600528bc652aa22f0c2cf14973a38efb596057ac899a953fc8b3e74eff4ba3436d98a9baafe008c9a016cb598184dc5d3ae222290c833f04e25')
depends=('intel-oneapi-common')
provides=("${pkgname}-eclipse-cfg")
conflicts=('intel-oneapi-basekit')

noextract=(
	"${pkgname}-${_pkgver}_amd64.deb"
	"${pkgname}-eclipse-cfg-${_majmin}-${_pkgver}_all.deb"
)

package() {
  cd "${srcdir}"
  ar x ${pkgname}-${_pkgver}_amd64.deb
  tar xvf data.tar.xz -C "${pkgdir}"
  rm data.tar.xz

  ar x ${pkgname}-eclipse-cfg-${_majmin}-${_pkgver}_all.deb
  tar xvf data.tar.xz -C "${pkgdir}"
  rm data.tar.xz

  local _dev_path='/opt/intel/oneapi/dev-utilities'
  ln -s "${_dev_path}/${_majmin}" "${pkgdir}/${_dev_path}/latest"

  install -d "${pkgdir}"/usr/share/licenses/"${pkgname}"
  ln -s /usr/share/licenses/intel-oneapi "${pkgdir}"/usr/share/licenses/"${pkgname}"/oneapi
}
