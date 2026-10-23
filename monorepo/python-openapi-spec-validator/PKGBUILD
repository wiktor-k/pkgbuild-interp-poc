# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>

pkgname=python-openapi-spec-validator
pkgver=0.9.0
pkgrel=1
pkgdesc="OpenAPI 2.0 (aka Swagger) and OpenAPI 3 spec validator"
url="https://github.com/p1c2u/openapi-spec-validator"
license=('Apache-2.0')
arch=('any')
depends=(
  'python'
  'python-jsonschema'
  'python-jsonschema-path'
  'python-lazy-object-proxy'
  'python-openapi-schema-validator'
  'python-pydantic'
  'python-pydantic-settings'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-poetry-core'
)
checkdepends=('python-pytest')
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('2652f1743596b8fb5cada7487a6880d9bd9122d63ca221448cd7b61ec87df69495b5f7591064ae8cb756647ad990ecef4a53504848b573dc2dda1978fd2dc5de')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  PYTHONPATH="$PWD" pytest --override-ini="addopts="
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
