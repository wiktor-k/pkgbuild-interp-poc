# Maintainer: Claudia Pellegrino <auerhuhn@archlinux.org>
# Contributor: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Josh Holbrook <josh.holbrook@gmail.com>

pkgname=python-pyee
_name=${pkgname#python-}
pkgver=13.0.1
pkgrel=1
pkgdesc="Port of node.js's EventEmitter to python"
arch=(any)
url=https://github.com/jfhbrook/pyee
license=(MIT)
depends=(python python-typing_extensions)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-pytest-asyncio
  python-pytest-trio
  python-twisted
)
optdepends=(
  python-trio
  python-twisted
)
source=("git+https://github.com/jfhbrook/pyee#tag=v$pkgver")
sha256sums=('2f3cb24b02767665060a63d5639c9b0b8185eea98419a9aa4037b1456d42e0d2')

build() {
  cd "$_name"
  python -m build --wheel --no-isolation
}

check() {
  cd "$_name"
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd "$_name"
  python -m installer --destdir="$pkgdir" dist/*.whl

  install -D -m 644 -t "$pkgdir/usr/share/licenses/$pkgname" \
    LICENSE
}
