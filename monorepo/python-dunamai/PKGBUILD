# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-dunamai
pkgver=1.26.3
pkgrel=1
pkgdesc='A library for producing dynamic version strings, derived from VCS tags'
arch=('any')
url='https://github.com/mtkennerly/dunamai'
license=('MIT')
depends=(
  'python'
  'python-packaging'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-poetry-core'
)
checkdepends=('python-pytest' 'python-setuptools')
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('c2ce11cf91e538ccbf3676906400d17587ec63fee9dbbe67dd0297b6b6a7f972bba483c12452696fc00a91ed500face0c55e2bdb1441d87aa4f71e734af7be95')
b2sums=('2ed86a839692711ba32782ae6af1b3440f66abd3c267cd46143430cfa97d8305e9b362c8a3a3ea9f023cea85d72a39bc0a181b662c27a6d3bb9a920feffedbcd')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # temporary install so that:
  # * pkg_resources can pick up the package
  # * the required binary can be added to $PATH
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PATH="$(pwd)/tmp/usr/bin:$PATH"
  export PYTHONPATH="$(pwd)/tmp/$site_packages"

  # skip tests that require messing with global git config
  local deselected=(
    tests/integration/test_dunamai.py::test__version__from_git__with_annotated_tags
    tests/integration/test_dunamai.py::test__version__from_git__with_lightweight_tags
    tests/integration/test_dunamai.py::test__version__from_git__with_mixed_tags
    tests/integration/test_dunamai.py::test__version__from_git__with_nonchronological_commits
    tests/integration/test_dunamai.py::test__version__from_git__gitflow
    tests/integration/test_dunamai.py::test__version__from_git__exclude_decoration
    tests/integration/test_dunamai.py::test__version__from_git__broken_ref
    tests/integration/test_dunamai.py::test__version__from_git__trace_env_var
    tests/integration/test_dunamai.py::test__version__from_git__initial_commit_empty_and_tagged
  )

  pytest -v ${deselected[@]/#/--deselect }
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
