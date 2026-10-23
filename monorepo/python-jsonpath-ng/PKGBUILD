# Maintainer: Chih-Hsuan Yen <yan12125@archlinux.org>
# Adopted from aur/python-jsonpath-ng, original contributors:
# Contributor: Paul Irofti <paul@irofti.net>

pkgname=python-jsonpath-ng
pkgver=1.8.0
pkgrel=1
pkgdesc="JSONPath for Python that aims to be standard compliant"
arch=('any')
url="https://github.com/h2non/jsonpath-ng"
license=('Apache-2.0')
depends=('python')
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-hypothesis'
  'python-pytest'
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('90786a347999bd644ad8e6fd5fb6d857c05295e899f0430fbd75dc1741cf2f00cbb8e5f39ce182417154ee12f523ee938e74656b271363a498ad3995af8ca3d2')

build() {
  cd "${pkgname#python-}"
  python -m build --wheel --no-isolation
}

check() {
  cd "${pkgname#python-}"
  pytest
}

package() {
  cd "${pkgname#python-}"
  python -m installer --destdir="$pkgdir" dist/*.whl
}
