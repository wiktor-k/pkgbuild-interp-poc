# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-msgspec
pkgver=0.22.0
pkgrel=1
pkgdesc='A fast serialization and validation library, with builtin support for JSON, MessagePack, YAML, and TOML'
arch=(x86_64)
url=https://github.com/jcrist/msgspec
license=(BSD-3-Clause)
depends=(
  glibc
  python
  python-attrs
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-msgpack
  python-pytest
)
optdepends=(
  'python-tomli-w: for TOML writing support'
  'python-yaml: for YAML support'
)
source=("$pkgname::git+$url.git#tag=$pkgver")
b2sums=('c4f4dfdcde33ad7890a2483a58ea76c30d1583bc3e4a3675f6e66d4298ebb402f924cb2c0aa4534ab3858dccc21948c5890097a7102aa63d8a5def8cba781701')

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest tests/unit
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
}
