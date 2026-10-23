# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Claudia Pellegrino <aur ät cpellegrino.de>
# Contributor: Felix Golatofski <contact@xdfr.de>
# Contributor: Johannes Dewender   arch at JonnyJD dot net

pkgname=distro-info-data
pkgver=2026.08.20
pkgrel=1
pkgdesc="Information about all releases of Debian and Ubuntu (data files)"
arch=('any')
url="https://salsa.debian.org/debian/distro-info-data"
license=('ISC')
makedepends=('git')
checkdepends=('python')
options=(!emptydirs)
source=("git+https://salsa.debian.org/debian/distro-info-data.git#tag=debian/$pkgver-1")
sha256sums=('8a6c004ef74d440b468e403ac2f90959e192bb6563b4dd4de3abdc26ac02a4ea')

check() {
  cd $pkgname
  make test
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
  install -Dm644 debian/copyright "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
  install -Dm644 -t "$pkgdir"/usr/share/doc/$pkgname README.md
}

# vim:set ts=2 sw=2 et:
