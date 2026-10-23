# Maintainer: Christian Heusel <christian@heusel.eu>
# Contributor: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: arkhan <arkhan@disroot.org>

pkgname=ueberzug
pkgver=18.3.2
pkgrel=1
pkgdesc='Command line util which allows to display images in combination with X11'
arch=('x86_64')
url=https://github.com/ueber-devel/ueberzug
license=('GPL-3.0-or-later')
options=(!lto)
depends=(
    'libxext'
    'libxres'
    'python-attrs'
    'python-docopt'
    'python-pillow'
)
makedepends=(
    'meson-python'
    'python-build'
    'python-installer'
    'python-wheel'
)
provides=("python-$pkgname")
replaces=("python-$pkgname")
source=("https://github.com/ueber-devel/ueberzug/archive/$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('c30c672611f1d5661916166ff6bf21c348354600d6662cb13823162b28cf9e78')
b2sums=('530a33054b8d5c39ba7c553e9aac319ba4947212ebd88a913115fb23ec62cc9bc7cd8da9e87cf2c22ccb7fdf9173e7be4d827668232305e3112136b49d66c255')

build() {
    cd $pkgname-$pkgver
    # set a fixed build dir to fix repro issue, see
    # https://github.com/mesonbuild/meson-python/issues/703 for details
    python -m build --wheel --no-isolation -Cbuild-dir=build
}

package() {
    cd $pkgname-$pkgver
    python -m installer --destdir="$pkgdir" dist/*.whl
}
