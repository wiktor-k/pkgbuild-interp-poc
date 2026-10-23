# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=pysolfc
pkgver=3.6.1
pkgrel=1
pkgdesc="PySol (Fan Club edition) is an exciting collection of more than 1000 solitaire games"
arch=('any')
url="http://pysolfc.sourceforge.net/"
license=('GPL-3.0-or-later')
depends=(
  'hicolor-icon-theme'
  'pysolfc-cardsets'
  'python'
  'python-attrs'
  'python-configobj'
  'python-pillow'
  'python-pygame'
  'python-pysol_cards'
  'tk'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
optdepends=('python-kivy: Kivy based interface')
source=("https://downloads.sourceforge.net/$pkgname/$pkgname-$pkgver.tar.xz")
b2sums=('f828ec45d5c141285e0ce5970af38757704a5d9121cc87e7cbeb83b92f4bac8f3f617c6dc5bcf2477e990fc12230ce988657bff0ab3909397a49da50307d2fae')

prepare() {
  cd $pkgname-$pkgver
  mv pysol.py pysol
  sed -i 's/pysol.py/pysol/' \
    MANIFEST.in \
    data/pysol.desktop \
    setup.py \
    tests/lib/pysol_tests/test_version_flag.py
}

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname-$pkgver
  python -m unittest discover -v
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
