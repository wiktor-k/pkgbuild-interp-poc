# Maintainer:

_pyname=pykakasi
pkgname=python-$_pyname
pkgver=2.3.0
pkgrel=3
pkgdesc='Lightweight converter from Japanese Kana-kanji sentences into Kana-Roman'
url='https://codeberg.org/miurahr/pykakasi'
license=(GPL-3.0-or-later)
arch=(any)
depends=(python
         python-deprecated
         python-jaconv)
makedepends=(git
             python-build
             python-installer
             python-setuptools-scm)
checkdepends=(python-pytest)
source=(git+https://codeberg.org/miurahr/pykakasi#tag=v$pkgver)
sha256sums=('SKIP')

prepare() {
  cd $_pyname
  git cherry-pick -n c27fa6eb714af6e34e302e479456f1cc68529f46 # Fix tests
  git cherry-pick -n e51d045f6bc2d70538c24a9997ca90500019d567 # Don't benchmark
}

build() {
  cd $_pyname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pyname
  pytest -v -m 'not benchmark'
}
  
package() {
  cd $_pyname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 COPYING -t "$pkgdir"/usr/share/licenses/$pkgname
}
