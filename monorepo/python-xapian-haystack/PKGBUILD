# Maintainer: David Runge <dvzrv@archlinux.org>

_haystack_name=django-haystack
_django_haystack_pkgver=3.3.0
pkgname=python-xapian-haystack
_name="${pkgname#python-}"
pkgver=4.0.0
pkgrel=1
pkgdesc="Xapian backend for Django-Haystack"
arch=(any)
url="https://github.com/notanumber/xapian-haystack"
license=(GPL-2.0-only)
depends=(
  python
  python-django
  python-django-haystack
  python-filelock
  python-xapian
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=(
  $pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz
  python-$_haystack_name-$_django_haystack_pkgver.tar.gz::https://github.com/django-haystack/django-haystack/archive/refs/tags/v$_django_haystack_pkgver.tar.gz
  python-3.14.patch
)
sha512sums=('8247ed8509ec1792e8c000fd413d2dcde5c867911ea85cc8460fbd739e4aa2b635fd27554d7240d2e346db2b2026c312639fd1397f66042b743771f0e0542c41'
            'f8fb6fe11957391d4dc48f279c07e29b2063c7268a4585c4d2d0462f9b17d3563997a29bed4da5577b4125e70668aa79ce86a553cda71cb94c1af19b9bd808fe'
            'b4129f6604d33d971e3f515c3bc3b601b08cc52a14cfc082f03d4524487863dd4974b0e96267840a25c4ccda414b8bb85a3120e373317bb66d5a6be274d1ecf1')
b2sums=('aab9e5b676577a02a29125551f06717bfdff83430e225c974c7d8df58521eb7b40ae1455391e8d1000aedb97839209d099459270921100e73e0a3a3750646c26'
        'd6193172f5baae90792a636f925f5d9eb6f886f362fcb36f4551dcf0e097fbb2ef58bd7a7bb8dc8401846df3eea4601566e58495f6ab8ce99b4498b8d9b55196'
        'bcdaeacd39f44b413c42e6ab19692ca879f5011cddfe61a5da10d65977ebfec839204c6a1693335328421937564618cae92ed9546224ad37baf3eb8faad0538f')

prepare() {
  # disable tests broken with django 5
  sed -e 's/test_date_facets_month/off1/;s/test_date_facets_seconds/off2/' -i $_name-$pkgver/tests/xapian_tests/tests/test_backend.py

  cd $_name-$pkgver
  patch -p1 -i ../python-3.14.patch # Fix tests with python 3.14
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  # inserting xapian-haystack tests into django-haystack (because... pain)
  ln -sv "$PWD/build/lib/xapian_backend.py" ../$_haystack_name-$_django_haystack_pkgver/haystack/backends/
  cp -rv "$PWD/tests/"* ../$_haystack_name-$_django_haystack_pkgver/test_haystack/
  ln -svf "$PWD/tests/xapian_tests/__init__.py" ../$_haystack_name-$_django_haystack_pkgver/test_haystack/

  cd ../$_haystack_name-$_django_haystack_pkgver
  PYTHONPATH="$PWD:$PYTHONPATH" django-admin test test_haystack.xapian_tests --settings=test_haystack.xapian_settings
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {AUTHORS,README.rst} -t "$pkgdir/usr/share/doc/$pkgname/"
}
