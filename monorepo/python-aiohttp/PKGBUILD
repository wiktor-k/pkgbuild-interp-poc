# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Jonas Witschel <diabonas@archlinux.org>
# Contributor: Philipp A. <flying-sheep@web.de>

# Check if new updates break python-engineio

pkgname=python-aiohttp
pkgver=3.14.3
pkgrel=1
pkgdesc='HTTP client/server for asyncio'
arch=(x86_64)
url=https://aiohttp.readthedocs.io
license=(Apache-2.0)
depends=(
  glibc
  llhttp
  python
  python-aiohappyeyeballs
  python-aiosignal
  python-attrs
  python-frozenlist
  python-multidict
  python-propcache
  python-yarl
)
makedepends=(
  cython
  git
  npm
  python-build
  python-installer
  python-pkgconfig
  python-setuptools
  python-wheel
)
checkdepends=(
  gunicorn
  python-aiodns
  python-blockbuster
  python-brotli
  python-coverage
  python-freezegun
  python-isal
  python-proxy.py
  python-pytest
  python-pytest-codspeed
  python-pytest-mock
  python-pytest-timeout
  python-pytest-xdist
  python-re-assert
  python-time-machine
  python-trustme
  python-uvloop
  python-zlib-ng
)
optdepends=(
  'gunicorn: to deploy using Gunicorn'
  'python-aiodns: for fast DNS resolving'
  'python-brotli: for Brotli transfer-encodings support'
)
source=("$pkgname::git+https://github.com/aio-libs/aiohttp#tag=v$pkgver")
b2sums=('fc16578f3c41ebfc9bba6cd9eecf555f2206292e0a167684660d92102a7f0ad2b6c20276e77a6668623370c4a047ab551a85d6d7069f888c0dcff86e9ab92b68')

prepare() {
  cd $pkgname
  sed 's|.install-cython ||' -i Makefile
  # Fix Cython 3.3.0 compiler crash when compiling the WebSocket reader
  git cherry-pick -n baa2709c4ca2365139d335b8542eb87e0ab6e150
}

build() {
  cd $pkgname
  export AIOHTTP_USE_SYSTEM_DEPS=1
  make cythonize
  python -m build --wheel --no-isolation
}

check() {
  local pytest_args=(
    -vvv
    --override-ini="addopts="

    # Fails, not sure why.
    --deselect=tests/test_http_parser.py::test_http_response_parser_bad_chunked_strict_c
    --deselect=tests/test_http_parser.py::test_http_response_parser_bad_chunked_strict_py
    --deselect=tests/test_http_parser.py::test_http_response_parser_strict_headers
    --deselect=tests/test_http_parser.py::test_http_response_parser_strict_obs_line_folding
    --deselect=tests/test_proxy_functional.py::test_uvloop_secure_https_proxy
    --deselect='tests/test_client_functional.py::test_invalid_idna[pyloop]'

    # Fails, calls the Python interpreter without -P, deselect to enable
    # venv-based testing.
    --deselect=tests/test_circular_imports.py::test_no_warnings

    --ignore=tests/autobahn/test_autobahn.py # Docker tests
    --ignore=tests/test_benchmarks_client.py
    --ignore=tests/test_benchmarks_client_request.py
    --ignore=tests/test_benchmarks_client_ws.py
    --ignore=tests/test_benchmarks_cookiejar.py
    --ignore=tests/test_benchmarks_http_websocket.py
    --ignore=tests/test_benchmarks_http_writer.py

    # Tests freeze and timeout when running serialized, no idea why
    -n auto

    -W ignore::DeprecationWarning
  )

  cd $pkgname
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -Pm pytest "${pytest_args[@]}"
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim: ts=2 sw=2 et:
