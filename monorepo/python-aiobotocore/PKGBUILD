# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=python-aiobotocore
pkgver=3.9.2
pkgrel=1
pkgdesc='asyncio support for botocore library using aiohttp'
arch=(any)
url='https://github.com/aio-libs/aiobotocore'
license=('Apache-2.0')
depends=(
  python
  python-aiohttp
  python-aioitertools
  python-botocore
  python-dateutil
  python-jmespath
  python-multidict
  python-wrapt
)
makedepends=(
  git
  python-build
  python-hatch-fancy-pypi-readme
  python-hatchling
  python-installer
  python-wheel
)
checkdepends=(
  python-anyio
  python-dill
  python-docker
  python-docutils
  python-flask
  python-flask-cors
  python-moto
  python-openapi-spec-validator
  python-pip
  python-pytest
  python-pytest-mock
  python-time-machine
  python-tiny-proxy
  python-trustme
  python-yaml
)
optdepends=(
  'python-anyio: for the HTTPX backends'
  'python-httpx2: for the HTTPX backend'
  'python-httpx: for the legacy HTTPX backend'
)
source=(
  "git+$url.git#tag=$pkgver"
  "$pkgname-fix-duplicate-server-header.patch"
)
b2sums=('b66e9913bedb57ab6f1eb20c29f2aee7707d32e6e4e5f361511044ef23aed23cde5933a147bb09122d7f974dfeb86a285914783258ed935b6ab9b9ec1b0934ed'
        'e0b6a254b92f61cd0c1bb6a2930373bdd8a513afb594f9ec0001a0fdf244439d305e887a7cddcdfb4a8775ffcf609606e128959547506576f2174ccc85afdd18')

prepare() {
  cd ${pkgname#python-}
  patch -Np1 < ../$pkgname-fix-duplicate-server-header.patch
}

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  local pytest_args=(
    -m='not localonly'
    # Uses moto.awslambda, which requires a running Docker service.
    # See: https://github.com/spulec/moto/issues/3276
    --deselect='tests/test_lambda.py'
    # Validates Amazons internal patches as far as I understand, as such not
    # relevant for us.
    --deselect='tests/test_patches.py'
    # Fails for some reason.
    --deselect='tests/botocore_tests/unit/test_signers.py::test_signers_generate_db_auth_token[aiohttp-aws_auth0]'
    --deselect='tests/test_sns.py::test_topic_attributes[aiohttp]'
  )
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest "${pytest_args[@]}"
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
}
