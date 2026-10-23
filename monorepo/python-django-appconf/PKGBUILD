# Maintainer: David Runge <dvzrv@archlinux.org>

_name=django-appconf
pkgname=python-django-appconf
pkgver=1.2.0
pkgrel=2
pkgdesc="An app to handle configuration defaults of packaged Django apps gracefully"
arch=(any)
url="https://github.com/django-compressor/django-appconf"
license=(BSD-3-Clause)
depends=(
  python
  python-django
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=($_name-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('e4385dad1ef870ebe124c5e1b192d942bf23bfb53b849929124f169439cda0cf9709d31d1c6fddf7149a3f572ee13f57af1206d290eb66f4c0af1a86494b9ec9')
b2sums=('fd23f9ce9cd3dc3b312ba0186d15afc6738ae4de804c3c5093ffc0c6412c386515bf87587aa014e7a9d9f99ea8cf488cff663c5e3093a8eaafd02316fe28942a')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  export PYTHONPATH="$PWD:build:$PYTHONPATH"
  DJANGO_SETTINGS_MODULE=tests.test_settings django-admin test -v2 tests
  rm -rf tests/__pycache__
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {AUTHORS,README.rst} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
