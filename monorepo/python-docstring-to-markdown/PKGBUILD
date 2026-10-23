# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Hugo Osvaldo Barrera <hugo@barrera.io>

pkgname=python-docstring-to-markdown
_name=${pkgname#python-}
pkgver=0.17
pkgrel=2
pkgdesc='On the fly conversion of Python docstrings to markdown'
arch=('any')
url=https://github.com/python-lsp/docstring-to-markdown
license=('LGPL-2.1-or-later')
depends=('python'
         'python-importlib-metadata'
         'python-typing_extensions')
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=('python-pytest')
source=("git+$url.git#tag=v$pkgver")
b2sums=('1aeb125b7fa61a2efba41313c5d8d102079bb6fc76692f31112d1456678e116e57c8ca9fe35a477dad780e0116dbcc318a0e1330d6874a9e782d326f96299982')

build() {
  cd "$_name"
  python -m build --wheel --no-isolation
}

check() {
  cd "$_name"
  # Override addopts as they invoke coverage testing
  PYTHONPATH="$PWD/$_name:$PYTHONPATH" pytest --override-ini="addopts="
}

package() {
  cd "$_name"
  python -m installer --destdir="$pkgdir" dist/*.whl
}
