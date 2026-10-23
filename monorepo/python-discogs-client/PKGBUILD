# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Guillaume Hayot <ghayot@postblue.info>
# Contributor: Arvedui <arvedui@posteo.de>
# Contributor: Marc Plano-Lesay <marc.planolesay@gmail.com>
# Contributor: Joost Bremmer <toost.b@gmail.com>

pkgname=python-discogs-client
pkgver=2.10
pkgrel=1
pkgdesc='Python Client for the Discogs API'
arch=(any)
url=https://github.com/joalla/discogs_client
license=(BSD-2-Clause)
depends=(
  python
  python-dateutil
  python-requests
  python-oauthlib
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=("$pkgname::git+https://github.com/joalla/discogs_client.git#tag=v$pkgver")
sha512sums=('f35d1bc8a99c5974a7c69dc0b8e7c44e95d6b026db36fb8e7150fb3044571caf68812008bdfe4f1c52c31b1548608f2773aa3f5bd1db430697e9241a0aabb6e6')
b2sums=('b92bec31ba7ecb7c1ab8c2ef21746d6d88afd4ab94dc8383ba695be24462de87551d7c2087a0f8fdeaa3c7c41b8ae6fa7b41f611078174f69207c8e080135ff7')

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
# vim: set ts=2 sw=2 et:
