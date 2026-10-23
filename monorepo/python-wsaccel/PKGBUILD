# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-wsaccel
pkgver=0.6.7
pkgrel=1
pkgdesc='Accelerator for ws4py and AutobahnPython'
arch=('x86_64')
url='https://github.com/methane/wsaccel'
license=('Apache')
depends=('python')
makedepends=('python-setuptools' 'cython' 'python-build' 'python-installer' 'python-wheel' 'git')
checkdepends=('python-pytest')
source=("git+https://github.com/methane/wsaccel.git#tag=v$pkgver")
sha512sums=('4506cb455c5cb1cf23b4aab101f2ada5a44c65fdbf81a25dcb7193e1e829c630920cd3af70798848e0a2c94b8de3c0e695ab06515735355be9db65c63c8d8d86')

prepare() {
  cd wsaccel
  sed -i 's/Cython~=3.0.11/Cython/' pyproject.toml
}

build() {
  cd wsaccel
  python -m build --wheel --no-isolation
}

check() {
  cd wsaccel
  local python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
  PYTHONPATH="$PWD/build/lib.linux-$CARCH-cpython-$python_version" pytest
}

package() {
  cd wsaccel
  python -m installer --destdir="$pkgdir" dist/*.whl
}
