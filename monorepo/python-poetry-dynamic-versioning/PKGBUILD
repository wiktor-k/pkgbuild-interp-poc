# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-poetry-dynamic-versioning
pkgver=1.10.1
pkgrel=1
pkgdesc='Plugin for Poetry to enable dynamic versioning based on VCS tags'
arch=(any)
url='https://github.com/mtkennerly/poetry-dynamic-versioning'
license=(MIT)
depends=(
  python
  python-dunamai
  python-tomlkit
  python-jinja
  python-poetry
  python-cleo
  python-poetry-core
  python-packaging
)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
)
checkdepends=(
  python-pytest
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('0d204403bae9e3fd1b23be05a68eae90cebde8fc448ffc41e862d5f62458f2f08666af7856e78fac8c7e18964a47e7c3e3aecac9ed40063a345c73fabda41c53')
b2sums=('a187056ccbaf8ffed25bab9b391947993556a192c367485699f0cd2ea7c79e17481f3eed891f90c8a285a656b639746d7b33dec24e5b31b576b223c0e8a16892')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

# poetry-related stuff is just icky
#check() {
#  cd "$pkgname"
#
#  # integration tests are sketchy
#  pytest -v tests/test_unit.py
#}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
