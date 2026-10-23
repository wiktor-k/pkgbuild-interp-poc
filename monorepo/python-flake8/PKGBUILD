# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Karol "Kenji Takahashi" Woźniak <kenji.sx>
# Contributor: spider-mario <spidermario@free.fr>
# Contributor: Allen Li <darkfeline at abagofapples.com>

pkgname=python-flake8
pkgver=7.4.1
pkgrel=1
epoch=1
pkgdesc="The modular source code checker: pep8, pyflakes and co"
arch=('any')
url='https://flake8.pycqa.org'
license=('MIT')
depends=(
  'python'
  'python-pyflakes'
  'python-mccabe'
  'python-pycodestyle'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=('python-pytest')
provides=('flake8')
replaces=('flake8')
source=("$pkgname::git+https://github.com/PyCQA/flake8#tag=$pkgver")
sha512sums=('0ab4ef9948ed37f9878ca2f0c73ccb4222a83756226b71072090086dc4fb8c5bb50a27016becd398950a04bbc94cdbbba4739a2bb7f84bd51e62b92d93a2cb54')
b2sums=('d521ca4a752ca15e8f27e8589be100235a44bc0073ed0285f3c5bb0b9410d6794f52ba96732ff2c82680964f8ff261f1344ab440056178226a5652366f99f9e3')

prepare() {
  cd "$pkgname"

  sed \
    -e 's/,<[0-9=.]*//' \
    -i setup.cfg

  sed \
    -e '/error/a \    ignore:Creating a LegacyVersion has been deprecated and will be removed in the next major release:DeprecationWarning' \
    -i pytest.ini
}

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # hack entry points by installing it to a temporary directory
  python -m installer --destdir="$PWD/tmp_install" dist/*.whl
  local python_version=$(python -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')
  export PYTHONPATH="$PWD/tmp_install/usr/lib/python${python_version}/site-packages:$PYTHONPATH"

  local deselected=(
    tests/unit/test_pyflakes_codes.py::test_all_pyflakes_messages_have_flake8_codes_assigned
    tests/integration/test_plugins.py::test_local_plugin_can_add_option
  )

  pytest ${deselected[@]/#/--deselect }

}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}

# vim:set ts=2 sw=2 et:
