# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Sven Schneider <archlinux.sandmann@googlemail.com>

pkgname=python-graphene
pkgver=3.4.3
pkgrel=4
pkgdesc='GraphQL framework for Python'
arch=(any)
url=https://graphene-python.org/
license=(MIT)
depends=(
  python
  python-dateutil
  python-graphql-core
  python-graphql-relay
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=(python-graphene::git+https://github.com/graphql-python/graphene.git#tag=v${pkgver})
b2sums=('7a117ab9dec1b011f5828e448fcebe8c2ff18d0f4793edf62d6d3aae9342b717b441913ca372eab7d3b317d1a8de43b5c8cd9a4c90ab61f8cdefeecf94909f85')

build() {
  cd python-graphene
  python -m build --wheel --no-isolation
}

package() {
  cd python-graphene
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 LICENSE -t "${pkgdir}"/usr/share/licenses/python-graphene
}
