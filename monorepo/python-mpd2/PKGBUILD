# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Jörg Thalheim <joerg@higgsboson.tk>
# Contributor: Emmanuel Gil Peyrot <linkmauve@linkmauve.fr>
# Contributor: Peter Richard Lewis <plewis@aur.archlinux.org>
# Contributor: Wieland Hoffmann <the_mineo@web.de>
# Contributor: stonecrest <stonecrest[at]gmail[dot]com>

pkgname=python-mpd2
pkgver=3.1.2
pkgrel=2
pkgdesc='Python library which provides a client interface for the Music Player Daemon'
url=https://github.com/Mic92/python-mpd2
license=(GPL-3.0-only)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
arch=(any)
_tag=abd04a5564fd038181fa6bc5e7afefdc63a06122
source=(git+https://github.com/Mic92/python-mpd2.git#tag=${_tag})
sha256sums=('89a87053e82f47a2c906b0c03beb4a76d69aefbb89a74015f1f3566a0fdb08e9')

pkgver() {
  cd python-mpd2
  git describe --tags | sed 's/^v//'
}

build() {
  cd python-mpd2
  python -m build --wheel --no-isolation
}

package() {
  python -m installer --destdir="${pkgdir}" python-mpd2/dist/*.whl
}

# vim: ts=2 sw=2 et:
