# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Jelle van der Waa <jelle@vdwaa.nl>
# Contributor: Eli Schwartz <eschwartz@archlinux.org>
# Contributor: Alexander Rødseth <rodseth@gmail.com>
# Contributor: Angel 'angvp' Velasquez <angvp[at]archlinux.com.ve>
# Contributor: Roman Kyrylych <Roman.Kyrylych@gmail.com>
# Contributor: iom <pizorn@gmail.com>

pkgname=python-pychm
pkgver=0.8.6
pkgrel=11
pkgdesc='Python bindings for chmlib'
arch=(x86_64)
url='https://github.com/dottedmag/pychm'
license=(GPL-2.0-only)
depends=(
  glibc
  python
  chmlib
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('379a60fac3f67cfc9c9eb2fa06768bb4bb7f49728cfd054a37994a8015adf2b1eb9d6155cf0202b433ac0f327c4aea1f3ceac12adf95c554ad229f132edaff58')
b2sums=('36c8bb95b9c406d96ae1b558fbfc52a60f51c4aed7a6992b7ad11c98b00379343d95199bc5a50f9915cc0af4b8a5b3d54f124c964a724cbd339c11d3f21efb4d')

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
