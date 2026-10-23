# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Fernando Jiménez Solano <fjim@sdfeu.org>
# Contributor: Sven-Hendrik Haase <svenstaro@archlinux.org>

# TODO follow this PR and related issue(s)
# https://github.com/beetbox/pyacoustid/pull/90

pkgname=python-pyacoustid
pkgver=1.3.1
pkgrel=1
pkgdesc='Bindings for Chromaprint acoustic fingerprinting and the Acoustid API'
arch=(any)
url='https://github.com/beetbox/pyacoustid'
license=(MIT)
depends=(
  chromaprint
  python
  python-requests
  python-audioread
)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('bfe9ac4648cb694a9c531a2b90caf9c19f4dd8c0627257935ec264d285590c8b06645cc82943c2519d57ee15383f087e62dfd8d413f81fd79de2c1a063e8ceb9')
b2sums=('8de868114135d9ed6b9ae070c0af3e8c64a267489ee6dbe6ba746c94032199ab5c53607433d4d459211ff0139ae30bc5dcb26449d1965fd06c9a66603b435752')

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
