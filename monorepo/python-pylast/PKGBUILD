# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Steven Allen <steven@stebalien.com>
# Contributor: Limao Luo <luolimao+AUR@gmail.com>
# Contributor: Wieland Hoffmann <the_mineo@web.de>
# Contributor: Amr Hassan <amr.hassan@gmail.com>

pkgname=python-pylast
pkgver=7.2.0
pkgrel=1
pkgdesc='A Python interface to Last.fm and Libre.fm'
arch=(any)
url='https://github.com/pylast/pylast'
license=(Apache-2.0)
depends=(
  python
  python-httpx2
)
makedepends=(
  git
  python-build
  python-hatch-vcs
  python-hatchling
  python-installer
)
source=("$pkgname::git+https://github.com/pylast/pylast.git#tag=$pkgver")
sha512sums=('23d8ef4de9e1f61bb40f98735f4a9b210c37a6f4eb921950896fc4c995f6c93053fade678bb3906c2056c8d3f5467ed4c0326d2c38594005f4a969ef1bbf0cc1')
b2sums=('9e01e70c0c40849bd44ee41103f4fd18f0a9a73bf8ed8e2b5cdc29b21a86d83598219b9d64ec8c733cff4264ccf407066251c8375c76a4d3925cc05bf05344c2')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
# vim: ts=2 sw=2 et:
