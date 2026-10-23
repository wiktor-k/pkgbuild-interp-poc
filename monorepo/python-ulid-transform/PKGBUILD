# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Lukasz Pozarlik <lpozarlik@gmail.com>

pkgname=python-ulid-transform
pkgver=2.2.9
pkgrel=3
pkgdesc='Create and transform ULIDs'
arch=(x86_64)
url=https://github.com/bdraco/ulid-transform
license=(MIT)
depends=(
  glibc
  libgcc
  libstdc++
  python
)
makedepends=(
  cython
  git
  python-build
  python-installer
  python-poetry-core
  python-setuptools
  python-wheel
)
source=(git+https://github.com/bdraco/ulid-transform.git#tag=v${pkgver})
b2sums=('478386edeb4986eb2cda2b63745f6d9715371f7dd845b27ecd2662212a4706fbd0a571c20c4cc97275a2d52234c6ba84a1f97665edc17615f0074ea944319ed1')


build() {
  cd ulid-transform
  python -m build --wheel --no-isolation
}

package() {
  python -m installer --destdir="${pkgdir}" ulid-transform/dist/*.whl
  install -Dm 644 ulid-transform/LICENSE -t "${pkgdir}"/usr/share/licenses/python-ulid-transform/
}

# vim: ts=2 sw=2 et:
