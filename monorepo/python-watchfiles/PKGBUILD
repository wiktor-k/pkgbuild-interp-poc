# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: @RubenKelevra <cyrond@gmail.com>

pkgname=python-watchfiles
pkgver=1.3.0
pkgrel=1
pkgdesc='Simple, modern and high performance file watching and code reload in Python'
arch=(x86_64)
url='https://github.com/samuelcolvin/watchfiles'
license=(MIT)
depends=(
  glibc
  libgcc
  python
  python-anyio
)
makedepends=(
  git
  python-build
  python-installer
  python-maturin
)
checkdepends=(
  python-dirty-equals
  python-pytest
  python-pytest-mock
  python-pytest-timeout
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('31f089f939be16a8ace3321134ee4f73e24545529f2dea89f05d43d6190f011594597f36e211391659ca27bf893e08c94c47da2920d48db8f57d0ac2b735e041')

prepare() {
  cd ${pkgname#python-}
  # Fix test collection with pytest >= 9.1
  git cherry-pick -n fe87d8a3e1cceb7979ca97e94f049ffcfa57c3bf

  # This prevents tests from detecting the watchfiles module.
  rm -v tests/__init__.py
}

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # Don't add CWD to PYTHONPATH, the watchfiles package in CWD will take
  # precedence over installed one.
  test-env/bin/python -Im pytest
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
