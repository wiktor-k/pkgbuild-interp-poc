# Maintainer: Chih-Hsuan Yen <yan12125@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-anyio
pkgver=4.15.1
pkgrel=2
pkgdesc='High level compatibility layer for multiple asynchronous event loop implementations'
arch=(any)
url='https://github.com/agronholm/anyio'
license=(MIT)
depends=(
  python
  python-idna
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-trio
  python-uvloop
  python-wheel
)
checkdepends=(
  python-hypothesis
  python-psutil
  python-pytest
  python-pytest-mock
  python-pytest-timeout
  python-trustme
  python-truststore
)
optdepends=(
  'python-outcome: trio backend'
  'python-pytest: pytest plugin'
  'python-trio: trio backend'
  'python-uvloop: use uvloop for asyncio backend'
)
source=("git+$url.git#tag=$pkgver")
b2sums=('98fd369a43a469da4c62b2e9b475d944dfc5e885b19c13d10832243c5c798c986bc172de4bfdaccee0003238dfe509c79ff9274b4aa4891003eb61271ad23d79')

prepare() {
  cd ${pkgname#python-}
  # pytest 9 validates native TOML types for pytest-timeout settings.
  sed -i 's/timeout = "20"/timeout = 20.0/' pyproject.toml
  # 3 seconds is too short for nested pytest subprocesses on riscv64 builders.
  sed -i 's/timeout=3)/timeout=30)/g' tests/test_pytest_plugin.py
}

build() {
  cd ${pkgname#python-}
  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
