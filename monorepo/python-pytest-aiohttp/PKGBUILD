# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

pkgname=python-pytest-aiohttp
_name="${pkgname#python-}"
pkgver=1.1.1
pkgrel=1
pkgdesc='pytest plugin for aiohttp support'
url='https://github.com/aio-libs/pytest-aiohttp'
arch=('any')
license=('Apache-2.0')
depends=('python' 'python-pytest' 'python-aiohttp')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-setuptools-scm' 'python-wheel')
source=(${pkgname}-${pkgver}.tar.gz::$url/archive/v${pkgver}.tar.gz)
sha256sums=('cc4f715f0088351e609825f5a0d36ebcfc140e3c8660067df1f18e31ffe02fb6')
sha512sums=('5afb49583d2f0e30569b5fd1e04a2cbfd45e32481adff2fd3c8de82852ef0e5a1b96130ecf6c904fe35b70e8c1699f20536da752c1126f0bc53e3d553dd330ba')

build() {
  cd ${_name}-${pkgver}
  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --no-isolation
}

package() {
  cd ${_name}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 README.rst CHANGES.rst -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
