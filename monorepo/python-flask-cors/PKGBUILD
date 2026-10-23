# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Johannes Löthberg <johannes@kyriasis.com>

pkgname=python-flask-cors
pkgver=6.0.5
pkgrel=1
pkgdesc='A Flask extension adding a decorator for CORS support'
url='https://github.com/corydolphin/flask-cors'
arch=('any')
license=('MIT')
depends=(
  'python'
  'python-flask'
  'python-werkzeug'
)
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-setuptools-scm'
  'python-wheel'
)
source=("git+$url.git#tag=$pkgver")
b2sums=('ce5e4043abc9b814e9c5b5336cafe2af2c63e9656596c4b20db19a06ba5f0ae39cc3abd47fcc9deab2cb2ae895469d2a5f8687049b8cfcc01a246bc4c5e33fa5')

build() {
	cd ${pkgname#python-}
	python -m build --wheel --no-isolation
}

check() {
	cd ${pkgname#python-}
	# Upstream uses nosetests, which appears to test all Pyton files,
	# while unittest discovers test*.py only by default.
	python -m unittest discover -p '*test*.py'
}

package() {
	cd ${pkgname#python-}
	python -m installer --destdir="$pkgdir" dist/*.whl
	install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
