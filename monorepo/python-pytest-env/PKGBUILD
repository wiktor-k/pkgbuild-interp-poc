# Maintainer: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Rafael Fontenelle <rafaelff@gnome.org>

pkgname=python-pytest-env
pkgver=1.8.0
pkgrel=1
pkgdesc='pytest plugin to set environment variables in pytest.ini or pyproject.toml file'
arch=(any)
url='https://github.com/pytest-dev/pytest-env'
license=(MIT)
depends=(
  python
  python-pytest
  python-dotenv
)
makedepends=(
  git
  python-build
  python-installer
  python-hatchling
  python-hatch-vcs
)
checkdepends=(python-pytest-mock)
source=("$pkgname::git+$url#tag=$pkgver")
sha512sums=('91730be4ecd01c46029e4a502f3cb00f0df0d0e8f0783414ded26ffe789834687ac1d1469577cde04ad64c4c95f8502a897a0b404ee0ba0ad9bf152419e71cec')
b2sums=('bc93c28b8bfdc1922039ccc6a517070f1f004d76ca1e492accec1d0c8e0d91948cdd67f8542008af8c7237dd5a35cedf8b2b4115e65298e5a0cc0049173f5c28')

build() {
  cd "$pkgname"

  SETUPTOOLS_SCM_PRETEND_VERSION="$pkgver" python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # temporary install
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$(pwd)/tmp/$site_packages"

  pytest -v \
    --deselect tests/test_env.py::test_env_via_toml["pyproject toml via tool.pytest"] \
    --deselect tests/test_env.py::test_env_via_toml["pytest toml over pyproject toml"] \
    --deselect tests/test_env.py::test_env_via_toml["subdir pytest toml over parent pyproject toml"] \
    --deselect tests/test_env.py::test_env_via_pyproject_toml_bad["pytest.toml"] \
    --deselect tests/test_env.py::test_env_via_pyproject_toml_bad[".pytest.toml"]

}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
