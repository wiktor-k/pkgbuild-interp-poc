# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: T.J. Townsend <blakkheim@archlinux.org>
# Contributor: Jaroslav Lichtblau <dragonlord@aur.archlinux.org>
# Contributor: Jens Adam <jra@byte.cx>
# Contributor: Abel Stern <abel.stern AT gmail.com>
# Contributor: Anton Bazhenov <anton.bazhenov at gmail>

pkgname=byacc
pkgver=20260126
pkgrel=1
pkgdesc='Berkeley reimplementation of the Unix parser generator Yacc'
url='https://invisible-island.net/byacc/'
arch=('x86_64')
license=('custom')
depends=('glibc')
makedepends=('git')
replaces=('byacc-noconflict')
source=(${pkgname}::git+https://github.com/ThomasDickey/byacc-snapshots.git#tag=t${pkgver})
sha512sums=('f95631b5cd1a52b31e0750a255d3d28a62f9b45031d1bb40fa1d22bdf12dba92364e35567dc4f95c21c5ab8cc0a0d322be6ac47055483075b34e50fe8e3ae847')
#validpgpkeys=('19882D92DDA4C400C22C0D56CC2AF4472167BE03') # Thomas Dickey <dickey@his.com>

build() {
  cd ${pkgname}
  sed -n '/is distributed/,/distributed freely/p' README > LICENSE
  ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --program-transform=s,^,b, \
    --enable-btyacc
  make
}

package() {
  cd ${pkgname}
  make DESTDIR="${pkgdir}" install
  install -Dm 644 README -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim:set ts=2 sw=2 et:
