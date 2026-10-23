# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-pytest-ruff
pkgver=0.5
pkgrel=1
pkgdesc="Pytest plugin to check ruff requirements"
url="https://github.com/businho/pytest-ruff"
license=('MIT')
arch=('any')
depends=('python' 'python-pytest' 'python-ruff')
makedepends=('git' 'python-build' 'python-installer' 'python-poetry-core'
             'python-poetry-dynamic-versioning')
checkdepends=('python-pytest-mock')
source=("git+https://github.com/businho/pytest-ruff.git#tag=v$pkgver")
sha512sums=('982894070cf3ac5ff88c44983ecc293b5399965029e424414c8b0bf637e7c8a02d6c8d134bc620a8eea86c557276faaf00122a87f4330dba174164b06c0cfaf2')

build() {
  cd pytest-ruff
  python -m build -nw
}

check() {
  cd pytest-ruff
  python -m installer -d tmp_install dist/*.whl
  local python_version=$(python -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')
  PYTHONPATH="$PWD/tmp_install/usr/lib/python${python_version}/site-packages" \
    pytest
}

package() {
  cd pytest-ruff
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
