# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-comrak
pkgver=0.0.16
pkgrel=1
pkgdesc='Python bindings for the Comrak Rust library, a fast CommonMark/GFM parser'
arch=(x86_64)
url='https://github.com/lmmx/comrak'
license=(BSD-2-Clause)
depends=(
  python
  glibc
  libgcc
)
makedepends=(
  git
  python-build
  python-installer
  python-maturin
  oniguruma
)
checkdepends=(python-pytest)
source=("$pkgname::git+$url#tag=py-$pkgver")
sha512sums=('75249f0ffdc49c3473784c94a5233bbf7bd7d964d06423b6f0e252b4dfa0fb1e8fc8ffc586547025c0ce7ea9893518328168e24d55acfa7e31d30c48b9ac98d5')
b2sums=('c0f9ba5aef9d2af1549ae74d4f57d0538bf48fcd7b34258abfe964b46cc4861ba24eb0db7a2f34e6bc514de8b619e96c4e0a203696a49f69e6dc74a3988b6535')

prepare() {
  cd "$pkgname"

  # download dependencies
  cargo fetch --locked --target host-tuple
}

build() {
  cd "$pkgname"

  # https://archlinux.org/todo/lto-fat-objects/
  CFLAGS+=' -ffat-lto-objects'

  export RUSTONIG_DYNAMIC_LIBONIG=1

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # temporary install
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$(pwd)/tmp/$site_packages"

  pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
