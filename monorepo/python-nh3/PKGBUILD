# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Marcus Hoffmann <bubu@bubu1.eu>

pkgname=python-nh3
pkgver=0.3.7
pkgrel=1
pkgdesc='Ammonia HTML sanitizer Python binding'
arch=(x86_64)
url='https://github.com/messense/nh3'
license=(MIT)
depends=(
  libgcc
  glibc
  python
)
checkdepends=(python-pytest)
makedepends=(
  git
  python-build
  python-installer
  python-maturin
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('1b07eac79c7c82affaf3efb7d5ea81b7eebe5ccec68a590e4b4a5c8ad2069070b2521aeeea69f7e2fe01d4b0baa1694d69cb219f0aaa52ef4c66857466571a17')
b2sums=('9b57c1e285d3d15e624c337d584c56727b21426bf83a0e69cb45feb7dc793387639c78c82fc1268ed0f5c8cd645df61977964e9a2f320efe456515933687c55b')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation 
}
 
check() {
  cd "$pkgname"

  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
