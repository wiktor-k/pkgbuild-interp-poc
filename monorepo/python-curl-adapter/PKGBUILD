# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Mark Wagie <mark.wagie@proton.me>

pkgname=python-curl-adapter
pkgver=1.2.3
pkgrel=1
pkgdesc='A curl HTTP adapter switch for requests library — make browser-like requests with custom TLS fingerprints'
arch=(any)
url=https://github.com/el1s7/curl-adapter
license=(MIT)
depends=(
  python
  python-brotli
  python-curl_cffi
  python-eventlet
  python-pycurl
  python-requests
  python-urllib3
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
optdepends=('python-gevent: gevent support')
source=(git+https://github.com/el1s7/curl-adapter.git#tag=v${pkgver})
b2sums=('3ad7aa044ff43d86897b59a6074f2af64344f11655fd02e27f4c7b62c6e940b61302afbc863144c91e5ac84a7eb3acf913b399648949764fe4481ef5ab28915c')

build() {
  cd curl-adapter
  python -m build --wheel --no-isolation
}

check() {
  cd curl-adapter
  pytest \
    --deselect tests/test_general.py::TestFunctions::test_ssl_verify \
    --deselect tests/test_general.py::TestFunctions::test_ssl_no_verify
}

package() {
  python -m installer --destdir="${pkgdir}" curl-adapter/dist/*.whl
  install -Dm 644 curl-adapter/LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}/"
}
