# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Daniel M. Capella <polyzen@archlinux.org>

pkgname=python-sphinx-autoapi
_name=${pkgname#python-}
pkgver=3.8.1
pkgrel=1
pkgdesc='Sphinx API documentation generator'
arch=(any)
url=https://github.com/readthedocs/sphinx-autoapi
license=(MIT)
depends=(
  python
  python-astroid
  python-docutils  # used directly, pulled in via python-sphinx
  python-jinja
  python-sphinx
  python-pyyaml
)
makedepends=(
  git
  python-build
  python-flit-core
  python-installer
)
checkdepends=(
  python-beautifulsoup4
  python-pytest
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('c0815844ca4884d25e6543f378ce01648358427d8a976de7c6738ad0452b04ee1ce56038dca9352fae17f9b03103849416a3e5c43ade45531f803b6b465cce9d')

build() {
  cd "$_name"
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd "$_name"
  PYTHONPATH="src" pytest -v
}

package() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  install -d "$pkgdir"/usr/share/licenses/$pkgname
  ln -s "$site_packages"/"${_name//-/_}"-$pkgver.dist-info/licenses/LICENSE.rst \
    "$pkgdir"/usr/share/licenses/$pkgname

  cd "$_name"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGELOG,README}.rst -t "$pkgdir"/usr/share/doc/$pkgname
}
