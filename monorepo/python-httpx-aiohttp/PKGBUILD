# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-httpx-aiohttp
pkgver=0.2.0
pkgrel=1
pkgdesc='aiohttp-powered httpx client'
arch=(any)
url=https://github.com/karpetrosyan/httpx-aiohttp
license=(BSD-3-Clause)
depends=(
  python
  python-aiohttp
  python-httpx
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
  python-brotli
  python-brotlicffi
  python-chardet
  python-h2
  python-pytest
  python-trio
  python-trustme
  python-zstandard
  uvicorn
)
source=("git+$url.git#tag=$pkgver")
b2sums=('415fae3b385e7ee91744145e8d03315ae6430bc7c521d60a20cd9768abb19a1372e43eed123ad41789079f9b44fcca6c1df4646edfeb1116291208b0750b068f')

prepare() {
  cd "$srcdir/${pkgname#python-}"
  rm -vr tests/httpx/
  ln -vs /usr/share/python-httpx tests/httpx
}

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python scripts/httpx_test.py \
    -k 'not [trio]' \
    --deselect 'tests/client/test_async_client.py::test_get_invalid_url' \
    --deselect 'tests/client/test_client.py::test_client_decode_text_using_autodetect' \
    --deselect 'tests/client/test_client.py::test_client_decode_text_using_explicit_encoding' \
    --deselect 'tests/client/test_proxies.py::test_async_proxy_close[asyncio]' \
    --deselect 'tests/client/test_proxies.py::test_proxies_environ' \
    --deselect 'tests/client/test_proxies.py::test_socks_proxy' \
    --deselect 'tests/models/test_responses.py::test_response_decode_text_using_autodetect' \
    --deselect 'tests/test_main.py' \
    --deselect 'tests/test_timeouts.py' \
    --deselect 'tests/test_utils.py::test_logging_redirect_chain' \
    --deselect 'tests/test_utils.py::test_logging_request'
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
