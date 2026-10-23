# Maintainer: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>

pkgname=python-openapi-schema-validator
# https://github.com/p1c2u/openapi-schema-validator/releases
pkgver=0.9.0
pkgrel=1
pkgdesc="OpenAPI schema validation for Python"
url="https://github.com/p1c2u/openapi-schema-validator"
# https://github.com/p1c2u/openapi-schema-validator/blob/0.7.0/pyproject.toml#L36
license=('BSD-3-Clause')
arch=('any')
depends=('python' 'python-jsonschema' 'python-jsonschema-specifications' 'python-rfc3339-validator' 'python-pydantic' 'python-pydantic-settings')
makedepends=('python-build' 'python-installer' 'python-poetry-core')
checkdepends=('python-pytest')
source=("https://github.com/p1c2u/openapi-schema-validator/archive/$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('723108995e01e4ce9f49e526d0ac1b0211aec0b34913b3913dde18c879c44264')

prepare() {
  cd openapi-schema-validator-$pkgver
  sed -i '/--cov/d' pyproject.toml
}

build() {
  cd openapi-schema-validator-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd openapi-schema-validator-$pkgver
  python -m pytest
}

package() {
  cd openapi-schema-validator-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
