# Maintainer: Maxime Gauduin <alucryd@archlinux.org>

pkgname=python-reflink
pkgver=0.2.2
pkgrel=6
pkgdesc='Reflink support for Python'
arch=(x86_64)
url='https://gitlab.com/rubdos/pyreflink'
license=(MIT)
depends=(
  glibc
  python
  python-cffi
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=("$pkgname::git+$url#tag=$pkgver")
sha512sums=('7814d63d294f37a5de46f13dccb14db84ebc92c23916cd8f165d43b0330a1e162dd1f543c4e8788bf7edb5b9c976954c3356e778e77b5a96653379e73632efb8')
b2sums=('2b6e4c572945d3598664f49853b248ddaa0689e3286d891011717ed19dc868c6f5ce1b60278155e9d7f3bd656bc321ca40c17a843071330c32bdd3096289f47f')

prepare() {
  cd "$pkgname"

  # remove unnecessary dependency
  sed \
    -e '/pytest-runner/d' \
    -i setup.py
}

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
# vim: ts=2 sw=2 et:
