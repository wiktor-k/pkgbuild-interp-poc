# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Maintainer: Peter Jung <ptr1337@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@proton.me>

pkgname=cosmic-wallpapers
pkgver=1.10.0
pkgrel=1
epoch=2
pkgdesc='Wallpapers for the COSMIC Desktop Environment'
arch=(any)
url=https://github.com/pop-os/cosmic-wallpapers
license=(LicenseRef-unknown)
groups=(cosmic)
makedepends=(
  git
  git-lfs
)
source=(git+https://github.com/pop-os/cosmic-wallpapers.git#tag=epoch-${pkgver})
b2sums=('3417048472aa110625c5c18164f8d417b5a8624f111e17d682bfc7ce11322a177e207ef8ba3cf16aabe38a9c9bca70c6c1bb97390c2bcda3ebf09eec3d5d92e9')

prepare() {
  cd cosmic-wallpapers
  git lfs install --local
  git remote add network-origin ${url}
  git lfs fetch network-origin
  git lfs checkout
}

package() {
  cd cosmic-wallpapers
  make DESTDIR="${pkgdir}" install
}

# vim: ts=2 sw=2 et:
