# Maintainer: David Runge <dvzrv@archlinux.org>

_name=cookiecutter
pkgname=python-cookiecutter
pkgver=2.7.1
pkgrel=1
pkgdesc="A command-line utility that creates projects from project templates"
arch=(any)
url="https://github.com/cookiecutter/cookiecutter"
license=(BSD-3-Clause)
depends=(
  python
  python-arrow
  python-binaryornot
  python-click
  python-jinja
  python-pyyaml
  python-requests
  python-rich
  python-slugify
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-freezegun
  python-pytest
  python-pytest-mock
)
source=($url/archive/v$pkgver/$_name-$pkgver.tar.gz)
sha512sums=('b24a08a41b25b37f30771b28398743cb164eedd34385ccbd309ace7bb4946f6e981f1df4e79a21f990e659765207f00dffc86262db087f434cfcab46d7672de3')
b2sums=('e7f548f6c9645f00b8843fe38504eccb6bcd8a71246f6a40e9ee3becca66d2f59361ece790896d085dd3d40eabc5878477aaed184a3415028d4bf20f759e2678')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    -o addopts=''  # we do not care about coverage
    # ignore tests that download the internet with git
    --deselect tests/vcs/test_clone.py::test_clone_handles_repo_typo
    --deselect tests/vcs/test_clone.py::test_clone_handles_branch_typo
    --deselect tests/vcs/test_clone.py::test_clone_unknown_subprocess_error
  )

  cd $_name-$pkgver
  export PYTHONPATH="build:$PYTHONPATH"
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
