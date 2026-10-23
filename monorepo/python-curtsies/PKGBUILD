# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Kyle Keen <keenerd@gmail.com>

pkgname=python-curtsies
pkgver=0.4.3
pkgrel=2
pkgdesc='Curses-like terminal wrapper with a display based on compositing 2d arrays of text'
arch=(any)
url='https://github.com/bpython/curtsies'
license=(MIT)
depends=(
  python
  python-blessed
  python-cwcwidth
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pyte python-pytest)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('0e47e42c0b96a6235e5316101e3680d50b5b094ecf479ff625ca2b42f9a4d268a557fa133b8fa0497c5c68fcf31e12b91c0860e02e4147e17d1d0c831535166c')
b2sums=('16a27151e6083bac70b8a87bd3d7c46f00559b74043b1ed4c261a727d8a33745589722fdebb4099a9017a58de137cba84b13401e61e2b60f1b6e2e6f1d2b2e67')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # temporary install so pytest can pick up the compiled extension
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$(pwd)/tmp/$site_packages"

  pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
