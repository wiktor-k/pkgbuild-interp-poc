# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-localzone
pkgver=0.9.8
pkgrel=8
pkgdesc='A simple library for managing DNS zones'
arch=(any)
url='https://localzone.iomaestro.com'
license=(BSD-3-Clause)
depends=(
  python
  python-dnspython
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=("$pkgname::git+https://github.com/ags-slc/localzone#tag=v$pkgver")
sha512sums=('b3321e338f62d720f083d2569707cb19fea1e357de00e7bf053caa7ea405d2eaa07023f86cf43431f4d8cfd39171488ffc2dabb59885f4c4ad3267a6f338a6c1')
b2sums=('f485584d18d2a622c2dd14f58e56c7e0fb9eb18eadcf1b61498a4a15d7ddfffa1fa365ec1b34da06505ef05a11268c26ced6e39ee829d44e623557f8f38d5bce')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  pytest -v --deselect tests/test_models.py::test_zone_add_record_no_content
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
