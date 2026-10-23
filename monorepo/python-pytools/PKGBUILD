# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Stéphane Gaudreault <stephane@archlinux.org>

pkgname=python-pytools
pkgver=2026.1.1
pkgrel=1
pkgdesc="A collection of tools for Python"
arch=(any)
url="https://github.com/inducer/pytools"
license=(MIT)
depends=(
  python
  python-platformdirs
  python-siphash24
  python-typing_extensions
)
optdepends=('python-numpy: NumPy-related functionality')
makedepends=(
  git
  python-build
  python-hatchling
  python-installer
)
checkdepends=(
  python-attrs
  python-numpy
  python-pytest
)
source=("git+https://github.com/inducer/pytools.git#tag=v$pkgver")
sha512sums=('5e227aadffc2794c7f2f2ea580f2ee36b04595fc34140a2ee22c7bc699e37f1ddd7823fd35b6bc994a1063ceaada6180802f02c2c98550e2a55b44279900b0aa')

build() {
  cd pytools
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    -W ignore::DeprecationWarning
  )

  cd pytools
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest "${pytest_options[@]}"
}

package(){
  cd pytools
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
