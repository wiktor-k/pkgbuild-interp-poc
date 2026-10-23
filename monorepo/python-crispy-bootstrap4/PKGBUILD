# Maintainer: David Runge <dvzrv@archlinux.org>

_name=crispy-bootstrap4
pkgname=python-crispy-bootstrap4
pkgver=2026.2
pkgrel=1
pkgdesc="Bootstrap4 template pack for django-crispy-forms"
arch=(any)
url="https://github.com/django-crispy-forms/crispy-bootstrap4"
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
sha512sums=('6a93c028337702ae921097b4018ebd1db08f0ed13dab4826910670595a17f833bb013f99fc2baffe597cb4cb9cb72faf6b88b12d6e4047e9a089c36118717374')
b2sums=('730b94aab236ca7b7a508ff7f8329f15114a4b336bdfb0e3acdc66600882466b29c0559615e96a6379808269df1bcfbd952a85f0c2235aec98be292166b10c3d')

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
