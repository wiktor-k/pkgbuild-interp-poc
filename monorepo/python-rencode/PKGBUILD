# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Simon Sapin <simon dot sapin at exyr dot org>
# Contributor: Kyle Keen <keenerd@gmail.com>

pkgname=python-rencode
pkgver=1.0.9
pkgrel=1
pkgdesc="A Module similar to bencode from the BitTorrent project"
url="https://github.com/aresch/rencode"
license=('GPL-3.0-or-later')
arch=('x86_64')
depends=('glibc' 'python')
makedepends=('git' 'cython' 'python-build' 'python-installer' 'python-setuptools')
checkdepends=('python-pytest')
source=("git+https://github.com/aresch/rencode.git#tag=v$pkgver")
sha512sums=('29e6b14c683a17dd2b08be5a0705025f5f8c70b63290f11e60d2959cedefdb21488b3cf0580db3d19af1e553cbc3c64ca6df84cb8e7c7dbe56c31c575317e2a0')

build() {
  cd rencode
  python -m build --wheel --no-isolation
}

check() {
  python -m venv --system-site-packages testenv
  testenv/bin/python -m installer rencode/dist/*.whl
  testenv/bin/python -m pytest rencode/tests
}

package() {
  cd rencode
  python -m installer --destdir="$pkgdir" dist/*.whl
}
