# Maintainer:
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=cln
pkgver=1.3.7
pkgrel=3
pkgdesc="Class library for numbers"
arch=(x86_64)
url="https://www.ginac.de/CLN/"
license=(GPL-2.0-or-later)
depends=(
  glibc
  gmp
  libgcc
  libstdc++
)
makedepends=('texlive-latex')
source=(https://www.ginac.de/CLN/${pkgname}-${pkgver}.tar.bz2)
sha256sums=('7c7ed8474958337e4df5bb57ea5176ad0365004cbb98b621765bc4606a10d86b')

build() {
  cd ${pkgname}-${pkgver}
  ./configure --prefix=/usr
  make all html pdf
}

check() {
  cd ${pkgname}-${pkgver}
  make check
}

package() {
  cd ${pkgname}-${pkgver}
  make DESTDIR="${pkgdir}" install install-html install-pdf
}
