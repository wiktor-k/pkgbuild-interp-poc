# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>

pkgname=python-cron-converter
pkgver=2.0.2
pkgrel=1
pkgdesc="Cron string converter for Python"
arch=(any)
url="https://github.com/Sonic0/cron-converter"
license=(MIT)
depends=(
  python
  python-dateutil
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=("$pkgname::git+$url.git#tag=v$pkgver")
sha512sums=('ac07e7312a277f8760add9f818d64b927fd3a97ef7bf96191ce147150dc03eefac4e426e97a7624fed9dec5e4e563c4da08e29501571ee69ee1cba8f614ad937')
b2sums=('4f87efa5ca13c05d06f8aee248a75dc19b4d94d34770a419c4d2c59aa2b2f45da1ada96004eb39f97bf2182d28492370638be3fed6d92763e4704e4a16e358dc')

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname
  python -m unittest discover -v tests/unit
  python -m unittest discover -v tests/integration
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
