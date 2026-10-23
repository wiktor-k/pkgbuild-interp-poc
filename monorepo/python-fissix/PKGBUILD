# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Eli Schwartz <eschwartz@archlinux.org>

pkgname=python-fissix
pkgver=24.4.24
pkgrel=3
pkgdesc='Backport of lib2to3, with enhancements'
arch=(any)
url="https://github.com/amyreese/fissix"
license=(PSF-2.0)
depends=(
  python
  python-appdirs
)
makedepends=(
  git
  python-build
  python-flit-core
  python-installer
)
checkdepends=(
  python-pytest
  python-tests
)
source=(
  "$pkgname::git+$url#tag=v$pkgver"
  remove-version-constraints.patch
)
sha512sums=('cc8a94812e7dee1d47f067b3f35e5634d2e8ba70533425985d677e03bf8d0cb9c2ccd439af83a0c8c4c62fb9b2044202423fd11729f179deb3fcb8f8cc3ba980'
            '0c38d36ff481acccde9f64f88cadb0cc01d727a7d4a2022669f0bed5379ffb3cb0dbf847ef10fb6a5054976ca429954b3aa4da57b2f516383d48ae12590ef3be')
b2sums=('4baea7a0818aecbfb60bb310bfa4cec8ef2cdad16d6dd73cddd416d48751b7253d52f92d0880d24d6cbb5a5dd0f896fe18fabd3818cb6bbcd18d61b645ab0793'
        '5144a9affb441d22802e158d9748c38055625c9b00e579155d51d4ce8e454721830245e0bbbc654b248bdfb39d62da83c8c9cfbb7c781ae017d05a4a09711cda')

prepare() {
  cd "$pkgname"

  # Remove upper bound on flit-core
  patch -p1 -i "$srcdir/remove-version-constraints.patch"
}

build(){
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
