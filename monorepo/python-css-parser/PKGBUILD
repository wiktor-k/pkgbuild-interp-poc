# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Eli Schwartz <eschwartz@archlinux.org>

pkgname=python-css-parser
pkgver=1.1.1
pkgrel=1
pkgdesc='CSS related utilities for Python'
arch=(any)
url='https://github.com/ebook-utils/css-parser'
license=(LGPL-3.0-or-later)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('1fb42aa5e6995f2f6f1ff1a167e8254878f6bff23647bf3b96f48b1aa265546aa9e5594ac77ebe8c08d4572e4cbc5d2b726ff967d6530adcda262f8a939db61f')
b2sums=('611bb527b0a2418cc3b724e1900e28d7468fd4a92eb1585376fa448430ad2f4fa309f87f0058aca75721ea9db506583c79603b366e26cc527d4a715849951b9e')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # temporary install
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$(pwd)/tmp/$site_packages"

  pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
