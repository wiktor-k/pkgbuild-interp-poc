# Maintainer: Balló György <ballogyor+arch at gmail dot com>

pkgname=python-gaphas
pkgver=5.1.2
pkgrel=1
pkgdesc='Diagramming widget library for Python'
arch=(any)
url='https://github.com/gaphor/gaphas'
license=(Apache-2.0)
depends=(
  glib2
  graphene
  gtk4
  python
  python-cairo
  python-gobject
)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
  python-sphinx-furo
  python-wheel
)
checkdepends=(
  python-pytest-asyncio
  xorg-server-xvfb
)
source=("git+https://github.com/gaphor/gaphas.git#tag=$pkgver")
b2sums=(0c9e536a1016189b525ec1ae10806ab76602f7bb891c5193a57710929cd5a809ca0bf651574b1999337e784e0156d4d80f0e94c5cf8c969efb9b0d9c665786cd)

build() {
  cd gaphas
  python -m build --wheel --no-isolation
}

check() {
  cd gaphas
  PYTHONPATH="$PWD/dist/gaphas-$pkgver-py3-none-any.whl" \
    dbus-run-session xvfb-run \
    -s '-screen 0 1920x1080x24 -nolisten local' \
    pytest --ignore=tests/test_architecture.py
}

package() {
  cd gaphas
  python -m installer --destdir="$pkgdir" dist/*.whl
  python -m sphinx -b html docs "$pkgdir/usr/share/doc/$pkgname/html"
  rm -r "$pkgdir/usr/share/doc/$pkgname/html/.doctrees"
}
