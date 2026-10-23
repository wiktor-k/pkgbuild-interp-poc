# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: minus <minus@mnus.de>

_pkgname=gevent-websocket
pkgname=python-${_pkgname}
pkgver=0.10.1
pkgrel=13
pkgdesc='WebSocket library for the gevent networking library'
url='https://www.gitlab.com/noppo/gevent-websocket'
arch=('any')
license=('Apache-2.0')
depends=('python' 'python-gevent' 'gunicorn' 'python-ujson')
makedepends=('cython' 'python-build' 'python-installer' 'python-setuptools' 'python-sphinx' 'python-wheel')
source=(https://gitlab.com/noppo/gevent-websocket/-/archive/${pkgver}/gevent-websocket-${pkgver}.tar.gz
        docs-2to3.patch)
sha512sums=('33473bd20bf8bdab95784bc46093c4f3567d61290731e1980de3b8bf96d71976dca658547320329d5ee5291f54e423a1c7ae8300358bc641eae3c71a1943ea84'
            '42dbcc66a091d3da2e0d42055c69d01e5fade27fb568959c0b4b8f4232dafc40823c09565e2d3fe6b1758e814356e3d81840710a7bd5f89bf02e8b1d05140b22')

prepare() {
  cd ${_pkgname}-${pkgver}
  patch -Np1 -i ${srcdir}/docs-2to3.patch
}

build() {
  cd ${_pkgname}-${pkgver}
  python -m build --wheel --no-isolation
  sphinx-build -b text docs docs/_build/text
  sphinx-build -b man docs docs/_build/man
}

package() {
  cd ${_pkgname}-${pkgver}
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 README.rst -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 docs/_build/text/*.txt -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -Dm 644 docs/_build/man/gevent-websocket.1 "${pkgdir}/usr/share/man/man1/${pkgname}.1"
}

# vim: ts=2 sw=2 et:
