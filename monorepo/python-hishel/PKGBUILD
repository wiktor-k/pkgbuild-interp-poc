# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-hishel
_name="${pkgname#python-}"
pkgver=1.4.0
pkgrel=1
pkgdesc="An elegant HTTP Cache implementation for HTTPX and HTTP Core"
arch=(any)
url="https://github.com/karpetrosyan/hishel"
license=(BSD-3-Clause)
depends=(
  python
  python-msgpack
  python-typing_extensions
)
makedepends=(
  python-build
  python-hatchling
  python-hatch-fancy-pypi-readme
  python-installer
)
checkdepends=(
  python-anyio
  python-anysqlite
  python-fakeredis
  python-httpx
  python-httpx2
  python-inline-snapshot
  python-pytest
  python-redis
  python-requests
  python-time-machine
  python-trio
)
optdepends=(
  'python-anyio: for general async as well as httpx and httpx2 support'
  'python-anysqlite: for general async as well as httpx and httpx2 support'
  'python-boto3: for S3 support'
  'python-fastapi: for FastAPI support'
  'python-httpx: for HTTP support with httpx'
  'python-httpx2: for HTTP support with httpx2'
  'python-redis: for Redis support'
  'python-requests: for requests support'
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz)
sha512sums=('255495cea3941778d13525bf8cc25f7ba406a0fc4a6986924110e4239afc79ba1ecc815b604b3b13d4163853837ca7bcda306d1b9431f77abfaa52d7e6e70729')
b2sums=('7d8745660a993b3279153912e488c0d1e14e69e2aa5634ac171af25cecf7c4d8cb9b16bd8e0d0e03c789f8c74a64fdcb83aa74b9dbc1115fb9b993ce1e2792b0')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vvv
  )

  cd $_name-$pkgver
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
