# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>

_name=pydantic-settings
pkgname=python-pydantic-settings
pkgver=2.16.1
pkgrel=1
pkgdesc="Settings management using pydantic"
arch=(any)
url="https://github.com/pydantic/pydantic-settings"
license=(MIT)
depends=(
  python
  python-dotenv
  python-pydantic
  python-pydantic-core  # directly used, but transitive dependency of python-pydantic
  python-typing_extensions
)
makedepends=(
  python-build
  python-hatchling
  python-installer
)
checkdepends=(
  python-pytest
  python-pytest-examples
  python-pytest-mock
  # TODO: add python-azure-keyvault-secrets and python-azure-identity for azure-key-vault optdepend
)
optdepends=(
  'python-pyyaml: for YAML support'
)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('13c0e04e6302d25975d3db158ea1ada74017689cee04ccadb09fe74ede42010e5b726a54e13cb96e1745e9ed70cfca444b737eaace0ebdbf87c0bb8fa98508e4')
b2sums=('88cfbb42d7cec585a3a44c790cd51722aba4e9bc97742936770d2f0145f070d14327c87092eab5d1eb003c5c64ed0601f148aec6b9f6e6812fc45115e585fdee')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    # we are not interested in linting and formatting
    --ignore tests/test_docs.py
    # we don't yet have azure integration in the repos
    --ignore tests/test_source_azure_key_vault.py
    # we don't yet have gcp integration in the repos
    --ignore tests/test_source_gcp_secret_manager.py
  )
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH"
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
