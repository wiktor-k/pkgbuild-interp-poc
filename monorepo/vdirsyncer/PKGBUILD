# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Hugo Osvaldo Barrera <hugo@barrera.io>

pkgname=vdirsyncer
pkgver=0.21.0
pkgrel=1
pkgdesc='Synchronize CalDAV and CardDAV'
arch=(any)
url=https://vdirsyncer.pimutils.org/en/stable/
license=(BSD-3-Clause)
depends=(
  python
  python-aiohttp
  python-aiostream
  python-click
  python-click-log
  python-requests
  python-tenacity
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools-scm
  python-sphinx
  python-wheel
)
checkdepends=(
  python-aioresponses
  python-hypothesis
  python-pytest
  python-pytest-asyncio
  python-pytest-httpserver
  python-pytest-subtesthack
  python-trustme
)
optdepends=('python-aiohttp-oauthlib: Google support')
source=("git+https://github.com/pimutils/vdirsyncer.git#tag=v$pkgver")
b2sums=('731897ddde72b9069b3d7508cb05e64258c39436b7d730f58d161e9f7d2454465a799c8962628d35f187bc79fb58c6d8c579fe5d357205382940c3b76bc4dd5f')
#validpgpkeys=('951082781CA308E4A529DEFD6F9FFF122B98C0CD') # Hugo Osvaldo Barrera <hugo@barrera.io>

build() {
  cd $pkgname
  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver
  python -m build --wheel --skip-dependency-check --no-isolation
  export PYTHONPATH="build:$PYTHONPATH"
  sphinx-build -b man docs/ build/
}

check() {
  cd $pkgname
  export DETERMINISTIC_TESTS=true
  # Override addopts as they invoke coverage testing
  pytest --override-ini="addopts="
}

package() {
  # symlink license
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  install -d "$pkgdir"/usr/share/licenses/$pkgname
  ln -s "$site_packages"/$pkgname-$pkgver.dist-info/licenses/LICENSE \
    "$pkgdir"/usr/share/licenses/$pkgname/LICENSE

  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  # systemd
  install -vDm 644 contrib/$pkgname.{service,timer} -t "$pkgdir"/usr/lib/systemd/user
  # man page
  install -vDm 644 build/$pkgname.1 -t "$pkgdir"/usr/share/man/man1
  # docs
  install -vDm 644 {AUTHORS,CHANGELOG,README}.rst config.example -t "$pkgdir"/usr/share/doc/$pkgname
}
