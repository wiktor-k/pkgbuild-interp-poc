# Maintainer: Bert Peters <bertptrs@archlinux.org>
# Contributor: Evangelos Foutras <foutrelis@archlinux.org>

pkgname='python-virtualenv-clone'
pkgver=0.5.7
pkgrel=8
pkgdesc="A script for cloning a non-relocatable virtualenv."
arch=('any')
url="https://github.com/edwardgeorge/virtualenv-clone"
license=('MIT')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-setuptools')
source=(https://files.pythonhosted.org/packages/source/v/virtualenv-clone/virtualenv-clone-$pkgver.tar.gz)
sha256sums=('418ee935c36152f8f153c79824bb93eaf6f0f7984bae31d3f48f350b9183501a')

build() {
  cd "$srcdir/virtualenv-clone-$pkgver"
  python -m build --wheel --no-isolation
}

package() {
  cd "$srcdir/virtualenv-clone-$pkgver"

  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}

# vim:set ts=2 sw=2 et:
