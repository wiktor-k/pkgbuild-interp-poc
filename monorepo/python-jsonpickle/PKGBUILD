# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-jsonpickle
pkgver=4.1.4
pkgrel=1
arch=('any')
pkgdesc="Python library for serializing any arbitrary object graph into JSON"
url="https://jsonpickle.github.io/"
license=('BSD-3-Clause')
depends=('python')
optdepends=("python-numpy: for serializing sklearn models, numpy arrays, and other numpy-based data"
            "python-gmpy2: for serializing ecdsa module's keys")
makedepends=('git' 'python-setuptools-scm' 'python-setuptools' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-numpy' 'python-feedparser' 'python-simplejson' 'python-gmpy2'
              'python-pymongo' 'python-ujson' 'python-pandas' 'python-pytest')
source=("git+https://github.com/jsonpickle/jsonpickle.git#tag=v$pkgver")
sha512sums=('09289352b880310563ee5df2883460f4f014b42c9c3eee51f638f517538cad4c70da8cc4b02de321b9250fe4b4bc49928536c587971acff72768a0f4f7701021')

prepare() {
  cd jsonpickle

  # Replace removed NumPy dtype aliases: https://github.com/jsonpickle/jsonpickle/pull/592
  sed -i -e 's/i2,a3,i4/i2,S3,i4/' -e 's/u1,f4,a1/u1,f4,S1/' tests/numpy_test.py
}

build() {
  cd jsonpickle
  python -m build --wheel --no-isolation
}

check() {
  cd jsonpickle
  pytest -W "ignore:keys will default to True in jsonpickle 5.0.0:DeprecationWarning" tests
}

package() {
  cd jsonpickle
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
