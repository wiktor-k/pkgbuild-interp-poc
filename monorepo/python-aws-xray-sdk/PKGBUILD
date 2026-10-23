# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Guillaume Horel <guillaume.horel@gmail.com>

pkgname=python-aws-xray-sdk
_pkgname=aws-xray-sdk-python
pkgver=2.15.0
pkgrel=4
pkgdesc='AWS X-Ray SDK for Python'
arch=(any)
url='https://github.com/aws/aws-xray-sdk-python'
license=('Apache-2.0')
depends=(
  python
  python-botocore
  python-wrapt
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-asyncio
  # Needed by tests/test_local_sampling_benchmark.py
  python-pytest-benchmark
)
source=(
  "git+$url.git#tag=$pkgver"
  "$pkgname-pytest-asyncio-compat.patch"
  "$pkgname-fix-asyncio-pytest.patch"
  "$pkgname-remove-pkg_resources.patch"
)
sha256sums=('98db847bce19d2ee82095c28425451397dbcdd0c16f87494fc6157b69c0d7d07'
            '9c0fab3fd7b94e22807913f4b0051a4485c0bf7bbe47baed4be2a1ca3b475a9f'
            'f347f4d8bcae8ef4e3d9a0ab19b882c903481f2ed7d2ee52b67b147c232643cc'
            'd1f752c4a3c25f9aa4ec33e0d498102e636fe99fa75d700545de56906e46ae6d')

prepare() {
  cd $_pkgname
  patch -Np1 < ../$pkgname-pytest-asyncio-compat.patch
  patch -Np1 < ../$pkgname-fix-asyncio-pytest.patch
  patch -Np1 < ../$pkgname-remove-pkg_resources.patch
}

build() {
  cd $_pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pkgname
  # See setenv= in upstream tox.ini
  export AWS_SECRET_ACCESS_KEY=fake_key
  export AWS_ACCESS_KEY_ID=fake_id
  # Ignore ext tests, which get broken from time to time with newer dependencies,
  # and result in a dependency cycle that makes rebuilding for newer Python harder
  pytest -v tests --ignore tests/ext -p pytest_asyncio
}

package() {
  cd $_pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
}
