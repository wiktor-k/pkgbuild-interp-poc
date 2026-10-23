# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-id
_name="${pkgname#python-}"
pkgver=1.6.1
pkgrel=1
pkgdesc="A tool for generating OIDC identities"
arch=(any)
url="https://github.com/di/id"
license=(Apache-2.0)
depends=(
  python
  python-urllib3
)
makedepends=(
  python-build
  python-flit-core
  python-installer
)
checkdepends=(
  python-pytest
  python-pretend
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('58ce4586eff00ad2ac36b4dff31199eedca7287a5a317175238b8ed0a015b8fd5e1b09ad66a5b2ad7042cefac6d8dd980ee1924d112afe6a62a56edf6f07ae24')
b2sums=('ddea443e5924c6925789a900ea837e0764b795cbea838a05518a4d07718de861e2a66d1f95a8053afd118cc73b4a57fedabb9092292912a0f1dad937fa4e18be')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  PYTHONPATH="." pytest -vv
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGELOG,README}.md -t "$pkgdir/usr/share/doc/$pkgname/"
}
