# Maintainer: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: razer <razer@neuf.fr>

pkgname=python-serpent
pkgver=1.43
pkgrel=1
pkgdesc='Serializer for literal Python expressions'
url='https://github.com/irmen/Serpent'
arch=(any)
license=(MIT)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-attrs python-pytz)
source=("$pkgname::git+$url#tag=serpent-$pkgver")
sha512sums=('aaefa719f1c819270be8169e3ab38492842df605dbefc9a53f0a0737ba1bdb1e11f02a21de258dc79b6ae3e3d6317f6499068850c016d5d2951431ba3279f94b')
b2sums=('a229e6dd5cc23b6d5b01279e68f02f7001e28db560ce8405fefc0086de8052297a63419e9089a88ad81c5df4fa1e2b12130e774ccb4ff872e3b5323f4472569e')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  python3 -m unittest discover -vs tests
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
