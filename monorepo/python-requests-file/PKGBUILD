# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Andrzej Giniewicz <gginiu@gmail.com>
# Contributor: lilydjwg <lilydjwg@gmail.com>

pkgname=python-requests-file
pkgver=2.1.0
pkgrel=3
pkgdesc='File transport adapter for Requests'
arch=('any')
url='https://github.com/dashea/requests-file'
license=('Apache-2.0')
depends=('python' 'python-requests')
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-setuptools-scm'
  'python-wheel'
)
source=("$pkgname::git+$url#tag=$pkgver")
sha512sums=('29cb33b3fc71fc865d5c729729885d797adfe129cf87feeb3443b97f3e04e3b599b08a4ffd989505f833d878d72b77a33d8ac537d976ec869d30fbcc22989a06')
b2sums=('460e5940c2f32eb8a114884eb766cd7ad90a968da6b95c7d129a6ef0ea1b13e35b11fa0cd9eff5cece4fbba6bca2de9aca6720ad9bd4b3b57a1d781d80c147e9')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"
  python -m unittest discover -vs tests
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
