# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@tutanota.com>

pkgname=python-moddb
pkgver=0.15.0
pkgrel=1
pkgdesc='A Python scraper to access ModDB mods, games and more as objects'
arch=(any)
url=https://github.com/ClementJ18/moddb
license=(MIT)
depends=(
  python
  python-beautifulsoup4
  python-curl-adapter
  python-requests
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-packaging
  python-setuptools
  python-wheel
)
#source=(git+https://github.com/ClementJ18/moddb.git#tag=v${pkgver})
source=(git+https://github.com/ClementJ18/moddb.git#commit=c94f7d826532eebd9b3d479552f8be0f3cb80b57)
b2sums=('c23d9aa3933b51438c41929750fcdbf0d3bdcc7557789fbdcfe83cbfc7e0baca6f748ac95a3a1b417a231ad45766c9a5f8fa33e4760381737c2b09dafb998be9')

build() {
  cd moddb
  python -m build --wheel --no-isolation
}

package() {
  python -m installer --destdir="${pkgdir}" moddb/dist/*.whl
  install -Dm 644 moddb/LICENSE -t "${pkgdir}"/usr/share/licenses/python-moddb/
}

# vim: ts=2 sw=2 et:
