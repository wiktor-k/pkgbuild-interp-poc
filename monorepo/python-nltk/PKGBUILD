# Maintainer: Alexander F Rødseth <xyproto@archlinux.org>
# Maintainer: Morten Linderud <foxboron@archlinux.org>
# Contributor: Thomas Dziedzic <gostrc@gmail.com>
# Contributor: Chris Baker <baker.chris.3@gmail.com>

pkgname=python-nltk
pkgver=3.10.3
pkgrel=1
pkgdesc='Natural language processing in Python'
arch=('any')
url='https://www.nltk.org/'
license=('Apache-2.0')
depends=('python-yaml' 'python-click' 'python-regex')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('nltk-data' 'python-requests' 'python-regex')
optdepends=('nltk-data: test data'
	    'python-numpy: used for calculations'
	    'python-matplotlib: used for plotting')
source=(git+https://github.com/nltk/nltk#tag=v$pkgver)
sha256sums=('10625f2a7c648af168d91ee5bc806c05a5f06bbc6373c658d49a6fe88db04276')

build() {
  cd nltk
  python -m build --wheel --no-isolation
}

# Need to fix the nltk-data split
#check(){
#  cd nltk-$pkgver
#  python ./nltk/test/runtests.py
#}

package() {
  cd nltk
  python -m installer --destdir="$pkgdir" dist/*.whl
}
