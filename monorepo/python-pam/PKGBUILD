# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Eli Schwartz <eschwartz@archlinux.org>

pkgname=python-pam
pkgver=2.1.0
pkgrel=1
pkgdesc='Pure Python interface to the Pluggable Authentication Modules system'
arch=(any)
url='https://github.com/FirefighterBlu3/python-pam'
license=(MIT)
depends=(python pam)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
)
checkdepends=(python-pytest)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('b8ad916b7355b79629f89738a55747858639e3da42525878604f217587361952c47900a0feeed129f3e3a4884c96997e08c603a4f0cb3e1c0d46c56f7d2188cc')
b2sums=('51042dfdf6aedea748b049cee7687ee8ac9fd53740f967916686689200c343e07b7ffee3f513ae473ee372f87c8205571dcac0e584f3e1c97e2856269b262923')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  # install to temporary directory
  python -m installer --destdir="$PWD/tmp_install" dist/*.whl

  PYTHONPATH="$PWD/tmp_install$site_packages" pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
