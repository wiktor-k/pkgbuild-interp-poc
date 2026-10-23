# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Evgeniy Filimonov <evgfilim1@gmail.com>

pkgname=python-aiogram
pkgver=3.31.0
pkgrel=1
pkgdesc="A modern and fully asynchronous framework for Telegram Bot API"
arch=('any')
url="https://github.com/aiogram/aiogram"
license=('MIT')
depends=(
  'python'
  'python-aiofiles'
  'python-aiohttp'
  'python-certifi'
  'python-magic-filter'
  'python-pydantic'
  'python-typing_extensions'
)
makedepends=(
  'git'
  'python-build'
  'python-hatchling'
  'python-installer'
  'python-wheel'
)
checkdepends=(
  'pifpaf'
  'python-aiohttp-socks'
  'python-aresponses'
  'python-babel'
  'python-cryptography'
  'python-motor'
  'python-pycryptodomex'
  'python-pymongo'
  'python-pytest'
  'python-pytest-aiohttp'
  'python-pytest-lazy-fixture'
  'python-pytz'
  'python-redis'
  'valkey'
)
optdepends=(
  'python-uvloop: fast, drop-in replacement of the built-in asyncio event loop'
  'python-ujson: ultra fast JSON encoder and decoder written in pure C'
  'python-rapidjson: extremely fast C++ JSON parser and serialization library'
  'python-emoji: emojize and demojize support'
  'python-redis: Redis storage support'
  'python-aiohttp-socks: SOCKS4(a) and SOCKS5 proxy support'
  'python-motor: MongoDB storage support'
  'python-babel: i18n support'
  'python-cryptography: webhook signature verification support'
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('7d770e3b2d2d34735bb300a20d10f45d4a385c2da92993e2a6abc43b0541b993dc387769117a8b4bb0e07f09e8a556d17a069dbc22c9fec6ba6b095c0642a4ca')

build() {
  cd aiogram
  python -m build --wheel --no-isolation
}

check() {
  cd aiogram
  # shellcheck disable=2016
  pifpaf run redis -- bash -c 'python -m pytest --redis $PIFPAF_REDIS_URL'
}

package() {
  cd aiogram
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" README.rst
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
