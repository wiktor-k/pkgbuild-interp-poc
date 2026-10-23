# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-feedparser-sgmllib
pkgver=2.1.0
pkgrel=1
pkgdesc='Provides sgmllib from Python 2'
arch=(any)
url='https://github.com/python-syndication/feedparser-sgmllib'
license=(PSF-2.0)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
)
checkdepends=(python-pytest)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('b0d614b99785681c4e48341b73a5dce3a6aa00915c44e27a8e16d1a54518180bb791ed5e24f3e76f878dd5bf72227e78d81d90dda85c80b53e1246130767b15e')
b2sums=('84a219f3b6eedd8ebcba0ea38ec9d5c508aa399f816ae0a7ddbae5892d1447d46fdf896bbb224e443b033faa84a1afd8f3cdc0800ac25f1d45ff4f6a0e0100c3')

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

  local pytest_options=(
    -vv
  )
  pytest "${pytest_options[@]}"
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
