# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-svglib
_name="${pkgname#python-}"
pkgver=2.3.0
pkgrel=1
pkgdesc="Read SVG files and convert them to other formats"
arch=(any)
url="https://github.com/deeplook/svglib"
license=(LGPL-3.0-only)
depends=(
  python
  python-lxml
  python-pillow
  python-reportlab
  python-tinycss2
  python-cssselect2
)
makedepends=(
  python-build
  python-hatchling
  python-installer
)
checkdepends=(
  python-pytest
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('067daf7c6e7d94d778094a657b214f1e647d01232f317f7e88d05adec50bd7c2610765ea242ee0617e7e8585b182f4ff25f6f9d1a9c2dcfeeb72d789c0ffe185')
b2sums=('7c0daa119f8f17e96a6a2547555cab74b3a104a1051c4b9838f6a91a5c73cf2fc31f1219f6eea724c5706c5cebf9c3cd736f2b3c38fdc0b1a8ee41a7254e1ab9')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    # disable tests that need remote resources
    --ignore tests/test_samples.py
  )

  cd $_name-$pkgver
  export PYTHONPATH="src:$PYTHONPATH"
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGELOG,CONTRIBUTORS,README}.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
