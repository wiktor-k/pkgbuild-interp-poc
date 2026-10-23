# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Antonio Rojas <arojas@archlinux.org>

pkgname=python-openapi-core
pkgver=0.23.1
pkgrel=1
pkgdesc='Client-side and server-side support for the OpenAPI Specification v3'
arch=(any)
url='https://github.com/p1c2u/openapi-core'
license=(BSD-3-Clause)
depends=(
  python
  python-isodate
  python-jsonschema
  python-jsonschema-path
  python-more-itertools
  python-openapi-schema-validator
  python-openapi-spec-validator
  python-werkzeug
)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
)
optdepends=(
  python-aiohttp
  python-aioitertools
  python-django
  python-falcon
  python-fastapi
  python-flask
  python-multidict
  python-requests
  python-starlette
)
checkdepends=(
  python-aiohttp
  python-aioitertools
  python-django
  python-django-rest-framework
  python-falcon
  python-fastapi
  python-flask
  python-httpx
  python-multidict
  python-pytest
  python-pytest-aiohttp
  python-pytest-asyncio # a dependency of python-pytest-aiohttp
  python-requests
  python-responses
  python-starlette
)
source=("git+$url#tag=$pkgver")
b2sums=('54bf2eea1c57779ea9144bbc56af2cd0ece68a6ef76a9d4057e251947d4a5e1f031949721deb8b2a46e3a7e5b0e9e70de23ddc8bf1256ae2bdfc2bb3c3554a6f')

build() {
  cd "${pkgname#python-}"
  python -m build --wheel --no-isolation
}

check() {
  cd "${pkgname#python-}"
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest --override-ini="addopts="
}

package() {
  cd "${pkgname#python-}"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
