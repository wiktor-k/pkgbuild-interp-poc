# Maintainer: Maxime Gauduin <alucryd@archlinux.org>

pkgname=python-graphql-relay
pkgver=3.2.0
pkgrel=8
pkgdesc='A library to help construct a graphql-py server supporting react-relay'
url=https://github.com/graphql-python/graphql-relay-py
arch=(any)
license=(MIT)
depends=(
  python
  python-graphql-core
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
  python-setuptools
  python-wheel
)
_tag=97ba15e573eabcefd13e249d05fc5557f0296804
source=(git+https://github.com/graphql-python/graphql-relay-py.git#tag=${_tag})
sha256sums=(SKIP)

pkgver() {
  cd graphql-relay-py
  git describe --tags | sed 's/^v//'
}

prepare() {
  cd graphql-relay-py
  sed -i 's/setuptools>=59,<70/setuptools/' pyproject.toml
  sed -i 's/poetry_core>=1,<2/poetry_core/' pyproject.toml
}

build() {
  cd graphql-relay-py
  python -m build --wheel --no-isolation
}

package() {
  python -m installer --destdir="${pkgdir}" graphql-relay-py/dist/*.whl
  install -Dm 644 graphql-relay-py/LICENSE -t "${pkgdir}"/usr/share/licenses/python-graphql-relay/
}

# vim: ts=2 sw=2 et:
