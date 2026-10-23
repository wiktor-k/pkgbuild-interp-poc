# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-tldextract
pkgver=5.4.0
pkgrel=1
pkgdesc="Accurately separate the TLD from the registered domain and subdomains of a URL, using the Public Suffix List"
arch=(any)
url='https://github.com/john-kurkowski/tldextract'
license=(BSD-3-Clause)
depends=(
  python
  python-idna
  python-requests
  python-requests-file
  python-filelock
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-mock
  python-responses
  python-sybil
  python-syrupy
)
source=("$pkgname::git+https://github.com/john-kurkowski/tldextract.git#tag=$pkgver")
sha512sums=('5a1e0eaaec52341f1e21247d62ebcf953591a87d22b6658c288f283637526f295f368d0d6031d64548df1207a49edc1c280b15124fc4e641b70b009e1a8657cb')
b2sums=('8dcb5eefbafbb9703c158c2de7b5a23f5f6b8a8199dfbe110933db4d175b59fbf1dafeb6254789bf54ba27952ac4af3509cd6d703dbe1319f791e364e4f585d5')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}

# vim:set sw=2 et:
