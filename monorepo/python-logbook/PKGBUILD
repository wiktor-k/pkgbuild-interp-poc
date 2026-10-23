# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Tomas Schertel<tschertel@gmail.com>

pkgname=python-logbook
pkgver=1.10.1
pkgrel=1
pkgdesc="Logging sytem for Python that replaces the standard library’s logging module"
arch=('x86_64')
url="https://logbook.readthedocs.io/en/stable/"
license=('BSD-3-Clause')
depends=('python' 'python-typing_extensions')
makedepends=('git' 'python-setuptools' 'python-setuptools-rust' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('valkey' 'python-pytest' 'python-sqlalchemy' 'python-redis' 'python-pyzmq'
              'python-execnet' 'python-jinja' 'python-brotli' 'pifpaf')
source=("git+https://github.com/mitsuhiko/logbook.git#tag=$pkgver")
sha512sums=('95eea24f1e75553f4f8a907874464d4778178e073c2c7329bc8c42cc35eba43b523d4b66f69dd27bf59a4422cc7125125ac10ece44071fee86984c3158f0c47f')

prepare() {
  cd logbook
  python -m build --wheel --no-isolation
}

build() {
  cd logbook
}

check() {
  cd logbook
  LC_CTYPE=en_US.UTF-8 PYTHONDONTWRITEBYTECODE=1 PYTHONPATH=src pifpaf run redis py.test tests
}

package() {
  cd logbook
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm664 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
