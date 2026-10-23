# Maintainer: David Runge <dvzrv@archlinux.org>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Johannes Dewender   arch at JonnyJD dot net

pkgname=python-discid
pkgver=1.4.2
pkgrel=1
pkgdesc="Python binding of MusicBrainz libdiscid"
url="https://github.com/JonnyJD/python-discid"
license=(LGPL-3.0-or-later)
arch=(any)
depends=(
  libdiscid
  python
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('86afde3063142a47ca42d1c412ff1d657e9cda958082e686744b2590e844f079e71766c79961b678d8a138f3af4d67e6ac27b7505b7b47848f5d27755786b206')
b2sums=('d052aface9ce4b65ac7029ac431d7472e6c0c9c969a1a3b0dbbbf2362a11bb70696d45f1e33ddbe02e220820c0bd34aba88969d7709b6ff475456c16fe025c55')

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname-$pkgver
  pytest -vv
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGES,README}.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
