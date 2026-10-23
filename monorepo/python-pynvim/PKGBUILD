# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
pkgname=python-pynvim
pkgver=0.6.0
pkgrel=2
pkgdesc='Python client for Neovim'
url='https://github.com/neovim/pynvim'
arch=('any')
license=('Apache')
depends=('python' 'python-msgpack' 'python-greenlet')
provides=('python-neovim')
makedepends=('python-setuptools' 'python-build' 'python-wheel' 'python-installer')
replaces=('python-neovim')
source=("$pkgname-$pkgver.tar.gz::https://github.com/neovim/pynvim/archive/${pkgver}.tar.gz")
sha256sums=('6649d630ce4a94007fe69a8debb010c7b185057f36d61576d2542c90f3c8db99')

prepare() {
  sed -i "s/msgpack-python/msgpack/g" pynvim-${pkgver}/setup.py
}

build() {
  cd "${srcdir}/pynvim-${pkgver}"
  python -m build --wheel --no-isolation
}

package() {
  cd "${srcdir}/pynvim-${pkgver}"
  python -m installer --destdir="$pkgdir" dist/*.whl
}

# vim:set sw=2 sts=2 et:
