# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-qt-material
pkgver=2.17
pkgrel=2
pkgdesc='Material inspired stylesheet for PySide & PyQt'
arch=('any')
url='https://qt-material.readthedocs.io/'
license=('BSD-2-Clause')
depends=(
  'python'
  'python-jinja'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-wheel'
  'python-setuptools'
)
optdepends=(
  'pyside6: Use PySide for Qt bindings'
  'python-pyqt6: Use PyQt for Qt bindings'
)
source=("$pkgname::git+https://github.com/dunderlab/qt-material#tag=v$pkgver")
sha512sums=('d3a1401fe322b06ec2ab707aec2f8fd67c75d81a47e7530dda211632b33145056b5b899acfd58e579873c09f5988d28a7fb85b5196a9ee5000d4dc50d05abd74')
b2sums=('bede50048f8ec00164a241b8b18dfcf296ac1737a77a0040cc147e23b8a98b28eea25dacfdc3714c6fc16b855d35327c1a3f0f79b4e42eb0e21d70db029da431')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
