# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Guillaume Horel <guillaume.horel@gmail.com>

pkgname=python-psycopg
pkgver=3.3.6
pkgrel=1
pkgdesc='PostgreSQL database adapter for Python'
arch=(x86_64)
url='https://www.psycopg.org/psycopg3/'
license=(LGPL-3.0-only)
depends=(
  glibc
  python
  postgresql-libs
)
makedepends=(
  git
  cython
  python-build
  python-installer
  python-wheel
  python-setuptools
)
checkdepends=(python-pytest python-pytest-asyncio)
source=(
  "$pkgname::git+https://github.com/psycopg/psycopg.git#tag=$pkgver"
  remove-version-constraint.patch
)
sha512sums=('5c0959b55f8cc257fa635f1686f315e2b0147d9d9efa453cf0c0442bdebc4955c37598dc0b79e2273a87bac718e80dc6936ea6fa58ba2ed92c17a91a2e3a6698'
            'ed8107454e0384003a2108a985f9f573b635c0838b9c45d91b67dd4c1578bdbdceb54f1ffaf17a323a83ac454d2f8e6f72e24ef58367c9175891212c70470b39')
b2sums=('0680426673d1644d9f1dbe4c865ec66c8863d56dcc9bed231f55e9a8d80b16a32a7bf9c81bedc696b73f996399973b411b0071a5fa15d4de54ce9a16964b4ced'
        'acd2b3e276efca144d6602116670fcc062fbb079b717ff9bfc32412c4850892f88af9f7c51b8bc5835338210a71ffa26bb0841ea48a4a505cd9ec054248724a7')

prepare() {
  cd "$pkgname"

  patch -p1 -i "$srcdir/remove-version-constraint.patch"
}

build(){
  cd "$pkgname"

  for pkg in psycopg psycopg_c; do
    pushd "$pkg"
    python -m build --wheel --no-isolation
    popd
  done
}

package(){
  cd "$pkgname"

  for pkg in psycopg psycopg_c; do
    pushd "$pkg"
    python -m installer --destdir="$pkgdir" dist/*.whl
    popd
  done
}
