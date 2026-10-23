# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>

pkgname=python-readme-renderer
pkgver=46.0
pkgrel=1
pkgdesc='A library that will safely render arbitrary README files into HTML'
arch=(any)
url='https://github.com/pypa/readme_renderer'
license=(Apache-2.0)
depends=(
  python
  python-docutils
  python-pygments
  python-nh3
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
  python-comrak
)
checkdepends=(python-pytest)
optdepends=('python-comrak: for Markdown support')
source=("$pkgname::git+$url#tag=$pkgver")
sha512sums=('fe8f9d01905cfca7904a0691425eb16982acc587c67fa6d84b0730258f9e313b8b3f1628e622b9e3057b828d18981feeec6e8b717ec53e40264e58a39e4b2c59')
b2sums=('f9d66559c82bec03e4d2687d1df8c674f7ba283c3dc766e1d8ae59f7723c349ad954a8feedaaf3203526293b85375ec7d32823f0f3903bcc7002556626eedce6')

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
}
