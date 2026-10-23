# Maintainer: Thore Bödecker <foxxx0@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Ian Beringer <ian@ianberinger.com>

pkgname=python-kubernetes
pkgver=37.0.1
pkgrel=1
license=('Apache-2.0')
pkgdesc='Python client for the kubernetes API'
arch=('any')
url='https://github.com/kubernetes-client/python'
depends=(
  'python'
  'python-adal'
  'python-aiohttp'
  'python-aiohttp-retry'
  'python-certifi'
  'python-dateutil'
  'python-durationpy'
  'python-google-auth'
  'python-requests'
  'python-requests-oauthlib'
  'python-six'
  'python-urllib3'
  'python-websocket-client'
  'python-yaml'
  'python-pydantic'
  'python-typing_extensions'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'python-isort'
  'python-pluggy'
  'python-py'
  'python-pytest'
  'python-pytest-cov'
  'python-coverage'
  'python-nose'
  'python-sphinx'
  'python-pycodestyle'
  'autopep8'
)
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('89854db20a0a5e19ecbc208eed673aedc46be47d2d1814caec161d909266c35a4dff8d3e6ef91d27c19a6dd40b93582744333087444e0d4ec0e8e3944321c38d')

build() {
  cd python-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd python-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest
}

package() {
  cd python-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
