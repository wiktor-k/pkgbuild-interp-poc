# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-update-checker
pkgver=1.0.0
pkgrel=1
pkgdesc="A python module that will check for package updates"
arch=('any')
license=('BSD-2-Clause')
url='https://github.com/bboe/update_checker'
depends=('python')
provides=('python-update_checker')
conflicts=('python-update_checker')
replaces=('python-update_checker')
makedepends=('git' 'python-build' 'python-installer' 'python-uv-build')
checkdepends=('python-aiohttp' 'python-pytest' 'python-pytest-asyncio')
optdepends=('python-aiohttp: async API support')
source=("git+https://github.com/bboe/update_checker.git#tag=v$pkgver")
sha512sums=('b384331d9de5a15ecde8232f3e5a4ccb00a6b35fe23046102d4ad9c7a0a46109d5abe28ece0f26cfe6b8258f426b5a493a90a9d4a8d7accbac2c44c8b605754d')

build() {
  cd update_checker
  python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd update_checker
  python -m installer --destdir=test_dir dist/*.whl
  PYTHONPATH="$PWD/test_dir/$_site_packages" pytest
}

package() {
  cd update_checker
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
