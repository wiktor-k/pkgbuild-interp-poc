# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>

_pkgname=poetry-plugin-export
pkgname=python-poetry-plugin-export
pkgver=1.10.1
pkgrel=1
pkgdesc="Poetry plugin to export the dependencies to various formats "
url="https://github.com/python-poetry/poetry-plugin-export"
license=('MIT')
arch=('any')
depends=('python-poetry')
makedepends=('git' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-pytest-mock' 'python-pytest-xdist')
source=("git+https://github.com/python-poetry/poetry-plugin-export.git#tag=$pkgver")
sha512sums=('85caa06348a02aed1f03c290a1e92d82fc4100aab7721f6729ea113d098d377bd129e5dcfb599cfc5a735a1b7e854fe047e3f2d75746403a3e7e1456d7f47114')

prepare() {
  cd poetry-plugin-export
}

build() {
  cd poetry-plugin-export
  python -m build -wn
}

check() {
  cd poetry-plugin-export
  python -m installer -d tmp_install dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  PYTHONPATH="$PWD/tmp_install/$site_packages" pytest
}

package() {
  cd poetry-plugin-export
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
}
