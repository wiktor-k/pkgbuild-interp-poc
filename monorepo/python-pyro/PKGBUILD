# Maintainer: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Angel Velasquez <angvp@archlinux.org>
# Contributor: Douglas Soares de Andrade <dsa@aur.archlinux.br>


# TODO: Switch to Pyro5: https://github.com/irmen/Pyro5
pkgname=python-pyro
pkgver=5.17
pkgrel=1
pkgdesc='Distributed object middleware (RPC)'
arch=(any)
url='https://github.com/irmen/Pyro5'
license=(MIT)
depends=(
  python
  python-serpent
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('a56626737c858e18d72e153adb064fb1244a4079658149ee11f98e737a6b1320d0a2adfd46fa60718c12fe36a197d908fb485d80bf819346038b5c8181e6965a')
b2sums=('82d9a3e1daffa106df0a0e78fb4029ad64aa743f4dd6eb12a06c9c08e7fbd0b07974973f92f40646c3973a42ab4498fd2a0ac1415bb80263a939fc448809a578')

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
