# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-deprecated
pkgver=3.0.0
pkgrel=1
pkgdesc="Python @deprecated decorator to deprecate old python classes, functions or methods"
url="https://github.com/tantale/deprecated"
license=('MIT')
arch=('any')
depends=('python' 'python-wrapt')
makedepends=('git' 'python-build' 'python-hatchling' 'python-installer')
checkdepends=('python-packaging' 'python-pytest')
source=("git+https://github.com/tantale/deprecated.git#tag=v$pkgver")
sha512sums=('e04190f497bd682c0475bf0d78c52fddbc77e7c6252faf5b6b6c22e3f54b430ca1bc917f4dda166ce17f9f31748086245927c6a56954799e758ffb343ee81c7f')

build() {
  cd deprecated
  python -m build --wheel --no-isolation
}

check() {
  cd deprecated
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd deprecated
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE.md -t "$pkgdir"/usr/share/licenses/$pkgname/
}
