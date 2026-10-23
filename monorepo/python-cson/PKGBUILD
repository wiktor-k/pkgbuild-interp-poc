# Maintainer: Claudia Pellegrino <auerhuhn@archlinux.org>
# Contributor: Rémy Oudompheng <remy@archlinux.org>

pkgname=python-cson
pkgver=0.8
pkgrel=12
pkgdesc='A Python parser for Coffeescript Object Notation (CSON)'
arch=('any')
url='https://github.com/avakar/pycson'
license=('MIT')
depends=('python' 'python-speg')
makedepends=('git'
             'python-build'
             'python-installer'
             'python-setuptools'
             'python-wheel')
checkdepends=('python-pytest')
source=("git+https://github.com/avakar/pycson.git#tag=${pkgver}")
sha256sums=('f957146918a0c2245346f2da63a95cf2e5ca636707287ed59df1e241161b8461')

build() {
  cd pycson
  python -m build --wheel --no-isolation
}

check() {
  cd pycson
  PYTHONPATH=. py.test
}

package() {
  cd pycson
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -D -m644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
