# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Simon Sapin <simon dot sapin at exyr dot org>
# Contributor: Kyle Keen <keenerd@gmail.com>

pkgname=python-factory-boy
pkgver=3.3.3
pkgrel=2
pkgdesc="A versatile test fixtures replacement based on thoughtbot's factory_bot for Ruby"
url="https://github.com/FactoryBoy/factory_boy"
license=('MIT')
arch=('any')
depends=('python-faker')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-pytest' 'python-flask' 'python-django' 'python-sqlalchemy'
              'python-sqlalchemy-utils' 'python-flask-sqlalchemy' 'python-mongoengine'
              'python-mongomock')
source=("$pkgname-$pkgver.tar.gz::https://github.com/FactoryBoy/factory_boy/archive/$pkgver.tar.gz")
sha512sums=('484b4b3db3bb88ccc63793ea5630c5b138208d9a802a2be77ed349fab504c0e548b9b4fa6e80615913c917532f3b50e8018645122d1e90acaf170666d4f8dd5f')

build() {
  cd factory_boy-$pkgver
  python -m build -nw
}

check() {
  cd factory_boy-$pkgver
  python -m installer -d tmp_install dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  PYTHONPATH=tmp_install/"${site_packages}" pytest
}

package() {
  cd factory_boy-$pkgver
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
