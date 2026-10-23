# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Ilya Elenskiy <elenskiy.ilya@gmail.com>
# Contributor: Julien Nicoulaud <julien.nicoulaud@gmail.com>

pkgname=python-reactivex
pkgver=5.1.0
pkgrel=1
pkgdesc='Reactive Extensions for Python'
arch=(any)
url=https://reactivex.io
license=(MIT)
depends=(
  python
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-hatchling
  python-installer
)
checkdepends=(
  python-pytest
  python-pytest-asyncio
)
source=(git+https://github.com/ReactiveX/RxPY.git#tag=v${pkgver})
b2sums=('d41379725b5375f5298c93f38cac87d0112890f401057a317cf0a9a32cc8a1c2429fdd3781f7501a1def4233bc253173ff2b92943acda2dffb32796cdcfdf278')

prepare() {
  cd "${srcdir}/RxPY"
}

build() {
  cd RxPY
  python -m build --wheel --no-isolation
}

check() {
  cd RxPY
  python -m pytest
}

package() {
  python -m installer --destdir="${pkgdir}" RxPY/dist/*.whl
  install -vDm 644 RxPY/LICENSE -t "${pkgdir}"/usr/share/licenses/${pkgname}/
}

# vim: ts=2 sw=2 et:
