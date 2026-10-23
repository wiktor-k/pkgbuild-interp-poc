# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Gordian Edenhofer <gordian.edenhofer[at]yahoo[dot]de>

pkgname=certbot-nginx
pkgver=5.8.0
pkgrel=1
pkgdesc='Nginx plugin for Certbot'
arch=(any)
license=(Apache-2.0)
url='https://pypi.python.org/pypi/certbot-nginx'
depends=(
  "certbot=$pkgver"
  "python-acme=$pkgver"
  python-pyopenssl
  python-pyparsing
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
# git repository is used because certbot is a huge monorepo and it's easier to
# share the entire repository across all certbot related packages than a few
# hundred tarballs.
_repo='github.com-certbot-certbot'
source=("$_repo::git+https://github.com/certbot/certbot#tag=v$pkgver")
sha512sums=('71d9a98365de26f7a8eeef40116f2afefdb4a539fa6d080024f1f5773299222974d79340d05fac83953aaffaa31ac169220a56001a8e3889b483365a65c8f78c')
b2sums=('969d082dfd6543e6c62e8374d9cead8889b62c3956b95055a39361423aeaf4f33b879b7b2958d1336b86d5d7d30190de7b0144f4f3e76bc305c73a2576b870e2')

build() {
  cd "$_repo/$pkgname"

  python -m build --wheel --no-isolation
}

package() {
  cd "$_repo/$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
