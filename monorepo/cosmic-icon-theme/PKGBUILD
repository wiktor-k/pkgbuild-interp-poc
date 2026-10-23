# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Maintainer: Peter Jung <ptr1337@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@proton.me>

pkgname=cosmic-icon-theme
pkgver=1.10.0
pkgrel=1
epoch=1
pkgdesc='Cosmic icon theme'
arch=(any)
url=https://github.com/pop-os/cosmic-icons
license=(CC-BY-SA-4.0)
depends=(pop-icon-theme)
makedepends=(
  git
  just
)
options=(!strip)
source=(git+https://github.com/pop-os/cosmic-icons.git#tag=epoch-${pkgver})
b2sums=('351dc348798d57ef4e4bc67cdd6f8e4fc339c0700c9d1b36334a6c8e11e394fe059f1592362a9923ff758f534ffc9d57f31021ab0831ca55d6cfe95333ccbb6c')

package() {
  cd cosmic-icons
  just rootdir="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
