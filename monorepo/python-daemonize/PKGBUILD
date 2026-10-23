# Maintainer: Johannes Löthberg <johannes@kyriasis.com>
# Contributor: Ivan Shapovalov <intelfx@intelfx.name>

pkgname=python-daemonize
pkgver=2.5.0
pkgrel=10

pkgdesc='Library to enable your code run as a daemon process on Unix-like systems'
url='https://pypi.python.org/pypi/daemonize/'
arch=('any')
license=('MIT')

depends=('python')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')

source=("https://pypi.org/packages/source/d/daemonize/daemonize-$pkgver.tar.gz")

sha256sums=('dd026e4ff8d22cb016ed2130bc738b7d4b1da597ef93c074d2adb9e4dea08bc3')

build() {
	cd "$srcdir"/daemonize-$pkgver
	python -m build --wheel --no-isolation
}

package() {
	cd daemonize-$pkgver
	python -m installer --destdir="$pkgdir" dist/*.whl

	install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}

# vim: set ts=4 sw=4 tw=0 ft=PKGBUILD :
