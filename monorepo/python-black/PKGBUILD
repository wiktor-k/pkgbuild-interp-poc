# Maintainer: Maxim Baz <archlinux at maximbaz dot com>
# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: James Zhu <jameszhu@berkeley.edu>

pkgname=python-black
pkgver=26.10.1
pkgrel=1
pkgdesc='Uncompromising Python code formatter'
arch=(any)
url='https://github.com/psf/black'
license=(MIT)
depends=(
  python
  python-click
  python-mypy_extensions
  python-packaging
  python-pathspec
  python-platformdirs
  python-pytokens
)
makedepends=(
  git
  python-build
  python-hatch-fancy-pypi-readme
  python-hatch-vcs
  python-hatchling
  python-installer
)
checkdepends=(
  ipython
  python-aiohttp
  python-parameterized
  python-pytest
  python-tokenize-rt
)
optdepends=(
  'ipython: for Jupyter notebook support'
  'python-tokenize-rt: for Jupyter notebook support'
  'python-aiohttp: for the blackd HTTP server'
  'python-colorama: for colored diffs'
)
source=("$pkgname::git+$url#tag=$pkgver")
sha512sums=('46cbeb33168d5292a9bedbf9adcbf3e02a15857f00007480c5c4789b378eaf128291221829f6a4b82c80b180b7035ef02210e9ad4b39b4ded710b715959ed125')
b2sums=('451841e95b5b984996793ad503b2b8fa6d7250bac5859c1bdff79a648a26623cfe9bf756027c764f6d4ce99a4b5c4d88ed09ef3596c29a1c6618dd2f3fff5ee8')

build() {
  cd "$pkgname"

  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd "$pkgname"

  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # https://github.com/psf/black/issues/3251#issuecomment-1236413890
  ulimit -n 2048
  PATH="$PWD/test-env/bin:$PATH" test-env/bin/python -m pytest
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE

  # vim plugin
  install -Dm644 -t "$pkgdir/usr/share/vim/vimfiles/plugin" plugin/black.vim
  install -Dm644 -t "$pkgdir/usr/share/vim/vimfiles/autoload" autoload/black.vim
}

# vim:set ts=2 sw=2 et:
