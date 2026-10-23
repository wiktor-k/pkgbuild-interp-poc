# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-configupdater
pkgver=4.0
pkgrel=1
pkgdesc="Parser like ConfigParser but for updating configuration files"
url="https://github.com/pyscaffold/configupdater"
license=('MIT')
arch=('any')
depends=('python')
makedepends=('git' 'python-hatchling' 'python-hatch-vcs' 'python-build' 'python-installer')
checkdepends=('python-pytest')
source=("git+https://github.com/pyscaffold/configupdater.git#tag=v$pkgver")
sha512sums=('227a11a3ca6edefa4a939073c50377096ef4e904b0457ee64716f5418b59552f283b8685a977009155cda221276e026f0387a405c542db0517e614222441ceaa')

build() {
  cd configupdater
  python -m build --wheel --no-isolation
}

check() {
  cd configupdater
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd configupdater
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE.txt -t "$pkgdir"/usr/share/licenses/$pkgname/
}
