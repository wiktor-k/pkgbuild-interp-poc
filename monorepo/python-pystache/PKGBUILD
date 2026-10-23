# Maintainer:
# Contributor: Florian Klink <flokli at flokli dot de>

pkgname=python-pystache
_pkgname=pystache
pkgver=0.6.8
pkgrel=1
pkgdesc='The mustache template engine written in python'
arch=(any)
url='https://github.com/PennyDreadfulMTG/pystache'
license=('MIT')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-setuptools-scm' 'python-wheel')
source=("https://github.com/PennyDreadfulMTG/$_pkgname/archive/v$pkgver/$_pkgname-$pkgver.tar.gz")
sha512sums=('b2eecabdc39f86713b4f3387b4a80afaa0a80ecc0a34482e3d9341467660666a7bdcd83d21fc9fd3b61a764a1c840a5ba8041f9a66a3b130c849e8bd8cf1f9b5')

build() {
  cd "$_pkgname-$pkgver"
  python -m build --wheel --no-isolation
}

check() {
  cd "$_pkgname-$pkgver"
  # This package uses 2to3 to convert itself to Python 3 on the fly
  # So we need to jump through some hoops here
  test_dir="../test_dir"
  rm -rf "$test_dir"
  mkdir "$test_dir"
  python -m installer --destdir="$test_dir" dist/*.whl
  PYTHONPATH=$(readlink -e "$test_dir"/usr/lib/python*/site-packages/) \
  	"$test_dir"/usr/bin/pystache-test .
}

package() {
  cd "$_pkgname-$pkgver"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/"$pkgname"/LICENSE
}
