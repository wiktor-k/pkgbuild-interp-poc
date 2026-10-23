# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-crispy-bootstrap5
_name="${pkgname#python-}"
pkgver=2026.9
pkgrel=1
pkgdesc="Bootstrap5 template pack for django-crispy-forms"
arch=(any)
url="https://github.com/django-crispy-forms/crispy-bootstrap5"
license=(MIT)
depends=(
  python
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-django
  python-django-crispy-forms
  python-pytest
  python-pytest-django
)
source=($url/archive/$pkgver/$_name-$pkgver.tar.gz)
sha512sums=('7674ea506821b2bce4319ac7709686159c5a694de9e4b5a6e20f03e3368d09646f9b227f0b11e995bee2aa019dc8ebb47a6a718363f35479aa3488bec99552b9')
b2sums=('2145f99cea1ff450669a67ba19a26e2b2648ec93ba40c8f93bdcb2d0970e91434c3ece65b391d417828be05333edbc4bb885ca00b81f665eecc1ce8ced574f0d')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name-$pkgver
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="test_dir/$site_packages:$PWD:$PYTHONPATH"
  pytest -vv tests
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGELOG,README}.md -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
