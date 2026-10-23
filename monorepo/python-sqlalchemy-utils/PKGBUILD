# Maintainer: Jelle van der Waa <jelle@archlinux.org
# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Husam Bilal <husam212 AT gmail DOT com>

_name=sqlalchemy-utils
pkgname=python-sqlalchemy-utils
pkgver=0.43.0
pkgrel=1
epoch=1
pkgdesc='Various utility functions, new data types and helpers for SQLAlchemy'
url='https://github.com/kvesteri/sqlalchemy-utils'
depends=('python' 'python-sqlalchemy' 'python-babel' 'python-arrow'
         'python-intervals' 'python-phonenumbers' 'python-passlib' 'python-colour'
         'python-dateutil' 'python-furl' 'python-cryptography' 'python-pendulum')
checkdepends=('python-pytest' 'python-flexmock' 'python-jinja')
makedepends=('python' 'python-setuptools' 'python-build' 'python-installer' 'python-wheel')
license=('BSD-3-Clause')
arch=('any')
source=(https://github.com/kvesteri/sqlalchemy-utils/archive/${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('79041c594d26c1c3584d4ec930c59f861a10ba3d5ef330e0c1da3268c787a03a24d2ee32df4a3f3687710506726bcda36a743d15d1d48089b68d856a9f228274')
b2sums=('83a808229a672448f07e1d6d88285662c863fb1aecad515698ac52f8a989450432554e2604ef58b6921f132a7dce77bcdb67d652a14b4832769829e0930e3195')

build() {
  cd ${_name}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${_name}-${pkgver}
  # Tests require a postgres and MySQL db
  export PYTHONWARNINGS="ignore"
  pytest -Wignore --disable-pytest-warnings tests/test_models.py
}

package() {
  cd ${_name}-${pkgver}
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
