# Maintainer: Peter Jung <ptr1337@archlinux.org>

pkgname=python-setuptools-git-versioning
_pkgname=setuptools_git_versioning
pkgver=3.2.0
pkgrel=1
pkgdesc='Use Git repo data for building a version number according to PEP 440.'
url='https://github.com/dolfinus/setuptools-git-versioning'
arch=('any')
license=('MIT')
depends=('python' 'python-setuptools' 'python-packaging')
makedepends=('python-build' 'python-installer' 'python-wheel')
source=("https://pypi.io/packages/source/s/$_pkgname/$_pkgname-$pkgver.tar.gz")
sha256sums=('27aa1ad0409b632ee49947ea592cc27e16fbb2a1605cc22efeb7c353fd790b31')

build() {
  cd "$srcdir/$_pkgname-$pkgver"
  python -m build --wheel --no-isolation
}

package() {
  cd "$srcdir/$_pkgname-$pkgver"
  python -m installer "--destdir=$pkgdir" "./dist/"*".whl"
  install -Dm644 "./LICENSE" "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
