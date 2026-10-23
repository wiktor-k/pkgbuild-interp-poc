# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-pytest-twisted
pkgver=1.14.3
pkgrel=4
pkgdesc='A twisted plugin for py.test'
arch=('any')
license=('BSD-3-Clause')
url='https://github.com/pytest-dev/pytest-twisted'
depends=('python-pytest' 'python-greenlet' 'python-decorator' 'python-twisted')
makedepends=('git' 'python-build' 'python-installer' 'python-wheel' 'python-setuptools')
source=("git+https://github.com/pytest-dev/pytest-twisted.git#tag=v$pkgver")
sha512sums=('afbc985721bb7d7b063c1181c0d295d930c0f6eb2a66434eee4980755c77657456bd9817bc42e29230341eb51cb59c0bd54f39484d4e95ff810cb3eaf5d09a61')

build() {
  cd pytest-twisted
  python -m build --wheel --no-isolation
}

check() {
  cd pytest-twisted
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd pytest-twisted
  python -m installer --destdir="$pkgdir" dist/*.whl
}
