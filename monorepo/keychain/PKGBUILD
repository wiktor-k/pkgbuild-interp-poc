# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=keychain
pkgver=3.0.8
pkgrel=2
pkgdesc='A front-end to ssh-agent, allowing one long-running ssh-agent process per system, rather than per login'
arch=(any)
url='https://www.funtoo.org/Keychain'
license=(GPL-2.0-only)
depends=(python)
makedepends=(git
             python-build
             python-installer
             python-setuptools)
source=("git+https://github.com/danielrobbins/keychain#tag=$pkgver")
sha256sums=('533c1e540a1ecbd5498972abbb336de332add3cfcf0f3decbf9a000137c1816c')

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  make DESTDIR="$pkgdir" PREFIX=/usr install-man
}
