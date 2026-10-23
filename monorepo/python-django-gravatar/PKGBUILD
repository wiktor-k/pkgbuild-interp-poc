# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-django-gravatar
_name="${pkgname#python-}"
pkgver=1.4.5
pkgrel=4
pkgdesc="Essential Gravatar support for Django"
arch=(any)
url="https://github.com/twaddington/django-gravatar"
license=(MIT)
depends=(
  python
  python-django
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=($pkgname::git+$url.git#tag=$pkgver)
sha512sums=('e3dc2ef4b80b55869de9342ceab72f4fd318d2ddec9c9fa53dad89777892af1c988e2c91dba030c9c84eedc5369274392f063a0f7a86e40eb6d6a996269ac004')
b2sums=('84c7af234b318d95ba3408593fd80f759683444e6d77cd3fed4154a8aed4629383676dbc8d63d300882710219138e4c7fc9fc0fee3138d91420f7a2d02bdcb6a')

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname
  export PYTHONPATH="build:$PYTHONPATH"
  python example_project/manage.py test --verbosity=2 django_gravatar
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
