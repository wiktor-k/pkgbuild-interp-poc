# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Mubashshir <ahmubashshir@gmail.com>
# Contributor: Frederik “Freso” S. Olesen <freso.dk@gmail.com>

pkgname=python-confuse
pkgver=2.3.0
pkgrel=1
pkgdesc='Painless YAML config files for Python'
arch=(any)
url=https://github.com/beetbox/confuse
license=(MIT)
depends=(
  python
  python-typing_extensions
  python-yaml
)
makedepends=(
  git
  python-build
  python-poetry-core
  python-installer
  python-click
  python-tomli
  python-sphinx
  python-sphinx_rtd_theme
)
checkdepends=(python-pytest)
source=("$pkgname::git+https://github.com/beetbox/confuse#tag=v$pkgver")
b2sums=('d9b3a9fd89ddc8ae10f01285565c6a056aa6a5835d3cf3f9192eba3ef113100a9585d4492d71fff9c2d84137abee92ca393f829624d9adfa18b6a21663c151d4')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  python -m unittest discover -vs .
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
# vim: ts=2 sw=2 et:
