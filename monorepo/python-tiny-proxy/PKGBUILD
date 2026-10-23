# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-tiny-proxy
pkgver=0.3.0
pkgrel=1
pkgdesc='Simple proxy (SOCKS4(a), SOCKS5(h), HTTP tunnel) server built with anyio'
arch=(any)
url='https://github.com/romis2012/tiny-proxy'
license=(Apache-2.0)
depends=(python python-anyio)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-trustme
  python-aiohttp
  python-httpx
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('327532687ec663fd98b866be74b2758cae01f087dc9aae1fc28b53971c2145cc1a7cc82b4a71097df6f12f814fa28659dde21d41aea760bc33fe613f2357ca23')
b2sums=('3d1f9e85a149429267eaed96abe2376597edd1e4d07351d42b9e34128520362397453cde380eb799574d78cb3fb0cc6086382b0bfccaded38247a34506d60f5c')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

# Need more AUR dependencies
#check() {
#  cd $_pyname-$pkgver
#  python -m venv --system-site-packages test-env
#  test-env/bin/python -m installer dist/*.whl
#  test-env/bin/python -m pytest
#}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
