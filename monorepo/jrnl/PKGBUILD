# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=jrnl
pkgver=4.6
pkgrel=1
pkgdesc="Collect your thoughts and notes without leaving the command line"
arch=('any')
url="https://jrnl.sh/"
license=('GPL-3.0-only')
depends=(
  'python'
  'python-colorama'
  'python-cryptography'
  'python-dateutil'
  'python-keyring'
  'python-parsedatetime'
  'python-pyxdg'
  'python-rich'
  'python-ruamel-yaml'
  'python-tzlocal'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-poetry-core'
)
checkdepends=(
  'python-pytest'
  'python-pytest-bdd'
  'python-pytest-xdist'
)
_url=https://github.com/jrnl-org/jrnl
source=("git+$_url#tag=v$pkgver")
b2sums=('ba07f30925dcbdbbf96b16e3edbddec95d3a45b55baa148bf878646bde05ad7c357764652e3d77892bd0a9da7d14f025693c1aa3e25e44b555eab4d0e0310b90')

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname
  pytest
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
}
