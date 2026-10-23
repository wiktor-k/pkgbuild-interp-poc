# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Caltlgin Stsodaat <contact@fossdaily.xyz>
# Contributor: Ekin Dursun <ekindursun@gmail.com>

pkgname=python-mpv
pkgver=1.0.8
pkgrel=2
pkgdesc='Python interface to the awesome mpv media player'
arch=(any)
url=https://github.com/jaseg/python-mpv
license=('GPL-2.0-or-later OR LGPL-2.0-or-later')
depends=(
  mpv
  python
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pyvirtualdisplay
  xorg-server-xvfb
)
optdepends=('python-pillow: for raw screenshot support')
source=("git+$url.git#tag=v$pkgver?signed")
b2sums=('401535f88b7214b6d370d85388f85d95fa4b4d4b72ca772dba94156b1a9dd627a1adb4b3ce657b331ab52de72b41aa83ea7ebe49baa63a63a2b0035d19dffddf')
validpgpkeys=(E36F75307F0A0EC2D145FF5CED7A208EEEC76F2D) # python-mpv Release Signing Key <python-mpv@jaseg.de>

build() {
  cd $pkgname
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd $pkgname
  pytest
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 -t "$pkgdir"/usr/share/doc/$pkgname README.rst
}
